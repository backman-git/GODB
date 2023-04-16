package main

import "io"

type Field interface {
	getType() Type
	compare(v interface{}) bool
	serialize(io.Writer)
}
