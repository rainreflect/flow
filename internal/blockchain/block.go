package blockchain

import (
	"bytes"
	"crypto/sha256"
)

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
	*header
}

type header struct {
}

func NewGenesisBlock() *Block {
	return &Block{}
}

func (b *Block) WithHash() {
	blockBytes := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	thisBlockHash := sha256.Sum256(blockBytes)
	b.Hash = thisBlockHash[:]
}

// TODO: из аргументов убрать prevHash
func CreateBlock(data string, prevHash []byte) *Block {
	b := &Block{
		Data:     []byte(data),
		PrevHash: prevHash,
	}
	b.WithHash()
	return b
}
