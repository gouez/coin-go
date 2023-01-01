package database

type (
	nodes  []*node
	inodes []inode
)

type node struct {
	//memory structure
	pgid     pgid
	parent   *node
	children nodes
	inodes   inodes
	//Physical structure
	isLeaf bool
}
type inode struct {
	flags uint32
	pgid  pgid
	key   []byte
	value []byte
}

func (n *node) root() *node {
	if n.parent == nil {
		return n
	}
	return n.parent.root()
}

// read convert page to node
func (n *node) read(p *page) {
	n.pgid = p.id
}

// write convert node to page
func (n *node) write(p *page) {

}
