package main

import (
	"fmt"
	"BasicChain/blockchain"
)


func main() {
	chain := blockchain.InitBlockChain()
	chain.AddBlock("first block")
	chain.AddBlock("second block")
	chain.AddBlock("third block")
	for _, block := range chain.Blocks {
		//fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		//fmt.Printf("Current Hash: %x\n", block.Hash)
		fmt.Printf("Data: %x\n", block.Data)
	}

}