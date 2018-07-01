package BLC

import (
	"github.com/boltdb/bolt"
	"log"
	"fmt"
	"math/big"
	"time"
)

const dbName = "blockchain.db"
const blockTableName = "blocks"

type BlockChain struct {
	/*
		存储有序的区块
	*/
	//最新区块的Hash
	Tip [] byte
	DB  *bolt.DB
}

var blockHash []byte
var tableBlock *bolt.Bucket

func CreateGenesisBlockChain() *BlockChain {

	db, err := bolt.Open(dbName, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Update(func(tx *bolt.Tx) error {

		if (tx.Bucket([]byte(blockTableName)) == nil) {
			//创建数据库表
			tableBlock, err = tx.CreateBucket([]byte(blockTableName))
		}

		if err != nil {
			log.Panic(err)
		}

		if tableBlock != nil {
			//创世区块
			genesisBlock := CreateGenesisBlock("Congration Genesis Block Cretared ")
			err := tableBlock.Put(genesisBlock.Hash, genesisBlock.Serialize())
			tableBlock.Put([]byte("l"), genesisBlock.Hash)
			if err != nil {
				log.Panic("Block Created,but don't save...Tray agean.")

			}
			blockHash = genesisBlock.Hash
		}

		return nil
	})

	return &BlockChain{blockHash, db}

}

//增加区块到区块链
func (blc *BlockChain) AddBlockchain(data string) {

	err := blc.DB.Update(func(tx *bolt.Tx) error {
		//获取表

		tb := tx.Bucket([]byte(blockTableName))
		if tb != nil {
			block := DeserializeBlock(tb.Get(blc.Tip))
			// 创建新区块
			newBlock := CreateBlock(data, block.Height+1, block.Hash)


			// 更新blockchain的tip  更新l对应的hash值
			err := tb.Put(newBlock.Hash, newBlock.Serialize())
			if err != nil {
				log.Panic(err)
			}
			blc.Tip = newBlock.Hash


		}
		return nil
	})


	if err != nil {
		log.Panic(err)
	}
}
func (chain *BlockChain) PrintChain() {
	var block *Block

	var currentHash []byte = chain.Tip
	for true {
		err := chain.DB.View(func(tx *bolt.Tx) error {
			bucket := tx.Bucket([]byte(blockTableName))
			if bucket != nil {
				blockBytes := bucket.Get(currentHash)
				block = DeserializeBlock(blockBytes)

				fmt.Printf("Height:%d\n", block.Height)
				fmt.Printf("ProblockHash:%x\n", block.ProBlockHash)
				fmt.Printf("Data:%s\n", block.Data)
				fmt.Printf("Nonce:%d\n", block.Nonce)
				fmt.Printf("Time:%d\n",time.Unix(block.Timestamp,0).Format("2006-01-02"))

			}

			return nil
		})

		if err != nil {
			log.Panic(err)
		}

		var hashInt big.Int
		hashInt.SetBytes(block.ProBlockHash)
		if big.NewInt(0).Cmp(&hashInt) == 0 {
			break
		}
		currentHash = block.ProBlockHash
	}
}
