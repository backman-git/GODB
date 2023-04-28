package main

import (
	"fmt"
	"reflect"
)

type Tuple struct {
	Td     *TupleDesc
	RID    *RecordID
	Fields []Field
}

func NewTuple(Td *TupleDesc) *Tuple {
	fieldNum := Td.numFields()
	return &Tuple{Td: Td, Fields: make([]Field, fieldNum)}
}

func (t Tuple) getTupleDesc() *TupleDesc {
	return t.Td
}

func (t Tuple) getRecordID() *RecordID {
	return t.RID
}

func (t *Tuple) setRecordID(rID *RecordID) {
	t.RID = rID
}

func (t *Tuple) setField(idx int, f Field) {
	if idx >= t.Td.numFields() {
		check(fmt.Errorf("The idx of field is over the numbers of the tuple"))
	}
	if f.getType() != t.Td.getFieldType(idx) {
		check(fmt.Errorf("mismatch type"))
	}

	t.Fields[idx] = f
}

func (t Tuple) getField(idx int) Field {
	if idx >= t.Td.numFields() {
		check(fmt.Errorf("The idx of field is over the numbers of the tuple"))
	}
	return t.Fields[idx]

}

func (t Tuple) isEqual(other Tuple) bool {
	return reflect.DeepEqual(t, other)
}
