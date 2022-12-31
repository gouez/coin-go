package database

import "unsafe"

type pgid uint64

const (
	magic uint32 = 0x85278527
)

const (
	branchPageFlag   = 0x01
	leafPageFlag     = 0x02
	metaPageFlag     = 0x04
	freelistPageFlag = 0x10
)

type page struct {
	id       pgid
	flags    uint16
	count    uint16
	overflow uint32
}

// meta page
type meta struct {
	magic uint32
	root  bucket
}

func (pg *page) meta() *meta {
	return (*meta)(unsafe.Pointer(uintptr(unsafe.Pointer(pg)) + unsafe.Sizeof(*pg)))
}
