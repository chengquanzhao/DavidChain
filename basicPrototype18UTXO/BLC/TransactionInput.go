package BLC

type TXinput struct {
	// 交易hash
	TxHash []byte
	// 存储Txoutput在Vout 里面的索引
	Vout int
	// 用户名 --- 数字签名
	ScriptSig string
}
