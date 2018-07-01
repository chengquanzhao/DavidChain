package BLC

type BlockChain struct {
	/*
		存储有序的区块
	*/
	Blocks []*Block
}

func CreateGenesisBlockChain() *BlockChain {
	//创世区块
	genesisBlock := CreateGenesisBlock("Genesis Data.......")
	// 包含创世区块的链条
	return &BlockChain{[]*Block{genesisBlock}}

}
