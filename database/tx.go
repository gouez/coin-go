package database

import (
	"sort"
	"unsafe"
)

type Tx struct {
	root  Bucket
	db    *DB
	pages map[pgid]*page
}

func newTx(db *DB) *Tx {
	tx := &Tx{
		db: db,
	}
	// tx.root = newBucket(tx)
	// *tx.root.bucket = db.meta.root
	return tx
}

// write writes any dirty pages to disk
func (tx *Tx) write() error {
	pages := make(pages, 0, len(tx.pages))
	for _, p := range tx.pages {
		pages = append(pages, p)
	}
	tx.pages = make(map[pgid]*page)
	sort.Sort(pages)
	// Write pages to disk in order.
	for _, p := range pages {
		size := int(p.overflow+1) * tx.db.pagesz
		offset := int64(p.id) * int64(tx.db.pagesz)
		ptr := (*[maxAllocSize]byte)(unsafe.Pointer(p))
		for {
			sz := size
			if sz > maxAllocSize-1 {
				sz = maxAllocSize - 1
			}
			buf := ptr[:sz]
			if _, err := tx.db.ops.writeAt(buf, offset); err != nil {
				return err
			}
			size -= sz
			if size == 0 {
				break
			}
		}
	}
	return nil
}

func (tx *Tx) CreateBucket(key []byte) error {
	return tx.root.CreateBucket(key)
}
func (tx *Tx) page(id pgid) *page {
	return tx.db.page(id)
}
