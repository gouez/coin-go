package blockchain

import "github.com/gouez/coin-go/lib/chainhash"

const dbPath = "./tmp/blocks"

// BlockChain stores all block
type BlockChain struct {
	blocks []*Block
}

func NewBlcokChain() *BlockChain {
	genesisBlock := GenesisBlock()
	chain := &BlockChain{
		blocks: []*Block{genesisBlock},
	}
	return chain
}

// AddBlock add a block to blcokchain
func (chain *BlockChain) AddBlock(data []byte) {
	if len(chain.blocks) == 0 {
		return
	}
	lastBlockHash := chain.blocks[len(chain.blocks)-1].Hash
	block := NewBlock(lastBlockHash, data)
	chain.blocks = append(chain.blocks, block)
}

// GenesisBlock first block
func GenesisBlock() *Block {
	b := NewBlock(chainhash.Hash{}, []byte("GenesisBlock"))
	return b
}
