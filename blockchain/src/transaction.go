package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
	"fmt"
	"log"
)

// 奖励的数额
const subsidy = 10

// Transaction 由交易 ID，输入和输出构成
type Transaction struct {
	ID   []byte
	Vin  []TXInput
	Vout []TXOutput
}

// 交易输入
type TXInput struct {
	Txid      []byte // 一个输入引用了之前一笔交易的一个输出：Txid 存储的是这笔交易的 ID
	Vout      int    // 存储的是该输出在这笔交易中所有输出的索引（因为一笔交易可能有多个输出，需要有信息指明是具体的哪一个）
	ScriptSig string // 一个脚本，提供了可作用于一个输出的 ScriptPubKey 的数据
}

// 交易输出 在比特币中，value 字段存储的是 satoshi 的数量，而不是>有 BTC 的数量。一个 satoshi 等于一百万分之一的 >BTC(0.00000001 BTC)，这也是比特币里面最小的货币单位>（就像是 1 分的硬币）
type TXOutput struct {
	Value        int    // 币
	ScriptPubKey string // 锁定和解锁输出的逻辑
}

// 创世块
func NewCoinbaseTX(to, data string) *Transaction {
	if data == "" {
		data = fmt.Sprintf("Reword to '%s'", to)
	}
	txin := TXInput{[]byte{}, -1, data}
	txout := TXOutput{subsidy, to}
	tx := Transaction{nil, []TXInput{txin}, []TXOutput{txout}}
	tx.SetID()
	return &tx
}

func (tx *Transaction) SetID() {
	var encoded bytes.Buffer
	var hash [32]byte

	enc := gob.NewEncoder(&encoded)
	err := enc.Encode(tx)
	if err != nil {
		log.Panic(err)
	}
	hash = sha256.Sum256(encoded.Bytes())
	tx.ID = hash[:]
}

// 输入的脚本是否可以解锁输出
func (in *TXInput) CanUnlockOutputWith(unlockingData string) bool {
	return in.ScriptSig == unlockingData
}

// 输出的的脚本是否可以解锁
func (out *TXOutput) CanBeUnlockedWith(unlockingData string) bool {
	return out.ScriptPubKey == unlockingData
}

// IsCoinbase 判断是否是 coinbase 交易
func (tx Transaction) IsCoinbase() bool {
	return len(tx.Vin) == 1 && len(tx.Vin[0].Txid) == 0 && tx.Vin[0].Vout == -1
}

// 发送币
//找到所有的未花费输出，并且确保它们存储了足够的值（value）
//对于每个找到的输出，会创建一个引用该输出的输入。接下来，我们创建两个输出：
//一个由接收者地址锁定。这是给实际给其他地址转移的币。
//一个由发送者地址锁定。这是一个找零。只有当未花费输出超过新交易所需时产生。记住：输出是不可再分的。
func NewUTXOTransaction(from, to string, amount int, bc *BlockChain) *Transaction {
	var inputs []TXInput
	var outputs []TXOutput
	acc, validOutputs := bc.FindSpendableOutputs(from, amount)
	if acc < amount {
		log.Panic("ERROR: Not enough funds")
	}
	// Build a list of inputs
	for txid, outs := range validOutputs {
		txID, _ := hex.DecodeString(txid)
		for _, out := range outs {
			input := TXInput{txID, out, from}
			inputs = append(inputs, input)
		}
	}
	// Build a list of outputs
	outputs = append(outputs, TXOutput{amount, to})
	if acc > amount {
		outputs = append(outputs, TXOutput{acc - amount, from})
	}
	tx := Transaction{nil, inputs, outputs}
	tx.SetID()
	return &tx
}
