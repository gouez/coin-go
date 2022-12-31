package blockchain

import (
	"testing"

	"github.com/gouez/coin-go/lib/logx"
)

// go test -timeout 30000s -run TestNewBlcokChain chan_test.go chan.go block.go pow.go blockheader.go
func TestNewBlcokChain(t *testing.T) {
	chain := NewBlcokChain()
	chain.AddBlock([]byte("1 block"))
	for i, block := range chain.blocks {
		logx.Info("hight:%v", i)
		logx.Info("hash:%v", block.Hash)
		logx.Info("noce:%v", block.BlockHeader.Nonce)
		logx.Info("prehash:%v", block.PrevBlock)
	}
}
