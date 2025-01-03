package util

import (
	"fmt"
	"io"
	"net/http"
)

// Wget fetches the named URL and returns its contents. It exits on any error.
func Wget(url string) []byte {
	res, err := http.Get(url)
	Must(err)

	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	Must(err)
	if res.StatusCode > 299 {
		Must(fmt.Errorf("fail: status code %d", res.StatusCode))
	}
	return body
}
