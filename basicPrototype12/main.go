package main

import (
	"DavidChain/basicPrototype12/BLC"
)

func main() {


	// 创世区块
	blockchain := BLC.CreateGenesisBlockChain()
	defer blockchain.DB.Close()
	blockchain.AddBlockchain("Send 1000 RMB To weixiaowei ")
	blockchain.AddBlockchain("Send 2000 RMB To weixiaowei ")
	blockchain.AddBlockchain("Send 3000 RMB To weixiaowei ")

	blockchain.PrintChain()

}
