/*
A "hello world" plugin in Go,
which reads a request header and sets a response header.
*/
package main

import (
	"github.com/Kong/go-pdk"
	"github.com/go-resty/resty/v2"
)

type Config struct {
	Url string
}

func New() interface{} {
	return &Config{}
}

func (conf Config) Access(kong *pdk.PDK) {
	Url := conf.Url
	if Url == "" {
		Url = "https://httpbin.org/get"
	}
	client := resty.New()

	resp, _ := client.R().
		EnableTrace().
		Get(Url)
	kong.Response.Exit(200, resp.String(), map[string][]string{"Content-type": []string{"application/json"}})
}
