package main

import (
	"testing"
)

func TestInsertTuple(t *testing.T) {

	/*
	   pageID := HeapPageID{0, 0}
	   data := make([]byte, PAGESIZE)
	   hp := NewHeapPage(pageID, data)

	   tuple := NewTuple(TMPTUPLEDESC)
	   tuple.setField(0, NewIntField(int(0x12)))
	   hp.insertTuple(tuple)

	   pBytes := hp.getPageData()

	   buf := bytes.Buffer{}

	   enc := gob.NewEncoder(&buf)
	   enc.Encode(int(0x12))

	   expectBytes := []byte{0x01, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	   expectBytes = append(expectBytes, buf.Bytes()...)

	   remainBytesNum := PAGESIZE - len(expectBytes)
	   remainBytes := make([]byte, remainBytesNum)

	   expectBytes = append(expectBytes, remainBytes...)
	   assert.EqualValues(t, expectBytes, pBytes)
	*/
}
