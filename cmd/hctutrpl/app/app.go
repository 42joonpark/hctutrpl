package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type userType struct {
	name     string
	password string
}

type OptionFlag struct {
	Method   *string
	UrlInput *string
	Output   *string
	Data     *string
	Header   *string
	User     *string
}

type App struct {
	Url    string
	Method string
	Header http.Header
	u      *url.URL
	res    *http.Response
	body   []byte
}

func (a *App) addHeader(header *string) error {
	var h map[string]interface{}

	err := json.Unmarshal([]byte(*header), &h)
	if err != nil {
		return err
	}

	for key, value := range h {
		a.Header.Add(key, fmt.Sprintf("%v", value))
	}

	return nil
}

func (a *App) addUser(u *string, req *http.Request) error {
	var user userType
	if len(*u) > 0 && strings.Contains(*u, ":") {
		strs := strings.Split(*u, ":")
		user.name = strs[0]
		user.password = strs[1]

		req.SetBasicAuth(user.name, user.password)
	}

	return nil
}

func (a *App) parseRequestUrl() error {
	u, err := url.ParseRequestURI(a.Url)
	if err != nil {
		return err
	}
	a.u = u
	return nil
}

func (a *App) sendRequest(options *OptionFlag) error {
	var req = &http.Request{
		Method: a.Method,
		URL:    a.u,
		Header: a.Header,
	}

	err := a.addUser(options.User, req)
	if err != nil {
		log.Fatal(err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	a.res = res
	return nil
}

func (a *App) readBody() error {
	resBody, err := ioutil.ReadAll(a.res.Body)
	if err != nil {
		return err
	}

	a.body = resBody
	return nil
}

func (a *App) Run(options *OptionFlag) {
	err := a.addHeader(options.Header)
	if err != nil {
		log.Fatal(err)
	}

	err = a.parseRequestUrl()
	if err != nil {
		log.Fatal(err)
	}

	err = a.sendRequest(options)
	if err != nil {
		log.Fatal(err)
	}

	err = a.readBody()
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("client: response body: %s\n", a.body)
}
