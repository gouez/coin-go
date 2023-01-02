package database

const (
	//default index file name
	IndexFileName = "index.db"
)

// memory struct for index
type index struct {
	indexes map[string]int64
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
	return i
}

// sync sync an index to disk
func (i *index) sync(ele *indexEle) {
	
}

// createIndex create a index for key
func (i *index) createIndex(key string, offset int64) {
	i.indexes[string(key)] = offset
}

// getIndex get index from memory
func (i *index) getIndex(key string) int64 {
	return i.indexes[string(key)]
}
