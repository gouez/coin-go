package database

import (
	"sync"
)

type DB struct {
	*dbfile
	isOpend bool
	mu      sync.RWMutex
	index   *index
}

// Put put a key/value to db
func (db *DB) Put(key []byte, value []byte) (err error) {
	if len(key) == 0 {
		return
	}
	v, err := db.Get(key)
	if len(v) > 0 {
		return
	}
	db.mu.Lock()
	defer db.mu.Unlock()

	err = db.dbfile.writemmap(newRecord(
		key, value, M_PUT,
	))
	if err != nil {
		return err
	}
	db.index.createIndex(string(key), db.offset)
	return err
}

// Get get value with given key
func (db *DB) Get(key []byte) ([]byte, error) {
	if len(key) == 0 {
		return nil, nil
	}
	db.mu.RLock()
	defer db.mu.RUnlock()

	offset := db.index.getIndex(string(key))
	r, err := db.dbfile.read(offset)
	if err != nil {
		return nil, err
	}
	return r.value(), nil
}

// Open opens a db file,created if it not existed
func Open(name string) (*DB, error) {
	f, err := newDBFile(name, 0666)
	if err != nil {
		return nil, err
	}
	db := &DB{
		dbfile:  f,
		isOpend: true,
		index:   newIndex(),
	}
	return db, nil
}

// Close close db
func (db *DB) Close() {
	if !db.isOpend {
		db.file.Close()
		db.dbfile.ummap()
	}
}
