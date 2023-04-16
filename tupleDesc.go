package main

import "fmt"

type TupleDesc struct {
	FieldDescs []FieldDesc
	TupleSize  int
}

type FieldDesc struct {
	TypeValue Type
	FieldName string
}

func NewTupleDesc(typeAr []Type, fieldAr []string) *TupleDesc {

	fieldDescs := make([]FieldDesc, 0)

	for idx, _ := range fieldAr {
		fieldDescs = append(fieldDescs, FieldDesc{TypeValue: typeAr[idx], FieldName: fieldAr[idx]})
	}

	tupleSize := 0
	for _, v := range fieldDescs {
		tupleSize += v.TypeValue.getLen()
	}
	return &TupleDesc{FieldDescs: fieldDescs, TupleSize: tupleSize}
}

func (td TupleDesc) numFields() int {
	return len(td.FieldDescs)
}

func (td TupleDesc) getFieldName(idx int) string {
	if idx > td.numFields() {
		check(fmt.Errorf("no such field"))
	}

	return td.FieldDescs[idx].FieldName
}

func (td TupleDesc) getFieldType(idx int) Type {
	if idx > td.numFields() {
		check(fmt.Errorf("no such field"))
	}
	return td.FieldDescs[idx].TypeValue
}

func (td TupleDesc) getSize() int {

	return td.TupleSize
}

func (td TupleDesc) equals(tdOther TupleDesc) bool {
	for idx, v := range td.FieldDescs {
		if tdOther.getFieldType(idx) != v.TypeValue {
			return false
		}
	}

	for idx, v := range td.FieldDescs {
		if tdOther.getFieldName(idx) != v.FieldName {
			return false
		}
	}

	return true
}
