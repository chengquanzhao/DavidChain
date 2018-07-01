package BLC

import (
	"math/big"
	"bytes"
	"crypto/sha256"
	"fmt"
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
	// 将block的属性拼接成字节数组

	// 生成hash

	// 判断hash 的有效性 阻塞循环
	nonce := 0
	// 存储新生成的hash
	var hashInt big.Int
	var hash [32]byte
	for true {
		dataBytes := ProofOfWork.prepareData(nonce)
		// 生成hash
		hash = sha256.Sum256(dataBytes)

		fmt.Printf("\r%x", hash)
		hashInt.SetBytes(hash[:])
		// 判断hashINT 是否小于target
		if ProofOfWork.target.Cmp(&hashInt) == 1 {
			break
		}

		nonce += 1
	}

	return hash[:], int64(nonce)
}

//数据拼接 返回字节数组

func (pow *ProofOfWork) prepareData(nonce int) []byte {
	data := bytes.Join([][]byte{
		pow.Block.Data,
		pow.Block.ProBlockHash,
		IntToHex(pow.Block.Timestamp),
		IntToHex(int64(targetBit)),
		IntToHex(int64(nonce)),
		IntToHex(int64(pow.Block.Height)),
	}, []byte{})

	return data
}
// 判断生成的hash 是否有效
func (proofOfWork *ProofOfWork) IsValid() bool {
	var hashint big.Int
	hashint.SetBytes(proofOfWork.Block.Hash)

	if proofOfWork.target.Cmp(&hashint)==1{
		return true
	}
	return false

}
