package ipieces

import (
	"fmt"
	"html/template"

	_ "embed"
)

var (
	//go:embed static/style.css
	stylesheet string
	//go:embed static/index.tmpl
	index string
	//go:embed static/error.html
	errorPage string
	//go:embed static/rate-limited.html
	rateLimitPage string
	//go:embed static/vpn.html
	vpnPage string
	//go:embed static/lights.gif
	lights string

	// All Templates are meant to be executed with a `data` struct.
	//
	// HTTP 200 template
	indexTmpl = template.Must(template.New("").Funcs(template.FuncMap{
		"comment": func(ip, hash string, index int) template.HTML {
			return template.HTML(fmt.Sprintf("<!--\n  SHA-256(%q) = %s\n  index = %d\n-->", ip, hash, index))
		},
		"format": func(d Digit) template.HTML { return d.format() },
	}).Parse(index))

	// /text template
	textTmpl = template.Must(template.New("").Parse("IP: {{.IP}} Digits[{{.Index}}]: {{.Revealed}}\n"))
)
