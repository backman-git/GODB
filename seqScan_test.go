package main

import (
	"fmt"
	"testing"
)

func TestSeqScan(t *testing.T) {

	tableID := db.getCatalog().getTableID("hello_DB")
	seqScan := NewSeqScan(0, tableID)

	seqScan.open()
	for seqScan.hasNext() {
		tuple := seqScan.next()
		fmt.Println(tuple)
	}

}
