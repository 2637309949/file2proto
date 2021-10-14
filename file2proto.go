package main

import (
	"fmt"
	"os"
	"text/template"
)

var transforms = map[string]file2proto{}

type file2proto interface {
	Check(string) bool
	Transform(string) ([]*message, error)
}

func register(name string, transform file2proto) {
	transforms[name] = transform
}

func transform(uri string) (f2p []file2proto) {
	for _, tf := range transforms {
		if tf.Check(uri) {
			f2p = append(f2p, tf)
		}
	}
	return
}

func buildFile2proto(uri string, outputFile string) error {
	tf := transform(uri)
	m, err := buildMessage(uri, tf...)
	if err != nil {
		panic(err)
	}
	err = writeOutput(m, outputFile)
	if err != nil {
		panic(err)
	}
	return nil
}

func buildMessage(uri string, f2t ...file2proto) (items []*message, err error) {
	for _, ft := range f2t {
		m, err := ft.Transform(uri)
		if err != nil {
			return []*message{}, err
		}
		items = append(items, m...)
	}
	return
}

func writeOutput(msgs []*message, outputFile string) error {
	msgTemplate := `syntax = "proto3";
package proto;

{{range .}}
message {{.Name}} {
{{- range .Fields}}
{{- if .IsRepeated}}
  repeated {{.TypeName}} {{.Name}} = {{.Order}};
{{- else}}
  {{.TypeName}} {{.Name}} = {{.Order}} [json_name = "{{.Name}}"];
{{- end}}
{{- end}}
}
{{end}}
`
	tmpl, err := template.New("test").Parse(msgTemplate)
	if err != nil {
		return err
	}

	f, err := os.Create(outputFile)
	if err != nil {
		return fmt.Errorf("unable to create file %s : %s", outputFile, err)
	}
	defer f.Close()

	return tmpl.Execute(f, msgs)
}
