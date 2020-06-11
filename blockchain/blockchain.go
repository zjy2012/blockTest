package blockchain

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
	gb := NewBlock("", "the gensis block.")
	bc.lastHash
	return bc
}
func (bc *BlockChian) AddBlock(txs string) *BlockChian {

	return bc
}
