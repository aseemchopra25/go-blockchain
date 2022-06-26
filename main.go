package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

type Block struct {
	PrevHash []byte
	Data     []byte
	Hash     []byte
}

type Blockchain struct {
	blocks []*Block
}

func (b *Block) SetHash() {
	hash := sha256.Sum256(bytes.Join([][]byte{b.PrevHash, b.Data}, []byte{}))
	b.Hash = hash[:]
}

func (chain *Blockchain) AddBlock(data string) {
	prevHash := chain.blocks[len(chain.blocks)-1].Hash
	newBlock := &Block{[]byte(prevHash), []byte(data), []byte{}}
	newBlock.SetHash()
	chain.blocks = append(chain.blocks, newBlock)
}

func (chain *Blockchain) Initialize() {
	genesisBlock := &Block{[]byte{}, []byte("Learn"), []byte{}}
	genesisBlock.SetHash()
	chain.blocks = append(chain.blocks, genesisBlock)
}

func main() {
	fmt.Println("Creating a simple blockchain using Golang!")
	chain := &Blockchain{}
	chain.Initialize()

	chain.AddBlock("BlockChain")
	chain.AddBlock("in ")
	chain.AddBlock("Golang")

	for _, v := range chain.blocks {
		fmt.Printf("Data: %s\n", v.Data)
		fmt.Printf("Previous Hash: %x\n", v.PrevHash)
		fmt.Printf("Hash:%x\n", v.Hash)
		fmt.Printf("\n")
	}

}
