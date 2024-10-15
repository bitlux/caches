package util

import (
	"fmt"
	"io"
	"net/http"
)

func Wget(url string) []byte {
	res, err := http.Get(url)
	Must(err)

	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	Must(err)
	if res.StatusCode > 299 {
		Must(fmt.Errorf("Fail: status code %d", res.StatusCode))
	}
	return body
}
