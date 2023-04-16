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

type INT_TYPE struct {
}

func (i INT_TYPE) getLen() int {
	return 4

}

func (i INT_TYPE) parse(reader io.Reader) Field {
	intBytes := make([]byte, 4)
	_, err := reader.Read(intBytes)
	check(err)
	return NewIntField(int(binary.BigEndian.Uint32(intBytes)))
}
