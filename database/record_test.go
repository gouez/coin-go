package database

import (
	"fmt"
	"testing"
)

func TestRecord1(t *testing.T) {
	r := newRecord([]byte("a"), []byte("test"), M_PUT)
	fmt.Printf("%s\n", r.key())
	fmt.Printf("%s\n", r.value())
}
