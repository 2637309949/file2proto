package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
	"unicode"
	"unicode/utf8"
)

func ensureOutputFile(dir string) (string, error) {
	targetFile, targetDir := "messages.proto", ""
	lstDir := filepath.Base(dir)
	if strings.Contains(lstDir, ".proto") {
		targetFile = lstDir
		targetDir = filepath.Dir(dir)
	} else {
		targetDir = dir
	}
	_, err := os.Stat(targetDir)
	if os.IsNotExist(err) {
		err = os.MkdirAll(targetDir, os.ModePerm)
		if err != nil {
			return "", err
		}
	}

	filePath := path.Join(targetDir, targetFile)
	_, err = os.Stat(filePath)
	if !os.IsNotExist(err) {
		return "", fmt.Errorf("output file %v already exists", filePath)
	}
	return filePath, nil
}

func camel2case(name string) string {
	buffer := newbuffer()
	for i, r := range name {
		if unicode.IsUpper(r) {
			if i != 0 {
				buffer.Append('_')
			}
			buffer.Append(unicode.ToLower(r))
		} else {
			buffer.Append(r)
		}
	}
	return buffer.String()
}

func toProtoFieldName(name string) string {
	if len(name) == 2 {
		return strings.ToLower(name)
	}
	r, n := utf8.DecodeRuneInString(name)
	name = string(unicode.ToLower(r)) + name[n:]
	return camel2case(name)
}

func normalizeType(name string) string {
	switch name {
	case "int":
		return "int64"
	case "float32":
		return "float"
	case "float64":
		return "double"
	default:
		return name
	}
}
