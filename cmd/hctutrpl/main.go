package main

import (
	"flag"
	"fmt"
	"hctutrpl/internal/input"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

var (
	method = flag.String("method", "GET", "http request method")
	output = flag.String("output", "", "file to store request result")
)

func init() {
	flag.Parse()
}

func main() {
	urlInput, _ := input.ReadUrl()
	u, err := url.ParseRequestURI(urlInput)
	if err != nil {
		log.Fatal(err)
	}

	var req = &http.Request{
		Method: *method,
		URL:    u,
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("client: response body: %s\n", resBody)
}
