package blockchain

import (
	"time"
	"bytes"
	"strconv"
	"crypto/sha256"
)

type Block struct {
	Timestamp int64
	Data []byte
	PrevBlockHash []byte
	Hash []byte
}

func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	payload := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(payload)
	b.Hash = hash[:]
}

func CreateBlock(Data string, PrevBlockHash []byte) *Block{
	block := &Block{
		time.Now().Unix(),
		[]byte(Data),
		PrevBlockHash,
		[]byte{},
	}
	block.SetHash()
	return block
}

type Blockchain struct {
	blocks []*Block
}

func (blockchain *Blockchain) AddBlock(Data string) {
	PrevBlock := blockchain.blocks[len(blockchain.blocks) - 1]
	NewBlock := CreateBlock(Data, PrevBlock.Hash)
	blockchain.blocks = append(blockchain.blocks, NewBlock)
}

func CreateGenesisBlock() *Block {
	return CreateBlock("Genesis Block", []byte{})
}

func CreateBlockchain() *Blockchain {
	return &Blockchain{[] *Block{CreateGenesisBlock()}}
}

