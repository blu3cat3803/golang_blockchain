package blockchain

type Blockchain struct {
	Blocks []*Block
}

func (blockchain *Blockchain) AddBlock(Data []byte) {
	PrevBlock := blockchain.Blocks[len(blockchain.Blocks) - 1]
	NewBlock := CreateBlock(Data, PrevBlock.Hash)
	blockchain.Blocks = append(blockchain.Blocks, NewBlock)
}

func CreateBlockchain() *Blockchain {
	return &Blockchain{[] *Block{CreateGenesisBlock()}}
}

