/*
A "hello world" plugin in Go,
which reads a request header and sets a response header.
*/
package main

import (
	"github.com/Kong/go-pdk"
)

type Config struct {
	Message string
}

func New() interface{} {
	return &Config{}
}

func (conf Config) Access(kong *pdk.PDK) {
	message := conf.Message
	if message == "" {
		message = "hello"
	}
	kong.Response.Exit(200, "{\"message\": "+message+"}", map[string][]string{"Content-type": []string{"application/json"}})
}
