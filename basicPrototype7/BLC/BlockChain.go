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

//增加区块到区块链
func (blc *BlockChain)AddBlockchain(data string,height int64,prehash []byte){
	newBlock:= CreateBlock(data, height, prehash)
	blc.Blocks = append(blc.Blocks,newBlock)
}

