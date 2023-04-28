package main

import (
	"fmt"
	"hash/fnv"
	"os"
)

func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

type heapFile struct {
	f        *os.File
	td       *TupleDesc
	numPages int
}

func NewHeapFile(f *os.File, td *TupleDesc) *heapFile {
	fState, err := f.Stat()
	check(err)
	numPages := fState.Size() / PAGESIZE
	return &heapFile{f: f, td: td, numPages: int(numPages)}
}

func (hf heapFile) insertTuple(tid TransactionID, t *Tuple) Page {

	if t.getTupleDesc() == nil {
		check(fmt.Errorf("Insert Empty Tuple"))
	}

	var page *HeapPage
	for idx := 0; idx <= hf.numPages; idx++ {
		pid := HeapPageID{tableID: hf.getID(), pageNo: idx}
		page = db.getBufferPool().getPage(tid, pid, READ_WRITE).(*HeapPage)

		if page.getNumEmptySlots() > 0 {
			page.insertTuple(t)
			break
		}

	}
	// page become dirty
	return page

}

func (hf heapFile) getFile() *os.File {
	return hf.f
}

func (hf heapFile) getID() int {

	return int(hash(hf.f.Name()))
}

func (hf heapFile) getTupleDesc() *TupleDesc {
	return hf.td
}

func (hf heapFile) readPage(pid PageID) Page {

	if pid == nil {
		return nil
	}

	pageNum := pid.pageno()
	fileOffset := pageNum * PAGESIZE
	pageData := make([]byte, PAGESIZE)

	_, err := hf.f.Seek(int64(fileOffset), 0)
	check(err)
	_, err = hf.f.Read(pageData)
	check(err)

	return NewHeapPage(pid, pageData)

}

func (hf *heapFile) writePage(p Page) {
	fileOffset := p.getID().pageno() * PAGESIZE
	_, err := hf.f.Seek(int64(fileOffset), 0)
	check(err)
	hf.f.Write(p.getPageData())
	check(err)
}

func (hf *heapFile) deleteTuple(tid TransactionID, t *Tuple) Page {
	return nil
}

func (hf heapFile) getNumPage() int {
	return hf.numPages
}

func (hf heapFile) getIterator() DBIterator {
	//return NewHeapFileIterator(hf.tid)
}

type heapFileIterator struct {
	tid       TransactionID
	tupleIter TupleIterator
	curPage   int
	file      heapFile
}

func NewHeapFileIterator(tid TransactionID) *heapFileIterator {

	return &heapFileIterator{tid: tid, curPage: 0}
}

func (hfIter *heapFileIterator) open() {
	hfIter.curPage = 0
	pageID := HeapPageID{}
	db.getBufferPool().getPage(hfIter.tid)

}
