package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	//	"strconv"
	"time"
)

type Block struct {
	// 区块头
	Timestamp    int64          // 区块创建的时间
	Transactions []*Transaction // 区块存储的实际有效的信息 交易
	//	Data         []byte // 区块存储的实际有效的信息 交易
	PreBlockHash []byte // 前一个块的哈希
	// 区块体
	Hash  []byte // 当前块的哈希
	Nonce int
}

func NewBlock(transactions []*Transaction, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), transactions, prevBlockHash, []byte{}, 0}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()
	block.Hash = hash
	block.Nonce = nonce
	return block
}

// 计算区块里所有交易的哈希
func (b *Block) HashTransactions() []byte {
	var txHashes [][]byte
	var txHash [32]byte
	for _, tx := range b.Transactions {
		txHashes = append(txHashes, tx.ID)
	}
	txHash = sha256.Sum256(bytes.Join(txHashes, []byte{}))
	return txHash[:]
}

func NewGenesisBlock(coinbase *Transaction) *Block {
	return NewBlock([]*Transaction{coinbase}, []byte{})
}

// Serialize
func (b *Block) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)
	encoder.Encode(b)
	return result.Bytes()
}

// Deserialize
func DeserializeBlock(d []byte) *Block {
	var b Block
	decoder := gob.NewDecoder(bytes.NewReader(d))
	decoder.Decode(&b)
	return &b
}
