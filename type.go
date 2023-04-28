package main

import (
	"encoding/binary"
	"io"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

type Type interface {
	getLen() int
	parse(io.Reader) Field
}

const INT_TYPE_LEN = 8

type INT_TYPE struct {
}

func (i INT_TYPE) getLen() int {
	return INT_TYPE_LEN

}

func (i INT_TYPE) parse(reader io.Reader) Field {
	intBytes := make([]byte, INT_TYPE_LEN)
	_, err := reader.Read(intBytes)
	check(err)
	return NewIntField(int(binary.LittleEndian.Uint64(intBytes)))
}
