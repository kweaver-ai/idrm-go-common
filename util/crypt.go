package util

import (
	"crypto/md5"
	"encoding/hex"
	"io"
)

func MD5(s string) string {
	hasher := md5.New()
	io.WriteString(hasher, s)
	hashBytes := hasher.Sum(nil)
	return hex.EncodeToString(hashBytes)
}
