package algo

import (
	"crypto/sha256"
	"github.com/itchyny/base58-go"
	"fmt"
	"os"
	"math/big"
)

func ToSHA256(input string) []byte {
	algorithm := sha256.New()
	algorithm.Write([]byte(input))

	return algorithm.Sum(nil)
}

func ToBase58(bytes []byte) string {
	encoding := base58.BitcoinEncoding
	encoded, err := encoding.Encode(bytes)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return string(encoded)
}

func GenerateURL(initialLink string, userID string) string {
	hashBytes := ToSHA256(userID + initialLink)
	hashNumber := new(big.Int).SetBytes(hashBytes).Uint64()
	resultString := ToBase58([]byte(fmt.Sprintf("%d",hashNumber)))
	result := resultString[:8]

	return result
}