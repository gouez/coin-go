package database

type nodes []*node

type node struct {
	pgid     pgid
	parent   *node
	children nodes
}

func (n *node) root() *node {
	if n.parent == nil {
		return n
	}
	return n.parent.root()
}
