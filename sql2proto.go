package main

import (
	"strings"
)

type sql2proto struct {
}

func (to *sql2proto) Transform(uri string) ([]*message, error) {
	return []*message{}, nil
}

func (to *sql2proto) Check(uri string) bool {
	if strings.Contains(uri, "mysql://") {
		return true
	}
	return true
}

func init() {
	register("sql2proto", new(sql2proto))
}
