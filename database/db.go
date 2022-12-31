package database

import (
	"fmt"
	"os"
	"syscall"
	"unsafe"
)

const maxMapSize = 0xFFFFFFFFFFFF

type DB struct {
	pagesz  int
	file    os.File
	path    string
	data    *[maxMapSize]byte
	dataref []byte
	datasz  int
}

// Open open db file
func Open(path string, mode os.FileMode) (*DB, error) {
	db := &DB{
		pagesz: os.Getpagesize(),
	}
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, mode)
	if err != nil {
		return nil, err
	}
	db.file = *file
	db.path = file.Name()
	if finfo, err := file.Stat(); err != nil {
		return nil, err
	} else if finfo.Size() == 0 {
		// init db file
		db.init()
	} else {
		//TODO check file
	}
	if err := db.mmap(); err != nil {
		return nil, err
	}
	return db, nil
}

// init init db file
func (db *DB) init() {
	buf := make([]byte, db.pagesz*2)
	pg := db.pageInbuffer(buf, 0)
	pg.id = 0
	m := pg.meta()
	m.magic = magic
	m.root = bucket{root: 1}

	pg = db.pageInbuffer(buf[:], pgid(1))
	pg.flags = leafPageFlag
	pg.id = 1
	pg.count = 0

	db.file.WriteAt(buf, 0)
}

func (db *DB) pageInbuffer(buf []byte, id pgid) *page {
	return (*page)(unsafe.Pointer(&buf[id*pgid(db.pagesz)]))
}

func (db *DB) mmap() error {
	info, err := db.file.Stat()
	if err != nil {
		return fmt.Errorf("mmap stat error: %s", err)
	}
	if err := db.munmap(); err != nil {
		return err
	}
	sz := info.Size()
	b, err := syscall.Mmap(int(db.file.Fd()), 0, int(sz), syscall.PROT_READ, syscall.MAP_SHARED)
	if err != nil {
		return err
	}

	err = syscall.Madvise(b, syscall.MADV_RANDOM)
	if err != nil && err != syscall.ENOSYS {
		return fmt.Errorf("madvise: %s", err)
	}

	db.dataref = b
	db.data = (*[maxMapSize]byte)(unsafe.Pointer(&b[0]))
	db.datasz = int(sz)

	return nil
}
func (db *DB) munmap() error {
	if db.dataref == nil {
		return nil
	}
	err := syscall.Munmap(db.dataref)
	db.dataref = nil
	db.data = nil
	db.datasz = 0
	return err
}
func (db *DB) page(id pgid) *page {
	return (*page)(unsafe.Pointer(&db.data[id*pgid(db.pagesz)]))
}
