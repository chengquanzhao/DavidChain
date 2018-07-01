package main

import "DavidChain/basicPrototype10/BLC"

func main() {


	// 创世区块
	blockchain := BLC.CreateGenesisBlockChain()
	defer blockchain.DB.Close()
	//fmt.Println(blockchain)
	//
	//fmt.Println(blockchain.Blocks)
	//fmt.Println(blockchain.Blocks[0])
	//// 新区快
	blockchain.AddBlockchain("Send 1000 RMB To weixiaowei ")
	blockchain.AddBlockchain("Send 20 RMB To weixiaowei ")
	blockchain.AddBlockchain("Send 3300 RMB To weixiaowei ")
	//blockchain.AddBlockchain("Send 44500 RMB To weixiaowei ", blockchain.Blocks[len(blockchain.Blocks)-1].Height+1, blockchain.Blocks[len(blockchain.Blocks)-1].Hash)
	//
	//fmt.Println(len(blockchain.Blocks))
	//block := BLC.CreateBlock("test", 1, []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
	//fmt.Printf("%d\n",block.Nonce)
	// fmt.Printf("%x\n",block.Hash)


}
