package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, err := http.Get("http://techmanski.net/geocaching/GC5ECV4/longitude.txt")
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	sum := 0
	for _, b := range body {
		sum += int(b) - 48
	}

	fmt.Println(sum * 3)
}
