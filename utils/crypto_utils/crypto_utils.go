package crypto_utils

import (
	"crypto/md5"
	"encoding/hex"
)

// Simple hashing with MD5
// Of course we need stronger security
// but for now it's okay
func GetMD5(input string) string {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(input))
	return hex.EncodeToString(hash.Sum(nil))
}
