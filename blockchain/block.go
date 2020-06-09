package blockchain

import (
	"crypto/sha256"
	"fmt"
	"time"
)

type Hash = string

const nodeVersion = 0

//区块主体
type Block struct {
	header    BlockHeader
	txs       string //交易列表
	txCounter int
	hashCurr  Hash
}

//区块头
type BlockHeader struct {
	version        int
	hashPrevBlock  Hash //前一个区块的Hash
	hashMerkleRoot Hash
	time           time.Time
	bits           int
	nonce          int
}

func NewBlock(preHash Hash, txs string) *Block {
	//实例化Block
	b := &Block{
		header: BlockHeader{
			version:       nodeVersion,
			hashPrevBlock: preHash,
			time:          time.Now(),
		},
		txs:       txs, //设置数据
		txCounter: 1,
	}
	b.SetHashCurr()
	return b
}

//头信息的字符串化
func (bh *BlockHeader) stringify() string {
	return fmt.Sprintf("%d%s%s%d%d%d",
		bh.version,
		bh.hashPrevBlock,
		bh.hashMerkleRoot,
		bh.time.UnixNano(), //得到时间戳，nano级别
		bh.bits,
		bh.nonce,
	)
}

//设置当前区块hash
func (b *Block) SetHashCurr() *Block {
	//生成头信息的拼接字符串
	headerStr := b.header.stringify()
	//计算hash值
	b.hashCurr = fmt.Sprintf("%x", sha256.Sum256([]byte(headerStr)))

	return b
}
