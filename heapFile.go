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
	f  *os.File
	td TupleDesc
}

func (hf *heapFile) NewHeapFile(f *os.File) *heapFile {
	return &heapFile{f: f}
}

func (hf heapFile) getFile() *os.File {
	return hf.f
}

func (hf heapFile) getID() int {

	return int(hash(hf.f.Name()))
}

func (hf heapFile) getTupleDesc() TupleDesc {
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

func (hf *heapFile) numPages() int {
	// TODO
	return 0
}

func (hf *heapFile) InsertTuple(tid TransactionID, *t Tuple) Page {

	if t == nil{
		check(fmt.Errorf("Tuple is empty"))
	}

	

	return nil
}

func (hf *heapFile) deleteTuple(tid TransacheapFile, t Tuple) Page {

}
