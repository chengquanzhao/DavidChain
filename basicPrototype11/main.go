package main

import (
	"DavidChain/basicPrototype11/BLC"
	"fmt"
	"github.com/boltdb/bolt"
)

func main() {


	// 创世区块
	blockchain := BLC.CreateGenesisBlockChain()
	defer blockchain.DB.Close()
	blockchain.AddBlockchain("Send 1000 RMB To weixiaowei ")
	blockchain.AddBlockchain("Send 2000 RMB To weixiaowei ")
	blockchain.AddBlockchain("Send 3000 RMB To weixiaowei ")
	fmt.Println(blockchain.DB.View(func(tx *bolt.Tx) error {
		tx.ForEach(func(name []byte, b *bolt.Bucket) error {
			fmt.Println(string(name))
			return nil
		})
		return nil
	}))
	blockchain.PrintChain()

}
