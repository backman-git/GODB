package main

type PageID interface {
	equals(v interface{}) bool
	getTableID() int
	pageno() int
}
