package hash

import (
	"crypto/md5"
	"encoding/hex"
	"log"
)

type Operator interface {
	Encode(content string) string
}

func Encode(content string) string {
	hash := md5.Sum([]byte(content))
	hashString := hex.EncodeToString(hash[:])
	log.Printf("Hash for %s is %s", content, hashString)
	return hashString
}
