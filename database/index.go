package database

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"io"
	"os"

	"github.com/gouez/coin-go/lib/logx"
)

const (
	//default index file name
	IndexFileName   = "index.db"
	IndexHeaderSize = 12
)

// memory struct for index
type index struct {
	indexes map[string]int64
	file    *os.File
	offset  int64
}

// disk struct for index
type indexEle struct {
	ksz    uint32
	offset int64
	key    []byte
}

// newIndex create a index instance , only one
func newIndex() *index {
	i := &index{
		indexes: make(map[string]int64),
	}
	i.init()
	return i
}

// init init index
// create index file if not existed
func (i *index) init() {
	f, err := os.OpenFile(IndexFileName, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	info, err := f.Stat()
	if err != nil {
		panic(err)
	}
	i.file = f
	i.offset = info.Size()
	//load index file
	i.loadIndex(*bufio.NewReader(f))
}

// sync sync an index to disk
func (i *index) sync(key []byte, offset int64) error {
	ele := &indexEle{
		ksz:    uint32(len(key)),
		key:    key,
		offset: offset,
	}
	buf := &bytes.Buffer{}
	fn := func(b *bytes.Buffer, element ...any) error {
		for _, e := range element {
			err := binary.Write(b, binary.LittleEndian, e)
			if err != nil {
				return err
			}
		}
		return nil
	}
	err := fn(buf, ele.ksz, ele.offset, ele.key)
	if err != nil {
		logx.Error(err.Error())
		return err
	}
	_, err = i.file.WriteAt(buf.Bytes(), i.offset)
	if err != nil {
		return err
	}
	i.offset += int64(len(buf.Bytes()))
	return nil
}

// loadIndex load index
func (i *index) loadIndex(read bufio.Reader) (error, map[string]int64) {
	buf := make([]byte, 1024)
	m := make(map[string]int64, 0)
	var (
		n   int
		err error
		// htmp   []byte
		// vtmp   []byte
		offset int64
	)
	for {
		if n > 0 {
			for {
				ksz := binary.LittleEndian.Uint32(buf[offset : 4+offset])
				koffset := int64(binary.LittleEndian.Uint64(buf[4+offset : 12]))
				offset += IndexHeaderSize
				key := string(buf[offset : offset+int64(ksz)])
				m[key] = koffset
				n -= int(IndexHeaderSize + ksz)

				if n <= 0 {
					break
				}
			}

		}
		n, err = read.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err, nil
		}
	}
	return err, nil
}

// createIndex create a index for key
func (i *index) createIndex(key []byte, offset int64) {
	i.sync(key, offset)
	i.indexes[string(key)] = offset
}

// getIndex get index from memory
func (i *index) getIndex(key string) int64 {
	return i.indexes[string(key)]
}
