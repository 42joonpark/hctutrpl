package main

import (
	"flag"
	"hctutrpl/cmd/hctutrpl/app"
	"net/http"
)

var (
	method   = flag.String("method", "GET", "http request method")
	urlInput = flag.String("url", "http://127.0.0.1:8080", "request URL")
	output   = flag.String("output", "", "file to store request result")
	data     = flag.String("data", "'{}'", "data to send")
	header   = flag.String("header", `{"Content_Type": "application/json"}`, "header content type")
	user     = flag.String("user", "", "basic authorize user")
)

func init() {
	flag.StringVar(method, "m", "GET", "http request method")
	flag.StringVar(urlInput, "U", "http://127.0.0.1:8080", "request URL")
	flag.StringVar(output, "o", "", "file to store request result")
	flag.StringVar(data, "d", "{}", "data to send")
	flag.StringVar(header, "H", `{"Content_Type": "application/json"}`, "header content type")
	flag.StringVar(user, "u", "", "basic authorize user")
	flag.Parse()
}

func main() {
	a := app.App{
		Url:    *urlInput,
		Method: *method,
		Header: make(http.Header),
	}
	options := app.OptionFlag{
		Method:   method,
		UrlInput: urlInput,
		Output:   output,
		Data:     data,
		Header:   header,
		User:     user,
	}
	a.Run(&options)
}
