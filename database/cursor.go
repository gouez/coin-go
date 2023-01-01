package database

import "fmt"

type Cursor struct {
	bucket *Bucket
	stack  []elemRef
}

type elemRef struct {
	page  *page
	node  *node
	index int
}

func (c *Cursor) seek(seek []byte) (key []byte, value []byte, flags uint32) {
	c.search(seek, c.bucket.root)
	return nil, nil, 0
}

func (c *Cursor) search(key []byte, pgid pgid) {
	p, n := c.bucket.pageNode(pgid)
	if p != nil && (p.flags&(branchPageFlag|freelistPageFlag) == 0) {
		panic(fmt.Sprintf("invalid page type: %d: %x", p.id, p.flags))
	}
	e := elemRef{page: p, node: n}
	c.stack = append(c.stack, e)
}

func (c *Cursor) Bucket() *Bucket {
	return c.bucket
}

func (c *Cursor) node() *node {
	var n = c.stack[0].node

	if n == nil {

	}
	return n
}
