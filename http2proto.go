package main

import (
	"strings"
)

type http2proto struct {
}

func (to *http2proto) Transform(uri string) ([]*message, error) {
	return []*message{}, nil
}

func (to *http2proto) Check(uri string) bool {
	if strings.Contains(uri, "http://") || strings.Contains(uri, "https://") {
		return true
	}
	return true
}

func init() {
	register("http2proto", new(http2proto))
}
