package main

// Blockchain stores blocks
type Blockchain struct {
	// this is slice syntax... not array
	blocks []*Block
}

// AddBlock adds a new block using NewBlock and appends to blocks slice
func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	// add new block to end of array
	bc.blocks = append(bc.blocks, newBlock)
}

// NewBlockchain creates a new block with a starting genesis block
func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}
