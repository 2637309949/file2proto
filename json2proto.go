package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type json2proto struct {
}

func (to *json2proto) Transform(uri string) ([]*message, error) {
	pkgs, err := loadJSON(uri)
	if err != nil {
		log.Fatalf("error loadHTTP: %s", err)
		return []*message{}, err
	}
	msgs := rspInterToMessages("Body", pkgs)
	return msgs, nil
}

func (to *json2proto) Check(uri string) bool {
	if strings.Contains(uri, ".json") {
		return true
	}
	s, err := os.Stat(uri)
	if err == nil && s.IsDir() {
		files, _ := ioutil.ReadDir(uri)
		for _, f := range files {
			if strings.Contains(f.Name(), ".json") {
				return true
			}
		}
	}
	return false
}

func loadJSON(uri string) (map[string]interface{}, error) {
	rspBytes, err := ioutil.ReadFile(uri)
	if err != nil {
		return nil, err
	}
	rspInter := map[string]interface{}{}
	err = json.Unmarshal(rspBytes, &rspInter)
	if err != nil {
		return nil, err
	}
	return rspInter, nil
}

func init() {
	register("json2proto", new(json2proto))
}
