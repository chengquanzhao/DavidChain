package BLC

import (
	"github.com/boltdb/bolt"
	"log"
	"fmt"
	"math/big"
	"time"
	"os"
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

func (blockChain *BlockChain) Iterator() *BlockChainIterator {

	return &BlockChainIterator{blockChain.Tip, blockChain.DB}

}

func CreateGenesisBlockChain() *BlockChain {
	if dbExists() {
		var blockchain *BlockChain
		db, err := bolt.Open(dbName, 0600, nil)
		if err != nil {
			log.Fatal(err)
		}

		err = db.View(func(tx *bolt.Tx) error {
			//获取表

			tb := tx.Bucket([]byte(blockTableName))
			if tb != nil {
				hash := tb.Get([]byte("l"))
				blockchain = &BlockChain{hash, db}

			}
			return nil
		})

		if err != nil {
			log.Panic(err)
		}

		return blockchain

	}
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
	fmt.Println("120 is in")
}
// 打印当前区块链信息
func (chain *BlockChain) PrintChain() {
	chainIterator := chain.Iterator()
	for true {
		block := chainIterator.Next()
		fmt.Printf("Height:%d\n", block.Height)
		fmt.Printf("ProblockHash:%x\n", block.ProBlockHash)
		fmt.Printf("currentHash:%x\n", block.Hash)
		fmt.Printf("Data:%s\n", block.Data)
		fmt.Printf("Nonce:%d\n", block.Nonce)
		fmt.Printf("Time:%v\n", time.Unix(block.Timestamp, 0).Format("2006-01-02 03:04:05 pm"))
		fmt.Println()

		var hashInt big.Int
		hashInt.SetBytes(block.ProBlockHash)
		if big.NewInt(0).Cmp(&hashInt) == 0 {
			break;
		}
	}
}

//判断数据库是否存在
func dbExists() bool {
	if _, err := os.Stat(dbName); os.IsNotExist(err) {
		return false
	}
	return true
}
