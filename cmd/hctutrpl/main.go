package main

import (
	"flag"
	"hctutrpl/cmd/hctutrpl/app"
)

var (
	method   = flag.String("method", "GET", "http request method")
	urlInput = flag.String("url", "http://127.0.0.1:8080", "request URL")
	output   = flag.String("output", "", "file to store request result")
)

func init() {
	flag.StringVar(method, "m", "GET", "http request method")
	flag.StringVar(urlInput, "u", "http://127.0.0.1:8080", "request URL")
	flag.StringVar(output, "o", "", "file to store request result")
	flag.Parse()
}

func main() {
	a := app.App{
		Url:    *urlInput,
		Method: *method,
	}

	a.Run()
}
