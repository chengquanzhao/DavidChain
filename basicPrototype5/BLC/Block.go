package BLC

import (
	"time"

)

//定义区块
type Block struct {
	//1.区块高度
	Height int64
	//2.上一个区块的Hash
	ProBlockHash []byte
	// 3.交易数据
	Data []byte
	//4.时间戳
	Timestamp int64
	//5.Hash
	Hash []byte
	//6.添加工作量证明
	Nonce int64
}

/*
	创建新的区块
*/
func CreateBlock(data string, heightBlock int64, preBlockHash []byte) *Block {

	// 创建一个没有Hash的区块
	block := &Block{
		heightBlock,
		preBlockHash,
		[]byte(data),
		time.Now().Unix(),
		nil,
		0}
	//调用工作量证明的方法，返回有效的hash 和nonce值
	pow := NewProofOfWork(block)
	hash,nonce := pow.Run()
	block.Hash = hash[:]
	block.Nonce = nonce
	return block
}

func CreateGenesisBlock(data string) *Block {
	return CreateBlock(data, 1, []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})

}
