package utils

import (
	"github.com/Luzifer/go-openssl"
	"go.uber.org/zap"
)

// DecrypteString 解密
func DecrypteString(key string, encrypted string, logger *zap.Logger) (plaintext string) {
	o := openssl.New()
	dec, err := o.DecryptBytes(key, []byte(encrypted))
	if err != nil {
		logger.Error("解密ssh密码失败", zap.String("secrect", encrypted), zap.String("errors", err.Error()))
	}
	plaintext = string(dec)
	return
}
