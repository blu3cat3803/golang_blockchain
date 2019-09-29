package blockchain

import (
	"time"
	//"bytes"
	//"strconv"
	//"crypto/sha256"
)
/*
// Original Block
type Block struct {
	Timestamp int64
	Data []byte
	PrevBlockHash []byte
	Hash []byte
}


//Original Block
func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	payload := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(payload)
	b.Hash = hash[:]
}

func CreateBlock(Data []byte, PrevBlockHash []byte) *Block{
	block := &Block{
		time.Now().Unix(),
		Data,
		PrevBlockHash,
		[]byte{},
	}
	block.SetHash()
	return block
}
*/
type Block struct {
	Timestamp int64
	Data []byte
	PrevBlockHash []byte
	Hash []byte
	Nonce int
}

func CreateBlock(Data []byte, prevBlockHash []byte) *Block {
    block := &Block{
    	time.Now().Unix(),
    	[]byte(Data),
    	prevBlockHash,
    	[]byte{},
    	0,
    }
    pow := NewProofOfWork(block)
    nonce, hash := pow.Proof()

    block.Hash = hash[:]
    block.Nonce = nonce

    return block
}

func CreateGenesisBlock() *Block {
	return CreateBlock([]byte("Genesis Block"), []byte{})
}
