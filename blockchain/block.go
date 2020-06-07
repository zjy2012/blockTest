package blockchain

import (
	"crypto/sha256"
	"fmt"
	time2 "time"
)

const nodeVersion = 0

type Hash = string

func NewBlock(prevHash Hash, tsx string) *Block {
	b := &Block{
		header: BlockHeader{
			version:       nodeVersion,
			hashPrevBlock: prevHash, // 设置前面的区块哈希
			time:          time2.Now(),
		},
		txs:       tsx, // 设置数据
		txCounter: 1,   // 计数交易
	} // 计算设置当前区块的哈希
	b.setHashCurr()
	return b
}

//区块主体
type Block struct {
	header    BlockHeader //结构体的嵌套
	txs       string      //交易列表
	txCounter int         //交易计数器
	hashCurr  Hash        //当前区块链哈希
}

//去块头
type BlockHeader struct {
	version        int
	hashPrevBlock  Hash
	hashMerkleRoot Hash
	time           time2.Time
	bits           int
	nonce          int
} //这里会涉及到一个结构体的封装性 ，外面的包不能随便的进行调用，所以会设置set 和个体
//方法
func (bh *BlockHeader) stringify() string {
	return fmt.Sprintf("%d%s%s%d%d%d",
		bh.version,
		bh.hashPrevBlock,
		bh.hashMerkleRoot,
		bh.time.UnixNano(), // 得到时间戳，nano 级别
		bh.bits,
		bh.nonce,
	)
}

//设置当前区块hash
func (b *Block) setHashCurr() *Block {
	headerStr := b.header.stringify()
	b.hashCurr = fmt.Sprintf("%x", sha256.Sum256([]byte (headerStr)))
	return b
}
