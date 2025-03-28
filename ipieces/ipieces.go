// Package ipieces allows users to create IP address-based Geocaching puzzles.
//
// To create a puzzle, you need to populate a [Puzzle] struct and call [Puzzle.Run] on it.
// For example:
//
//	package main
//
//	import (
//	  "github.com/bitlux/caches/ipieces"
//	  "github.com/bitlux/vpnapi"
//	)
//
//	func main() {
//	  p := ipieces.Puzzle{
//	    Final: []ipieces.Digit{
//	      ipieces.Digit{Value: "3", Status: ipieces.VISIBLE}
//	      ipieces.Digit{Value: "7", Status: ipieces.VISIBLE}
//	      ipieces.Digit{Value: "2", Status: ipieces.VISIBLE}
//	      ipieces.Digit{Value: "4", Status: ipieces.HIDDEN}
//	      ipieces.Digit{Value: "0", Status: ipieces.HIDDEN}
//	      ipieces.Digit{Value: "0", Status: ipieces.HIDDEN}
//	      ipieces.Digit{Value: "0", Status: ipieces.HIDDEN}
//	      ipieces.Digit{Value: "1", Status: ipieces.VISIBLE}
//	      ipieces.Digit{Value: "2", Status: ipieces.VISIBLE}
//	      ipieces.Digit{Value: "2", Status: ipieces.VISIBLE}
//	      ipieces.Digit{Value: "0", Status: ipieces.VISIBLE}
//	      ipieces.Digit{Value: "4", Status: ipieces.HIDDEN}
//	      ipieces.Digit{Value: "0", Status: ipieces.HIDDEN}
//	      ipieces.Digit{Value: "0", Status: ipieces.HIDDEN}
//	      ipieces.Digit{Value: "0", Status: ipieces.HIDDEN}
//	    },
//	    IndexFunc: func(b [sha256.Size]byte) int {
//	      return int(b[sha256.Size-1]) % 8
//	    },
//	    // Setting Client is optional.
//	    Client:   vpnapi.New("YOUR-API-KEY-HERE"),
//	    Backdoor: "topsecret",
//	    GCCode:   "GCB2PKC",
//	  }
//	  p.Run()
//	}
//
// [Puzzle.Run] creates two handlers:
//   - a text endpoint at /text which responds with a short plaintext page with the client's IP,
//     the computed index into the final coordinates, and the revealed coordinate, and
//   - a default endpoint, which serves any path other than /text, and responds with an HTML page.
package ipieces

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"net/netip"
	"os"
	"slices"

	"github.com/bitlux/vpnapi"
)

// TODO: All logging is printed to stdout. This is fine for Google Cloud Run,
// but consider using log/slog.

func init() {
	hostname, _ := os.Hostname()
	fmt.Printf("STARTING on %s, pid %d\n", hostname, os.Getpid())
}

// VPNChecker determines whether a client's IP address belongs to a VPN or proxy. Its concrete
// implementation is *github.com/bitlux/vpnapi.Client. In order to use vpnapi.Client, you must
// first obtain an API key from http://vpnapi.io.
type VPNChecker interface {
	Query(string) (*vpnapi.Response, error)
}

// Display controls how a Digit is displayed.
type Display int

const (
	// Visible by default
	VISIBLE Display = iota
	// Hidden by default
	HIDDEN
	// Hidden by default, revealed in this rendering
	REVEALED
)

// A Digit is one of the digits that make up the coordinates of the final.
type Digit struct {
	// Value is a single digit, in string form.
	Value string
	// Status is how this digit should be displayed: visible, hidden, or revealed.
	Status Display
}

func (d Digit) format() template.HTML {
	switch d.Status {
	case VISIBLE:
		return template.HTML(d.Value)
	case HIDDEN:
		return template.HTML(`<span class="underline">&nbsp;</span>`)
	case REVEALED:
		return template.HTML(fmt.Sprintf(`<span class="red">%s</span>`, d.Value))
	default:
		return template.HTML("ERROR")
	}
}

type Puzzle struct {
	// Final is the full final coordinates of the puzzle.
	Final []Digit

	// IndexFunc determines which digit of Final is revealed. The return value must be less than the number
	// of hidden Digits, because it is used as an index into the hidden Digits in Final.
	IndexFunc func([sha256.Size]byte) int

	// Client determines how to handle requests from IP addresses that belong to VPNs or proxies.
	// If Client is nil, no VPN checking will be done. In order to do VPN checking, first obtain an
	// API key from http://vpnapi.io. Pass that key to github.com/bitlux/vpnapi.Client.New to create
	// a *vpnapi.Client, and set Client to that value.
	Client VPNChecker

	// Backdoor allows you to test how the server handles a specific IP address. Backdoor will be accepted
	// as an HTTP header name. The server will read the header value as the client's IP address. For
	// example, if Backdoor is "geocache" and the server is running on localhost:8080, the following
	// command will tell the server that the request is coming from IP 1.2.3.4:
	//   curl -H "geocache: 1.2.3.4" localhost:8080
	// Any string is a valid .
	Backdoor string

	// GCCode is used to link back to the puzzle on geocaching.com.
	GCCode string

	hiddenCount int
}

