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

func (hf heapFile) getIterator(tid TransactionID) DBFileIterator {
	return hf.NewHeapFileIterator(tid)
}

type heapFileIterator struct {
	tid     TransactionID
	curItr  TupleIterator
	curPage int
	file    *heapFile
}

func (hf heapFile) NewHeapFileIterator(tid TransactionID) DBFileIterator {
	return &heapFileIterator{tid: tid, curPage: 0, file: &hf}
}

func (hfIter *heapFileIterator) open() {
	hfIter.curPage = 0

	if hfIter.curPage >= hfIter.file.getNumPage() {
		return
	}

	hfIter.curItr.tuples = db.getBufferPool().getPage(hfIter.tid, HeapPageID{tableID: hfIter.file.getID(), pageNo: hfIter.curPage}, READ_ONLY).iterator()
	hfIter.curItr.idx = 0
}

func (hfIter *heapFileIterator) advance() {

	for !hfIter.curItr.hasNext() {
		hfIter.curPage++
		if hfIter.curPage < hfIter.file.getNumPage() {
			hfIter.curItr.tuples = db.getBufferPool().getPage(hfIter.tid, HeapPageID{tableID: hfIter.file.getID(), pageNo: hfIter.curPage}, READ_ONLY).iterator()
		} else {
			break
		}
	}

}

func (hfIter heapFileIterator) hasNext() bool {
	return hfIter.curPage < hfIter.file.getNumPage()
}

func (hfIter *heapFileIterator) next() Tuple {

	if !hfIter.hasNext() {
		return Tuple{}
	}
	tuple := hfIter.curItr.next()
	hfIter.advance()
	return tuple
}

func (hfIter *heapFileIterator) close() {
	hfIter.curItr = TupleIterator{}
	hfIter.curPage = 0
}

func (hfIter heapFileIterator) getTupleDesc() {}
