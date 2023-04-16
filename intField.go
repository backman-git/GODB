package main

import (
	"encoding/gob"
	"io"
	"log"
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
	enc := gob.NewEncoder(w)
	if err := enc.Encode(f.Value); err != nil {
		log.Fatal("Encode Error")
	}
}
