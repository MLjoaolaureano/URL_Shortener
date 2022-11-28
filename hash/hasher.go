package hash

import (
	"crypto/md5"
	"encoding/hex"
)

type HashOperator interface {
	Encode(content string) string
	Compare(content string, hash string) bool
}

func Encode(content string) string {
	hash := md5.Sum([]byte(content))
	return hex.EncodeToString(hash[:])
}

func Compare(content string, hash string) bool {
	contentHash := md5.Sum([]byte(content))
	return hex.EncodeToString(contentHash[:]) == hash
}
