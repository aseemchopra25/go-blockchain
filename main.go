package main

import (
	"fmt"
	"strconv"

	"github.com/aseemchopra25/go-blockchain/blockchain"
)

func main() {
	fmt.Println("Creating a simple blockchain using Golang!")
	fmt.Println("-------------------------------------------------------------------------")
	fmt.Println()
	chain := blockchain.InitializeBlockchain()
	chain.AddBlock("BlockChain")
	chain.AddBlock("in ")
	chain.AddBlock("Golang")

	for _, v := range chain.Blocks {
		fmt.Printf("Data: %s\n", v.Data)
		fmt.Printf("Previous Hash: %x\n", v.PrevHash)
		fmt.Printf("Hash:%x\n", v.Hash)
		fmt.Printf("Nonce:%d\n", v.Nonce)
		pow := blockchain.NewProofOfWork(v)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Printf("\n")
	}

}
