package blockchain

import (
	"math/big"

	"github.com/gouez/coin-go/lib/chainhash"
)

type ProofOfWork struct {
	block  *Block
	target big.Int
}

func NewProofOfWork(block *Block) *ProofOfWork {
	pow := &ProofOfWork{
		block: block,
	}
	targetStr := "00f000000000000000000000000000000000000000000000000000000000000"
	tmpInt := big.Int{}
	tmpInt.SetString(targetStr, 16)
	pow.target = tmpInt
	return pow
}

func (pow *ProofOfWork) Run() (chainhash.Hash, uint32) {
	var nonce uint32
	for {
		pow.block.BlockHeader.Nonce = nonce
		hash := pow.block.BlockHeader.BlockHash()
		tmpInt := big.Int{}
		tmpInt.SetBytes(hash[:])
		if tmpInt.Cmp(&pow.target) == -1 {
			return hash, nonce
		}
		nonce++
	}
}
