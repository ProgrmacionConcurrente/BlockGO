package main

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

type Block struct {
	Timestamp     int64  //Cuando se creo
	Data          []byte //La data del bloque
	PrevBlockHash []byte //Hash del anterior
	Hash          []byte
}

func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10)) //Casteo de int a byte
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)
	b.Hash = hash[:] //Pasamos tod o el contenido del hash
}

func NewBlock(data string, prevBliackHash []byte) *Block {
	block := &Block{
		Timestamp:     time.Now().Unix(),
		Data:          []byte(data),
		PrevBlockHash: prevBliackHash,
		Hash:          []byte{},
	}
	block.SetHash()
	return block
}
