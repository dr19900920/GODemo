package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"math/big"
)

const targetBits = 24 // 挖矿的难度值 这里的 24 指的是算出来的哈希前 24 位必须是 0 用 16 进制表示化的话 // 0000010000000000000000000000000000000000000000000000000000000000
const maxNonce = math.MaxInt64

type ProofOfWork struct {
	Blo    *Block   // 指向一个块
	Target *big.Int // 一个目标的指针
}

// 在 NewProofOfWork 函数中，我们将 big.Int 初始化为 1，然后左移 256 - targetBits 位。256 是一个 SHA-256 哈希的位数，我们将要使用的是 SHA-256 哈希算法。
// target（目标） 的 16 进制形式为：0x10000000000000000000000000000000000000000000000000000000000
// 它在内存上占据了 29 个字节
func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBits))
	pow := &ProofOfWork{b, target}
	return pow
}

// 下面是与前面例子哈希的形式化比较：
// 0fac49161af82ed938add1d8725835cc123a1a87b1b196488360e58d4bfb51e3
// 0000010000000000000000000000000000000000000000000000000000000000
// 0000008b0f41ec78bab747864db66bcb9fb89920ee75f43fdaaeb5544f7f76ca
// 第一个哈希（基于 “I like donuts” 计算）比目标要大，因此它并不是一个有效的工作量证明。第二个哈希（基于 “I like donutsca07ca” 计算）比目标要小，所以是一个有效的证明。

func (pow *ProofOfWork) prepareData(nonce int) []byte {
	data := bytes.Join([][]byte{
		pow.Blo.PreBlockHash,
		pow.Blo.HashTransactions(),
		IntToHex(pow.Blo.Timestamp),
		IntToHex(int64(targetBits)),
		IntToHex(int64(nonce)),
	}, []byte{})

	return data
}

//准备数据
//用 SHA-256 对数据进行哈希
//将哈希转换成一个大整数
//将这个大整数与目标进行比较
func (pow *ProofOfWork) Run() (int, []byte) {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0
	fmt.Printf("Mining the block containing\n")
	for nonce < maxNonce {
		data := pow.prepareData(nonce)

		hash = sha256.Sum256(data)
		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(pow.Target) == -1 {
			fmt.Printf("\r%x", hash)
			break
		} else {
			nonce++
		}
	}
	fmt.Print("\n\n")
	return nonce, hash[:]
}

// 将一个 int64 转化为一个字节数组(byte array)
func IntToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}
	return buff.Bytes()
}

// Validate block's PoW
func (pow *ProofOfWork) Validate() bool {
	var hashInt big.Int
	data := pow.prepareData(pow.Blo.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])
	isValid := hashInt.Cmp(pow.Target) == -1

	return isValid
}
