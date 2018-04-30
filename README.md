# GO language take aways

## receiver functions
receiver functions declared like this:
```go
func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)
	b.Hash = hash[:]
}
```

called like this:
```go
block.SetHash()
```

## Slices
easiest way to create slice from array is:
```go
b.Hash = hash[:]
```

## Arrays/Slices declaration
directly declared like this:
```go
[]byte{value, value, value}
```

# Pointers
used to change function args like this:
```go
func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	// add new block to end of array
	bc.blocks = append(bc.blocks, newBlock)
}
```
