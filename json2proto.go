package main

import (
	"io/ioutil"
	"os"
	"strings"
)

type json2proto struct {
}

func (to *json2proto) Transform(uri string) ([]*message, error) {
	return []*message{}, nil
}

func (to *json2proto) Check(uri string) bool {
	if strings.Contains(uri, ".json") {
		return true
	}
	s, _ := os.Stat(uri)
	if s.IsDir() {
		files, _ := ioutil.ReadDir(uri)
		for _, f := range files {
			if strings.Contains(f.Name(), ".json") {
				return true
			}
		}
	}
	return true
}

func init() {
	register("json2proto", new(json2proto))
}
