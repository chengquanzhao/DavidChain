package BLC

import (
	"time"
	"fmt"
	"bytes"
	"encoding/gob"
	"log"
	"crypto/sha256"
)

//定义区块
type Block struct {
	//1.区块高度
	Height int64
	//2.上一个区块的Hash
	ProBlockHash []byte
	// 3.交易数据
	TXs []*Transaction
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
func CreateBlock(TXs []*Transaction, heightBlock int64, preBlockHash []byte) *Block {

	// 创建一个没有Hash的区块
	block := &Block{
		heightBlock,
		preBlockHash,
		TXs,
		time.Now().Unix(),
		nil,
		0}
	//调用工作量证明的方法，返回有效的hash 和nonce值
	pow := NewProofOfWork(block)
	// 挖矿验证
	hash, nonce := pow.Run()
	fmt.Println()
	block.Hash = hash[:]
	block.Nonce = nonce
	return block
}

func CreateGenesisBlock(TXs []*Transaction) *Block {
	return CreateBlock(TXs, 1, []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})

}

//将区块序列化成字节数组
func (block *Block) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)
	err := encoder.Encode(block)
	if err != nil {
		log.Panic(err)
	}
	return result.Bytes()
}
//将交易数据转化为字节数组 既TXS
func (block *Block) HashTransactions() []byte {

	var txHashes [][]byte
	var txHash [32]byte

	for _,tx := range block.TXs {

		txHashes = append(txHashes,tx.TxHash)

	}

	txHash = sha256.Sum256(bytes.Join(txHashes,[]byte{}))
	return txHash[:]
}

//反序列化
func DeserializeBlock(blockBytes []byte) *Block {
	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(blockBytes))
	err := decoder.Decode(&block)
	if err!=nil{
		log.Panic(err)
	}

	return &block
}
