package BLC

import (
	"time"
	"strconv"
	"bytes"
	"crypto/sha256"
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
}

func (block *Block) SetSelfHash() {
	// 1. 块高 转换为自己数组
	heightBytes := IntToHex(block.Height)
	// 2. 时间戳 转化为自己数组
	timeString:= strconv.FormatInt(block.Timestamp, 2)
	timeBytes := []byte(timeString)
	// 3. 将属性数据拼接起来
	blockBytes:= bytes.Join([][]byte{heightBytes, block.ProBlockHash,
		block.Data, timeBytes, block.Hash}, []byte{})
	// 4.生成Hash
	sum256Hah := sha256.Sum256(blockBytes)

	block.Hash = sum256Hah[:]
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
		nil}

	//设置自身hash值
	block.SetSelfHash()

	return block
}
