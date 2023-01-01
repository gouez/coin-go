package database

import "unsafe"

type (
	pages []*page
	pgid  uint64
	pgids []pgid
)

const (
	magic        uint32 = 0x85278527
	maxAllocSize        = 0x7FFFFFFF
)

const (
	branchPageFlag        = 0x01
	leafPageFlag          = 0x02
	metaPageFlag          = 0x04
	freelistPageFlag      = 0x10
	pageSize              = unsafe.Sizeof(page{})
	leafPageElementSize   = unsafe.Sizeof(leafPageElement{})
	branchPageElementSize = unsafe.Sizeof(leafPageElement{})
)

// Physical structure
type page struct {
	id       pgid
	flags    uint16
	count    uint16
	overflow uint32
}

// meta page
// Physical structure
type meta struct {
	magic uint32
	root  bucket
}

// Physical structure
type branchPageElement struct {
	pos   uint32
	ksize uint32
	pgid  pgid
}

// Physical structure
type leafPageElement struct {
	flags uint32
	pos   uint32
	ksize uint32
	vsize uint32
}

func (pg *page) leafPageElement(index uint16) *leafPageElement {
	return (*leafPageElement)(unsafe.Pointer(uintptr(unsafe.Pointer(pg)) + pageSize + uintptr(index)*leafPageElementSize))
}

func (pg *page) branchPageElement(index uint16) *branchPageElement {
	return (*branchPageElement)(unsafe.Pointer(uintptr(unsafe.Pointer(pg)) + pageSize + uintptr(index)*branchPageElementSize))
}

func (pg *page) meta() *meta {
	return (*meta)(unsafe.Pointer(uintptr(unsafe.Pointer(pg)) + unsafe.Sizeof(*pg)))
}

func (s pgids) Len() int           { return len(s) }
func (s pgids) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s pgids) Less(i, j int) bool { return s[i] < s[j] }

func (s pages) Len() int           { return len(s) }
func (s pages) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s pages) Less(i, j int) bool { return s[i].id < s[j].id }
