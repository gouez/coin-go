package database

import (
	"unsafe"
)

const (
	M_PUT = iota
	M_DEL
)

const (
	rheaderSize  = unsafe.Sizeof(record{})
	rdataMaxSize = 1<<31 - 1
)

// header
type record struct {
	ksz  uint32
	vsz  uint32
	mark uint16
}

// data
type rdata struct {
	k []byte
	v []byte
}

func newRecord(k []byte, v []byte, mark uint16) *record {
	r := &record{
		mark: mark,
		ksz:  uint32(len(k)),
		vsz:  uint32(len(v)),
	}
	data := (*[rdataMaxSize]byte)(unsafe.Pointer(uintptr(unsafe.Pointer(r)) + rheaderSize))[:r.ksz+r.vsz]
	copy(data, append(k, v...))
	return r
}

func decodeRecordHeader(buf []byte) *record {
	return (*record)(unsafe.Pointer(&buf[0]))
}

// r to []byte
func (r *record) write() []byte {
	data := (*(*[rdataMaxSize]byte)(unsafe.Pointer(r)))[:rheaderSize+uintptr(r.ksz)+uintptr(r.vsz)]
	return data
}

func (r *record) key() []byte {
	return (*[rdataMaxSize]byte)(unsafe.Pointer(uintptr(unsafe.Pointer(r)) + rheaderSize))[:r.ksz]
}

func (r *record) value() []byte {
	return (*[rdataMaxSize]byte)(unsafe.Pointer(uintptr(unsafe.Pointer(r)) + rheaderSize))[r.ksz : r.ksz+r.vsz]
}
