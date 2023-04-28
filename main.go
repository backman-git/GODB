package main

import "fmt"

func main() {
	fmt.Println("hi")
}

var TMPTUPLEDESC = NewTupleDesc([]Type{INT_TYPE{}}, []string{"integer"})
