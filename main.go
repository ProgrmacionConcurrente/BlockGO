package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const myurl = "https://serverblockapp.herokuapp.com/api/users"

func main() {
	data := DecodeJson()
	bc := NewBlockChain()
	for i := 0; i < len(data); i++ {
		bc.AddBlock(data[i])
	}

	for _, Block := range bc.blocks {
		fmt.Printf("Hash B.A.: %x\n", Block.PrevBlockHash)
		fmt.Printf("Data: %s\n", Block.Data)
		fmt.Printf("Hash: %x\n\n", Block.Hash)
	}

}

func PerformGetRequest() []byte {
	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)
	return content
	//fmt.Println(content) //content es []unit8
}

func DecodeJson() []Data {
	jsonDataFromWeb := PerformGetRequest()
	var data []Data
	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("Json is accepted")
		json.Unmarshal(jsonDataFromWeb, &data)
	} else {
		fmt.Println("Json no accepted")
	}
	return data
}
