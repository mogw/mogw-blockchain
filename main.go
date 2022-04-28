package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

// BlockChain is an array of pointers
type BlockChain struct {
	blocks []*Block
}

// Block is a single unit in the blockchain
type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
}

// CreateBlock creates a block
func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash}
	block.DeriveHash()
	return block
}

// AddBlock Will add a Block type unit to a blockchain
func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.blocks[len(chain.blocks)-1]
	newBlock := CreateBlock(data, prevBlock.PrevHash)
	chain.blocks = append(chain.blocks, newBlock)
}

func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	// This will join our previous block's relevant info with the new blocks
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

// Genesis needs to be the first block in a chain, as the first block doesn't have an address to point back to
func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

// InitBlockChain will be what starts a new blockChain
func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}

func main() {

	chain := InitBlockChain()

	chain.AddBlock("first block after genesis")
	chain.AddBlock("second block after genesis")
	chain.AddBlock("third block after genesis")

	for _, block := range chain.blocks {
		fmt.Printf("Previous hash: %x\n", block.PrevHash)
		fmt.Printf("data: %s\n", block.Data)
		fmt.Printf("hash: %x\n", block.Hash)
	}

}
