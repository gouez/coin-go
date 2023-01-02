package database

import (
	"os"
	"syscall"
	"unsafe"
)

const (
	MaxMMapSize = 1 << 30
)

type dbfile struct {
	file   *os.File
	offset int64
	//file data
	data    *[MaxMMapSize]byte
	dataRef []byte
}

func newDBFile(name string, mode os.FileMode) (*dbfile, error) {
	f, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE, mode)
	if err != nil {
		return nil, err
	}
	info, err := f.Stat()
	if err != nil {
		return nil, err
	}
	df := &dbfile{
		file:   f,
		offset: info.Size(),
	}
	// open mmap
	df.mmap()
	return df, nil
}

func (f *dbfile) mmap() error {
	buf, err := syscall.Mmap(int(f.file.Fd()), 0, 1<<30, syscall.PROT_READ|syscall.PROT_WRITE, syscall.MAP_SHARED)
	if err != nil {
		return err
	}
	f.dataRef = buf
	f.data = (*[MaxMMapSize]byte)(unsafe.Pointer(&buf[0]))
	f.grow(1024)
	return nil
}

// MAYBE NOT USE IT
func (f *dbfile) write(r *record) (err error) {
	buf := r.write()
	_, err = f.file.WriteAt(buf, f.offset)
	if err != nil {
		return err
	}
	f.offset += int64(len(buf))
	return
}

// writemmap use mmap system call to write file
func (f *dbfile) writemmap(r *record) error {
	buf := r.write()
	copy(f.data[f.offset:], buf)
	f.offset += int64(len(buf))
	return nil
}

func (f *dbfile) ummap() error {
	err := syscall.Munmap(f.dataRef)
	if err == nil {
		return err
	}
	f.data = nil
	f.dataRef = nil
	return nil
}

func (f *dbfile) read(offset int64) (r *record, err error) {
	buf := make([]byte, rheaderSize)
	if _, err = f.file.ReadAt(buf, offset); err != nil {
		return
	}
	r = decodeRecordHeader(buf)
	offset += int64(rheaderSize)
	buf = (*[rdataMaxSize]byte)(unsafe.Pointer(uintptr(unsafe.Pointer(r)) + rheaderSize))[:r.ksz+r.vsz]
	if _, err = f.file.ReadAt(buf, offset); err != nil {
		return
	}
	return
}

func (f *dbfile) grow(size int64) error {
	if info, _ := f.file.Stat(); info.Size() >= size {
		return nil
	}
	return f.file.Truncate(size)
}
