package database

import "testing"

func TestTx_write(t *testing.T) {
	// _, _ = Open("./db.db", 0666)
	db, _ := Open("./db.db", 0666)
	tx := newTx(db)
	tx.pages = map[pgid]*page{
		2: {
			id: 2,
		},
		3: {
			id: 3,
		},
	}
	tx.write()
}
