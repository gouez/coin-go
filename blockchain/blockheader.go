package blockchain

import (
	"bytes"
	"encoding/binary"
	"io"
	"time"

	"github.com/gouez/coin-go/lib/chainhash"
)

const MaxBlockHeaderPayload = 16 + (chainhash.HashSize * 2)

type BlockHeader struct {
	Version    uint32
	PrevBlock  chainhash.Hash
	MerkelRoot chainhash.Hash
	Timestamp  time.Time
	Bits       uint32
	Nonce      uint32
}

// BlockHash computes the block identifier hash
func (h *BlockHeader) BlockHash() chainhash.Hash {
	buf := bytes.NewBuffer(make([]byte, MaxBlockHeaderPayload))
	_ = writeBlockHeader(buf, h.Version, h.PrevBlock, h.MerkelRoot, uint32(h.Timestamp.Unix()), h.Bits, h.Nonce)
	return chainhash.DoubleHashH(buf.Bytes())
}

// writeBlockHeader write block header to w
func writeBlockHeader(w io.Writer, e ...any) error {
	for _, element := range e {
		err := binary.Write(w, binary.LittleEndian, element)
		if err != nil {
			return err
		}
	}
	return nil
}
