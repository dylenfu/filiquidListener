package utils

import (
	"encoding/json"
	"io"
	"net/http"

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

// func ToJson(data any) []byte {
// 	jsonData, err := json.Marshal(data)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	return jsonData
// }

// func Min(i, j int) int {
// 	if i < j {
// 		return i
// 	} else {
// 		return j
// 	}
// }

func GetUsers(url string) ([]common.Address, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	result := make([]common.Address, 0)

	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func Keccak256Hash(s1, s2 string) common.Hash {
	hasher := sha3.NewLegacyKeccak256()
	hasher.Write([]byte(s1))
	hasher.Write([]byte(s2))
	data := hasher.Sum(nil)
	return common.BytesToHash(data)
}
