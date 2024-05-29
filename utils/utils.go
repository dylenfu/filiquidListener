package utils

import (
	"encoding/json"

	"github.com/ethereum/go-ethereum/common"
	"github.com/filiquid/listener/config"
	"golang.org/x/crypto/sha3"
)

func ToJson(data any) ([]byte, error) {
	return json.Marshal(data)
}

func ToString(data any) string {
	raw, err := json.Marshal(data)
	if err != nil {
		return ""
	}
	return string(raw)
}

func HeightToUnix(height uint64) uint64 {
	return (height * 30) + config.CALIBRATION_GENESIS_UNIX_EPOCH
}

func Min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func Keccak256Hash(s1, s2 string) common.Hash {
	hasher := sha3.NewLegacyKeccak256()
	hasher.Write([]byte(s1))
	hasher.Write([]byte(s2))
	data := hasher.Sum(nil)
	return common.BytesToHash(data)
}
