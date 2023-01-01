package database

import (
	"fmt"
	"testing"
)

func TestOpen(t *testing.T) {
	db, _ := Open("./db.db", 0666)
	fmt.Println(db.page(0).meta().magic)
}

func TestDB_Update(t *testing.T) {
	db, _ := Open("./db.db", 0666)
	db.Update(func(tx *Tx) error {
		tx.CreateBucket([]byte("user"))
		return nil
	})
}
