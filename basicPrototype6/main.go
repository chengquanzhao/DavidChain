package main

import (
	"DavidChain/basicPrototype6/BLC"
	"fmt"
)

func main() {

	// 创世区块
	blockchain := BLC.CreateGenesisBlockChain()
	fmt.Println(blockchain)

	fmt.Println(blockchain.Blocks)
	fmt.Println(blockchain.Blocks[0])
	// 新区快
	blockchain.AddBlockchain("Send 1000 RMB To weixiaowei ", blockchain.Blocks[len(blockchain.Blocks)-1].Height+1, blockchain.Blocks[len(blockchain.Blocks)-1].Hash)
	blockchain.AddBlockchain("Send 20 RMB To weixiaowei ", blockchain.Blocks[len(blockchain.Blocks)-1].Height+1, blockchain.Blocks[len(blockchain.Blocks)-1].Hash)
	blockchain.AddBlockchain("Send 3300 RMB To weixiaowei ", blockchain.Blocks[len(blockchain.Blocks)-1].Height+1, blockchain.Blocks[len(blockchain.Blocks)-1].Hash)
	blockchain.AddBlockchain("Send 44500 RMB To weixiaowei ", blockchain.Blocks[len(blockchain.Blocks)-1].Height+1, blockchain.Blocks[len(blockchain.Blocks)-1].Hash)


	fmt.Println(len(blockchain.Blocks))
}
