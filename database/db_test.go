package database

import (
	"fmt"
	"testing"
)

func TestOpen(t *testing.T) {
	db, _ := Open("./data.db")
	// db.Put([]byte("cccc"), []byte("dddd"))
	c, _ := db.Get([]byte("cccc"))
	fmt.Printf("%s\n", c)
	// os.Remove("./data.db")
	// os.Remove("./index.db")
}
