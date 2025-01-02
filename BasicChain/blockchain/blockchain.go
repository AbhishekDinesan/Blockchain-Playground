package blockchain
// go.sum is like a package.json, depdendency stuffs

import (
	"bytes"
	"crypto/sha256"
)

type Block struct{
	Hash []byte // derived from other two fields
	Data []byte
	PrevHash []byte // last block hash, like a shitty linked list
}

type BlockChain struct{
	Blocks []*Block
}


func (b *Block) DeriveHash(){
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:] // throws hash into block
}

func CreateBlock(Data string, prevHash []byte) *Block{
	block := &Block{[]byte{}, []byte(Data), prevHash}
	block.DeriveHash()
	return block
}

func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

func (chain *BlockChain) AddBlock(data string){
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	new := CreateBlock(data, prevBlock.Hash)
	chain.Blocks = append(chain.Blocks, new)
}

func InitBlockChain() *BlockChain{
	return &BlockChain{[]*Block{Genesis()}}
}