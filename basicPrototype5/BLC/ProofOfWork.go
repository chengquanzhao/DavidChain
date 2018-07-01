package BLC

type ProofOfWork struct {
	Block *Block
}
// 创建新的工作量证明 对象

func NewProofOfWork(block *Block) *ProofOfWork {

	return &ProofOfWork{block}
}

func (ProofOfWork *ProofOfWork) Run()([]byte,int64) {

	return nil, 0
}