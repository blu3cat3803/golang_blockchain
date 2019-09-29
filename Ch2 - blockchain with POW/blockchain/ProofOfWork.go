package blockchain

import (
	"fmt"
	"bytes"
	"crypto/sha256"
	"math/big"
	"math"
)

const targetBits = 12
const maxNonce = math.MaxInt64

type ProofOfWork struct {
	block *Block
	target *big.Int
}

func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBits))

    pow := &ProofOfWork{b, target}

    return pow
}

func (pow *ProofOfWork) prepareData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.block.PrevBlockHash,
			pow.block.Data,
			IntToHex(pow.block.Timestamp),
			IntToHex(int64(targetBits)),
			IntToHex(int64(nonce)),
		},
		[]byte{},
	)
	return data
}

func (pow *ProofOfWork) Proof() (int, []byte) {
	var hashInt big.Int
	var hash [32]byte
	n := 0

	fmt.Printf("Mining the block containing \"%s\"\n", pow.block.Data)
	for n < maxNonce {
		data := pow.prepareData(n)

		hash := sha256.Sum256(data)
		fmt.Printf("\r%x", hash)
		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(pow.target) == -1 {
			break
		} else {
			n++
		}
	}
	fmt.Print("\n\n")

	return n, hash[:]
}

func (pow *ProofOfWork) Validate() bool {
	var hashInt big.Int

	data := pow.prepareData(pow.block.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])

	isValid := hashInt.Cmp(pow.target) == -1

	return isValid
}
