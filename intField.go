package main

import (
	"encoding/binary"
	"io"
)

type IntField struct {
	Value int
}

func NewIntField(v int) IntField {
	return IntField{Value: v}
}

func (f IntField) getType() Type {
	return INT_TYPE{}
}

func (f IntField) getValue() int {
	return f.Value
}

func (f IntField) compare(v interface{}) bool {
	if v == nil {
		return false
	}
	vOther := v.(IntField)
	if f.Value == vOther.getValue() {
		return true
	}
	return false
}

func (f IntField) serialize(w io.Writer) {
	intBytes := make([]byte, f.getType().getLen())
	binary.LittleEndian.PutUint64(intBytes, uint64(f.Value))
	w.Write(intBytes)
}
