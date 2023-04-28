package main

import (
	"bytes"
	"fmt"
	"io"
	"math"
)

type HeapPage struct {
	pid          PageID
	td           *TupleDesc
	numSlots     int
	header       *BitMap
	tuples       []*Tuple
	lastDirtyTid TransactionID
}

func NewHeapPage(id PageID, data []byte) Page {

	// TODO which is better? all empty or 0 but have length
	if len(data) == 0 {
		return nil
	}

	byteReader := bytes.NewReader(data)

	// TODO
	td, err := db.getCatalog().getTupleDesc(id.getTableID())
	if err != nil {
		check(err)
	}
	hp := &HeapPage{pid: id, td: td}
	// byte to bit
	hp.header = NewBitMap(hp.getHeaderSize() * 8)
	hp.tuples = make([]*Tuple, hp.getNumTuples())

	// read header
	byteReader.Read(hp.header.bits)

	// read tuple
	for idx := 0; idx < hp.getNumTuples(); idx++ {
		hp.tuples[idx] = hp.readTuple(byteReader, idx)
	}

	return hp
}

func (hp HeapPage) getID() PageID {
	return hp.pid
}

func (hp HeapPage) getNumTuples() int {
	return PAGESIZE * 8 / (hp.td.getSize()*8 + 1)
}

func (hp HeapPage) getHeaderSize() int {
	return int(math.Ceil(((float64)(hp.getNumTuples())) / 8.0))
}

func (hp HeapPage) insertTuple(t *Tuple) {

	if t == nil {
		check(fmt.Errorf("Tuple is nil"))
	}

	if !hp.td.equals(*t.getTupleDesc()) {
		check(fmt.Errorf("Tuple Desc is mismatch"))
	}

	// first empty slot
	emptySlot := 0
	for emptySlot = 0; hp.isSlotUsed(emptySlot); emptySlot++ {
	}

	hp.tuples[emptySlot] = t
	hp.markSlotUsed(emptySlot)

	recordID := &RecordID{hp.pid, emptySlot}
	t.setRecordID(recordID)
}

func (hp HeapPage) deleteTuple(t Tuple) {

	if !t.getRecordID().getPageID().equals(hp.pid) {
		check(fmt.Errorf("Delete Tuple from wrong the page"))
	}

	locSlot := t.getRecordID().getTupleNo()
	if !hp.isSlotUsed(locSlot) {
		check(fmt.Errorf("Slot is already empty"))
	}
	hp.tuples[locSlot] = &Tuple{}
	hp.markSlotUsed(locSlot)
}

func (hp HeapPage) readTuple(dataStream io.Reader, slotID int) *Tuple {

	RecordID := RecordID{Pid: hp.getID(), TupleNo: slotID}
	tuple := NewTuple(hp.td)
	tuple.setRecordID(&RecordID)
	for idx := 0; idx < tuple.getTupleDesc().numFields(); idx++ {
		f := hp.td.getFieldType(idx)

		switch f := f.(type) {
		case INT_TYPE:
			intField := f.parse(dataStream).(IntField)
			tuple.setField(idx, intField)
		}
	}
	return tuple
}

func createEmptyPageData() []byte {
	return make([]byte, PAGESIZE)
}

func (hp HeapPage) getNumEmptySlots() int {
	numEmptySlot := 0
	for idx := 0; idx < hp.header.getSize(); idx++ {
		if !hp.header.Check(idx) {
			numEmptySlot++
		}
	}
	return numEmptySlot
}

func (hp HeapPage) isSlotUsed(idx int) bool {
	return hp.header.Check(idx)
}

func (hp HeapPage) markSlotUsed(idx int) {
	hp.header.Set(idx)
}

func (hp HeapPage) getPageData() []byte {
	buf := bytes.Buffer{}

	buf.Write(hp.header.bits)

	for idx, v := range hp.tuples {

		if hp.isSlotUsed(idx) {

			for idx := 0; idx < v.Td.numFields(); idx++ {
				f := v.getField(idx)
				f.serialize(&buf)
			}
		} else {
			for idx := 0; idx < v.Td.getSize(); idx++ {
				buf.WriteByte(0)
			}
		}
	}

	//padding
	remainBytes := PAGESIZE - buf.Len()
	for idx := 0; idx < remainBytes; idx++ {
		buf.WriteByte(0)
	}

	return buf.Bytes()
}

func (hp HeapPage) markDirty(dirty bool, tid TransactionID) {

	if dirty {
		hp.lastDirtyTid = tid
	} else {
		hp.lastDirtyTid = nil
	}
}

func (hp HeapPage) isDirty() TransactionID {
	return hp.lastDirtyTid
}

func (hp HeapPage) getHeader() *BitMap {
	return hp.header
}
