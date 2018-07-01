package BLC

import (
	"math/big"
)

const targetBit = 16

type ProofOfWork struct {
	//区块当前需要验证的
	Block *Block
	// 代表大数据存储
	target *big.Int
}

// 创建新的工作量证明 对象 一个区块 对应一个 工作量证明
func NewProofOfWork(block *Block) *ProofOfWork {
	/*
		big.Int 对象
		0000 0001 -> 向左边移动6位
		0100 0000 -> 64
		0010 0000 -32 < 64
		当难度为16时候
		只要小于 0000 0000 0000 0000 0000 0001 0011 ....{共256位置}....1010

			1.
			左移256 - targetbit
	*/
	//实现为  创建一个初始值 为1 的target
	target := big.NewInt(1)
	target = target.Lsh(target, 256-targetBit)

	return &ProofOfWork{block, target}
}

func (ProofOfWork *ProofOfWork) Run() ([]byte, int64) {

	return nil, 0
}
