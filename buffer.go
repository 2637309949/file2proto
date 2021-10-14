package main

import (
	"bytes"
	"strconv"
)

type buffer struct {
	*bytes.Buffer
}

func newbuffer() *buffer {
	return &buffer{Buffer: new(bytes.Buffer)}
}

func (b *buffer) Append(i interface{}) *buffer {
	switch val := i.(type) {
	case int:
		b.append(strconv.Itoa(val))
	case int64:
		b.append(strconv.FormatInt(val, 10))
	case uint:
		b.append(strconv.FormatUint(uint64(val), 10))
	case uint64:
		b.append(strconv.FormatUint(val, 10))
	case string:
		b.append(val)
	case []byte:
		b.Write(val)
	case rune:
		b.WriteRune(val)
	}
	return b
}

func (b *buffer) append(s string) *buffer {
	b.WriteString(s)
	return b
}
