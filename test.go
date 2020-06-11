package main

import (
	"blockTest/blockchain"
)

func main() {
	//b := blockchain.NewBlock("", "Gensis Block.")
	//fmt.Println(b)

	//区块链测试
	bc:=blockchain.NewBlockchain()
	bc.AddGensisBlock()
	bc.
		AddBlock("first block").
		AddBlock("second block")
	bc.Iterate()
}
