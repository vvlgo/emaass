package utils

import (
	"crypto/sha256"
	"fmt"
)

// Encryption sha256加密
func Encryption(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	bs := h.Sum(nil)
	return fmt.Sprintf("%x", bs)
}
