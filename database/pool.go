package database

import "sync"

type pool struct {
	p *sync.Pool
}

func newPool() pool {
	return pool{p: &sync.Pool{
		New: func() interface{} {
			return &buf{bs: make([]byte, 0, _size)}
		},
	}}
}
func (p pool) get() *buf {
	buf := p.p.Get().(*buf)
	buf.Reset()
	buf.pool = p
	return buf
}

func (p pool) put(buf *buf) {
	p.p.Put(buf)
}
