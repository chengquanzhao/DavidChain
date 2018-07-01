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

func CreateGenesisBlockChain(txs []*Transaction) *BlockChain {
	if DBExists() {
		fmt.Println("GenesisBlock 已经存在")
		os.Exit(1)

	}


	db, err := bolt.Open(dbName, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Update(func(tx *bolt.Tx) error {

		//创建数据库表
		tableBlock, err = tx.CreateBucket([]byte(blockTableName))

		if err != nil {
			log.Panic(err)
		}

		if tableBlock != nil {
			//创世区块
			genesisBlock := CreateGenesisBlock(txs)
			err := tableBlock.Put(genesisBlock.Hash, genesisBlock.Serialize())
			tableBlock.Put([]byte("l"), genesisBlock.Hash)
			if err != nil {
				log.Panic(err)

			}
			blockHash = genesisBlock.Hash
		}

		return nil
	})

	return &BlockChain{blockHash, db}

}

//增加区块到区块链
func (blc *BlockChain) AddBlockchain(txs []*Transaction) {

	err := blc.DB.Update(func(tx *bolt.Tx) error {
		//获取表
		tb := tx.Bucket([]byte(blockTableName))
		// 创建新区块
		if tb != nil {

			blockBytes := tb.Get(blc.Tip)
			block := DeserializeBlock(blockBytes)

			newBlock := CreateBlock(txs, block.Height+1, block.Hash)

			err := tb.Put(newBlock.Hash, newBlock.Serialize())
			if err != nil {
				log.Panic(err)
			}
			err = tb.Put([]byte("l"), newBlock.Hash)

			// 更新blockchain的tip  更新l对应的hash值
			blc.Tip = newBlock.Hash

		}
		return nil
	})

	if err != nil {
		log.Panic(err)
	}
}

// 打印当前区块链信息
func (chain *BlockChain) PrintChain() {
	chainIterator := chain.Iterator()
	for true {
		block := chainIterator.Next()
		fmt.Printf("Height:%d\n", block.Height)
		fmt.Printf("ProblockHash:%x\n", block.ProBlockHash)
		fmt.Printf("currentHash:%x\n", block.Hash)
		fmt.Printf("Data:%v\n", block.TXs)
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
func DBExists() bool {
	if _, err := os.Stat(dbName); os.IsNotExist(err) {
		return false
	}
	return true
}
func BlockChainObject() *BlockChain {

	db, err := bolt.Open(dbName, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	var hash [] byte
	err = db.View(func(tx *bolt.Tx) error {

		b := tx.Bucket([]byte(blockTableName))
		if b!=nil {
			// 读取最新区块的hash
			hash = b.Get([]byte("l"))
		}
		return nil
	})
	
	
	
	return &BlockChain{hash,db}
}
