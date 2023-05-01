package main

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_TupleInsertDownToFileAndReadBack(t *testing.T) {

	bp := db.getBufferPool()

	tableID := db.getCatalog().getTableID("hello_DB")
	tid := 0

	td, err := db.catalog.getTupleDesc(tableID)

	if err != nil {
		check(err)
	}
	tp := NewTuple(td)
	tp.setField(0, NewIntField(111))
	tp.setField(1, NewIntField(111))
	tp.setField(2, NewIntField(111))

	bp.insertTuple(tid, tableID, tp)
	tp2 := NewTuple(td)
	tp2.setField(0, NewIntField(222))
	tp2.setField(1, NewIntField(222))
	tp2.setField(2, NewIntField(222))
	bp.insertTuple(tid, tableID, tp2)

	// flush the page
	bp.flushPage(tp.getRecordID().getPageID())

	// check from file
	byteSlice, err := os.ReadFile("/home/backman/GODB/DATA/hello_DB.data")
	check(err)
	bReader := bytes.NewReader(byteSlice)
	bp2 := NewHeapPage(tp.getRecordID().getPageID(), byteSlice)

	skipBytes := make([]byte, bp2.getHeaderSize())
	bReader.Read(skipBytes)

	tpOut := bp2.readTuple(bReader, tp.getRecordID().getTupleNo())

	assert.Equal(t, tp.isEqual(*tpOut), true)
}
