package BLC

import (
	"github.com/boltdb/bolt"
	"log"
)

type BlockChainIterator struct {
	currentHash []byte   // 当前正在遍历的区块的hash值
	DB          *bolt.DB //当前的数据库
}


func (blockChainIterator *BlockChainIterator) Next() *Block {
	var block *Block

	err := blockChainIterator.DB.View(func(tx *bolt.Tx) error {

		bucket := tx.Bucket([]byte(blockTableName))

		if bucket != nil {
			currenHashBytes:= bucket.Get(blockChainIterator.currentHash)

			//获取到当前迭代器当中的 currenthash 所对应的区块
			block = DeserializeBlock(currenHashBytes)

			// 更新迭代器 hash
			blockChainIterator.currentHash = block.ProBlockHash

		}


		return nil
	})

	if err != nil {
		log.Panic(err)
	}

	return block
}