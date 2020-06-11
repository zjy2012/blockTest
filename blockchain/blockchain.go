package blockchain

import (
	"fmt"
	. "time"
)

type BlockChian struct {
	lastHash Hash
	blocks   map[Hash]*Block
}

func NewBlockchain() *BlockChian {
	bc := &BlockChian{
		blocks: map[Hash]*Block{},
	}
	return bc
}
func (bc *BlockChian) AddGensisBlock() *BlockChian {
	if bc.lastHash!=""{
return  bc
	}

	return bc.AddBlock("The Gensis Block")
}
func (bc *BlockChian) AddBlock(txs string) *BlockChian {
b:=NewBlock(bc.lastHash,txs)
bc.blocks[b.hashCurr]=b
bc.lastHash=b.hashCurr

	return bc
}
func (bc*BlockChian)Iterate(){
	for hash :=bc.lastHash;hash!="";{
		b:=bc.blocks[hash]
		fmt.Println("HashCurr",b.hashCurr)
		fmt.Println("Tsx",b.txs)
		fmt.Println("Time",b.header.time.Format(UnixDate))
		fmt.Println("HashPrev",b.header.hashPrevBlock)
		fmt.Println("-----------------------------")
		hash=b.header.hashPrevBlock
	}
}