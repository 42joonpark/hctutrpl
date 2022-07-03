package app

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

type App struct {
	Url    string
	Method string
}

func (a *App) Run() {
	u, err := url.ParseRequestURI(a.Url)
	if err != nil {
		log.Fatal(err)
	}

	var req = &http.Request{
		Method: a.Method,
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
