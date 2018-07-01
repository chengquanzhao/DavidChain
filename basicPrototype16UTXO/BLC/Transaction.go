package BLC

import (
	"bytes"
	"encoding/gob"
	"log"
	"crypto/sha256"
)

//UTXO
type Transaction struct {
	//1  交易hash值
	TxHash []byte
	//2  输入
	Vins []*TXinput
	//3  输出
	Vouts []*TxOutput
}

//	1 创世区块创建时候的transaction
func NewCoinBaseTransation(address string) *Transaction {

	//代表消费
	tXinput := &TXinput{[]byte{}, -1, "Genesis Data"}

	txOutput := &TxOutput{10, address}

	txCoinBase := Transaction{[]byte{}, []*TXinput{tXinput}, []*TxOutput{txOutput}}
	//设置hash值
	txCoinBase.HashTransation()

	return &txCoinBase
}

//将transation  序列化 后生成hash 值
func (transation *Transaction) HashTransation() {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)
	err := encoder.Encode(transation)
	if err != nil {
		log.Panic(err)
	}

	hash := sha256.Sum256(result.Bytes())
	transation.TxHash = hash[:]
}
