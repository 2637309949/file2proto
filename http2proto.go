package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type http2proto struct {
}

func (to *http2proto) Transform(uri string) ([]*message, error) {
	pkgs, err := loadHTTP(uri)
	if err != nil {
		log.Fatalf("error loadHTTP: %s", err)
		return []*message{}, err
	}
	msgs := rspInterToMessages("Body", pkgs)
	return msgs, nil
}

func (to *http2proto) Check(uri string) bool {
	if strings.Contains(uri, "http://") || strings.Contains(uri, "https://") {
		return true
	}
	return false
}

func loadHTTP(uri string) (map[string]interface{}, error) {
	res, err := http.Get(uri)
	if err != nil {
		return nil, err
	}
	rspBytes, err := ioutil.ReadAll(res.Body)
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

func rspInterToMessages(name string, objs map[string]interface{}) (out []*message) {
	msg := message{
		Name:   name,
		Fields: []*field{},
	}
	i := 0
	for k, v := range objs {
		newField := &field{
			Name:       toProtoFieldName(k),
			IsRepeated: false,
			Order:      i + 1,
		}
		switch v.(type) {
		case int64, int32, int:
			newField.TypeName = "int64"
		case float64, float32:
			newField.TypeName = "float64"
		case string:
			newField.TypeName = "string"
		case bool:
			newField.TypeName = "bool"
		default:
			rspBytes, err := json.Marshal(v)
			if err != nil || string(rspBytes) == "{}" || string(rspBytes) == "[]" || string(rspBytes) == "null" {
				continue
			}
			if strings.HasPrefix(string(rspBytes), "{") {
				rspInter := map[string]interface{}{}
				err = json.Unmarshal(rspBytes, &rspInter)
				if err != nil || rspInter == nil {
					continue
				}
				newField.TypeName = case2camel(name + "_" + k)
				out = append(out, rspInterToMessages(newField.TypeName, rspInter)...)
			} else if strings.HasPrefix(string(rspBytes), "[") {
				rspInter := []map[string]interface{}{}
				err = json.Unmarshal(rspBytes, &rspInter)
				if err != nil || rspInter == nil {
					continue
				}
				newField.TypeName = case2camel(name + "_" + k)
				out = append(out, rspInterToMessages(newField.TypeName, rspInter[0])...)
			}
		}
		msg.Fields = append(msg.Fields, newField)
	}
	out = append(out, &msg)
	return
}

func init() {
	register("http2proto", new(http2proto))
}
