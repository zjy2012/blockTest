package main

import (
	"blockTest/blockchain"
	"fmt"
)

func main() {
	b := blockchain.NewBlock("", "Gensis Block.")
	fmt.Println(b)
}
