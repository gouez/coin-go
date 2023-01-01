package database

type Bucket struct {
	*bucket
	tx       *Tx
	rootNode *node
}

type bucket struct {
	root pgid
}

func newBucket(tx *Tx) Bucket {
	var b = Bucket{tx: tx}
	return b
}

// CreateBucket create bucket,return error if it is existed or other error
func (b *Bucket) CreateBucket(key []byte) error {
	c := b.Cursor()
	//seek(key)
	_, _, _ = c.seek(key)
	//create new bucket
	var _ = Bucket{
		rootNode: &node{
			isLeaf: true,
		},
	}

	return nil
}

func (b *Bucket) pageNode(id pgid) (*page, *node) {
	return b.tx.page(id), nil
}

func (b *Bucket) DelBucket(key []byte) error {

	return nil
}

func (b *Bucket) Cursor() *Cursor {
	return &Cursor{
		bucket: b,
	}
}