type data struct {
	IP       string
	Hash     string
	Digits   []Digit
	Index    int
	Revealed string
	GCCode   string
}

func ipFromHeaders(h http.Header, backdoor string) string {
	if vals := h.Values(("X-Forwarded-For")); len(vals) > 1 {
		fmt.Println("X-Forwarded-For:", vals)
	}
	if ip := h.Get(backdoor); ip != "" {
		fmt.Println("Setting IP to", ip, "via header")
		return ip
	}
	return h.Get("X-Forwarded-For")
}

func (p Puzzle) dataFromReq(req *http.Request) (*data, error) {
	ip := ipFromHeaders(req.Header, p.Backdoor)
	if ip == "" {
		ap, err := netip.ParseAddrPort(req.RemoteAddr)
		if err != nil {
			return nil, fmt.Errorf("IP error: ParseAddrPort(%s) returned %v", req.RemoteAddr, err)
		}
		ip = ap.Addr().String()
	}

	sha := sha256.Sum256([]byte(ip))
	d := &data{
		IP:     ip,
		Hash:   hex.EncodeToString(sha[:]),
		Digits: slices.Clone(p.Final),
		Index:  p.IndexFunc(sha),
		GCCode: p.GCCode,
	}
	d.flip()
	return d, nil
}

func (d *data) flip() {
	count := 0
	for i := range d.Digits {
		if d.Digits[i].Status == HIDDEN {
			if count == d.Index {
				d.Digits[i].Status = REVEALED
				d.Revealed = d.Digits[i].Value
				return
			}
			count++
		}
	}
}

func writeResponse(w http.ResponseWriter, code int, body string, format string, args ...any) {
	fmt.Printf(format, args...)
	w.WriteHeader(code)
	if _, err := io.WriteString(w, body); err != nil {
		fmt.Println("WriteString failed:", err)
	}
}

func (p Puzzle) handle(w http.ResponseWriter, req *http.Request, tmpl *template.Template) {
	d, err := p.dataFromReq(req)
	fmt.Printf("IP %s index %d\n", d.IP, d.Index)
	if err != nil {
		writeResponse(w, http.StatusInternalServerError, errorPage, "dataFromReq failed: %v\n", err)
		return
	}

	if p.Client != nil {
		resp, err := p.Client.Query(d.IP)
		fmt.Println(resp)
		if err != nil {
			if err == vpnapi.ErrRateLimited {
				writeResponse(w, http.StatusTooManyRequests, rateLimitPage, "rate limited: %v\n", err)
				return
			}
			// TODO: This fails closed. This may be too strict, especially if vpnapi.io is unreliable.
			writeResponse(w, http.StatusInternalServerError, errorPage, "Query failed: %v\n", err)
			return
		}

		if resp.Security.VPN || resp.Security.Proxy || resp.Security.Tor || resp.Security.Relay {
			writeResponse(w, http.StatusForbidden, vpnPage, "VPN: %t Proxy: %t Tor: %t Relay: %t\n", resp.Security.VPN, resp.Security.Proxy, resp.Security.Tor, resp.Security.Relay)
			return
		}
	}

	if err := tmpl.Execute(w, d); err != nil {
		writeResponse(w, http.StatusInternalServerError, errorPage, "tmpl.Execute failed: %v\n", err)
		return
	}
}

// Runs starts an HTTP server and blocks forever (or until a fatal error occurs).
func (p Puzzle) Run() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		p.handle(w, req, indexTmpl)
	})
	http.HandleFunc("/text", func(w http.ResponseWriter, req *http.Request) {
		p.handle(w, req, textTmpl)
	})
	http.HandleFunc("/style.css", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "text/css; charset=utf-8")
		if _, err := io.WriteString(w, stylesheet); err != nil {
			fmt.Println("WriteString failed:", err)
		}
	})
	http.HandleFunc("/lights.gif", func(w http.ResponseWriter, _ *http.Request) {
		if _, err := io.WriteString(w, lights); err != nil {
			fmt.Println("WriteString failed:", err)
		}
	})

	for _, d := range p.Final {
		if d.Status == HIDDEN {
			p.hiddenCount++
		}
	}

	// Google Cloud Run passes the port in the PORT environment variable.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Println("Listening on port", port)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
