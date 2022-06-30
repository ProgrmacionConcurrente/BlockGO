package main

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

type Data struct {
	Id           string    `json:"id"`
	Username     string    `json:"username"`
	Password     string    `json:"password"`
	Amount       float32   `json:"amount"`
	DataCreated  time.Time `json:"dataCreated"`
	LastModified time.Time `json:"lastModified"`
}

type Block struct {
	Timestamp     int64  //Cuando se creo
	Data          Data   //La data del bloque
	PrevBlockHash []byte //Hash del anterior
	Hash          []byte
}

func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10)) //Casteo de int a byte
	headers := bytes.Join([][]byte{b.PrevBlockHash, timestamp}, []byte{})
	hash := sha256.Sum256(headers)
	b.Hash = hash[:] //Pasamos tod o el contenido del hash
}

func NewBlock(data Data, prevBliackHash []byte) *Block {
	block := &Block{
		Timestamp:     time.Now().Unix(),
		Data:          data,
		PrevBlockHash: prevBliackHash,
		Hash:          []byte{},
	}
	block.SetHash()
	return block
}
