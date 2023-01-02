package blockchain

import (
	"math/big"

	"github.com/gouez/coin-go/lib/chainhash"
)

const Difficulty = 12

// pow
type ProofOfWork struct {
	block  *Block
	target *big.Int
}

func NewProofOfWork(block *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-Difficulty))
	pow := &ProofOfWork{
		block:  block,
		target: target,
	}
	return pow
}

func (pow *ProofOfWork) Run() (chainhash.Hash, uint32) {
	var nonce uint32
	for {
		pow.block.BlockHeader.Nonce = nonce
		hash := pow.block.BlockHeader.BlockHash()
		tmpInt := big.Int{}
		tmpInt.SetBytes(hash[:])
		if tmpInt.Cmp(pow.target) == -1 {
			return hash, nonce
		}
		nonce++
	}
}
