package main

import "time"

type Blockchain struct {
	blocks []*Block
}

func (bc *Blockchain) AddBlock(data Data) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}

func NewGenesisBlock() *Block {
	data := Data{100, "profesor", "profesor", 1000, time.Now(), time.Now()}
	return NewBlock(data, []byte{})
}

func NewBlockChain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}
