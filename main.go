package main

import "fmt"

func main() {
	fmt.Println("hi")
}

// TODO
const (
	PAGESIZE = 4096
)

var TMPTUPLEDESC = NewTupleDesc([]Type{INT_TYPE{}}, []string{"integer"})
