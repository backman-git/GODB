package main

type BitMap struct {
	bits []byte
	len  int
}

func NewBitMap(size int) *BitMap {

	bm := &BitMap{}
	bm.len = size

	allocSize := (size + 7) / 8
	bm.bits = make([]byte, allocSize)
	return bm
}

func (bm *BitMap) Set(idx int) {

	if idx > bm.len {
		return
	}
	bm.bits[idx/8] |= 1 << (idx % 8)
}

func (bm *BitMap) UnSet(idx int) {
	if idx > bm.len {
		return
	}

	bm.bits[idx/8] &^= 1 << (idx % 8)
}

func (bm *BitMap) Check(idx int) bool {
	if idx >= bm.len {
		return false
	}
	return bm.bits[idx/8]&(1<<(idx%8)) != 0
}

func (bm *BitMap) getSize() int {
	return bm.len
}
