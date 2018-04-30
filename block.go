package main

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

// Block stores data for each block in Blockchain
type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
}

// (args) before funcName() is a 'receiver'. is called as: Block.SetHash()
// pointer is passed so that Hash can be set on block

// SetHash sets the hash on a block using other data to create the hash
func (b *Block) SetHash() {
	// converts timestamp to base10 string and then converts to byte array
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	// arrays are declared with {} in go... they are printed in [] format
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	// create sha256 hash of headers which includes previous PrevBlockHash
	hash := sha256.Sum256(headers)
	// finally set the hash as a slice of array
	b.Hash = hash[:]
}

// NewBlock creates a new block based on previous block's hash
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}}
	block.SetHash()
	return block
}

// NewGenesisBlock creates a starting block for the blockchain
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}
