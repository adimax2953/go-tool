package encrypttool

import (
	"bytes"
	"crypto/aes"
	"encoding/base64"
)

// PKCS7Padding 填充
func pkcs7Padding(data []byte, blockSize int) []byte {
	padding := blockSize - (len(data) % blockSize)
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

// AES-128-ECB 加密
func EcbEncryptCJ(origData string, key []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	blockSize := block.BlockSize()
	plainText := pkcs7Padding([]byte(origData), blockSize)

	// 分塊進行 ECB 加密
	cipherText := make([]byte, len(plainText))
	for bs, be := 0, blockSize; bs < len(plainText); bs, be = bs+blockSize, be+blockSize {
		block.Encrypt(cipherText[bs:be], plainText[bs:be])
	}

	// CryptoJS 預設輸出的 ciphertext.toString() 為 Base64 編碼
	return base64.StdEncoding.EncodeToString(cipherText), nil
}
