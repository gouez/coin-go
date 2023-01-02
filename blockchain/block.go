package blockchain

import (
	"bytes"
	"encoding/gob"
	"log"
	"time"

	"github.com/gouez/coin-go/lib/chainhash"
)

// BlockHeader defines information about a block
type Block struct {
	BlockHeader
	Hash chainhash.Hash
	Data []byte
}

// NewBlockHeader return a new BlockHeader
func NewBlock(prevHash chainhash.Hash, data []byte) *Block {
	block := &Block{
		BlockHeader: BlockHeader{
			PrevBlock: prevHash,
			Timestamp: time.Unix(time.Now().Unix(), 0),
		},
		Data: data,
	}
	pow := NewProofOfWork(block)
	hash, nonce := pow.Run()
	block.BlockHeader.Nonce = nonce
	block.Hash = hash
	return block
}

func (b *Block) Serialize() []byte {
	var res bytes.Buffer
	encoder := gob.NewEncoder(&res)

	err := encoder.Encode(b)

	if err != nil {
		log.Panic(err)
	}

	return res.Bytes()
}

func Deserialize(data []byte) *Block {
	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(data))
	err := decoder.Decode(&block)
	if err != nil {
		log.Panic(err)
	}
	return &block
}
