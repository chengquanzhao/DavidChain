package BLC

import (
	"github.com/boltdb/bolt"
	"log"
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

		if  (tx.Bucket([]byte(blockTableName)) == nil) {
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
	//blc.Blocks = append(blc.Blocks, newBlock)

	err := blc.DB.Update(func(tx *bolt.Tx) error {
		//获取表

		tb := tx.Bucket([]byte(blockTableName))
		if tb == nil {

		}
		block := DeserializeBlock(tb.Get(blc.Tip))
		// 创建新区块
		newBlock := CreateBlock(data, block.Height, block.ProBlockHash)

		// 将区块序列化

		// 更新blockchain的tip  更新l对应的hash值
		err := tb.Put(newBlock.Hash, newBlock.Serialize())
		if err != nil {
			log.Panic(err)
		}
		blc.Tip = newBlock.Hash

		return nil
	})
	if err != nil {
		log.Panic(err)
	}
}
