package encrypttool

import (
	"crypto/aes"
	"encoding/base64"
	"fmt"
	"net/url"
)

// PKCS7Unpadding 去除填充內容
func pkcs7Unpadding(origData []byte) ([]byte, error) {
	length := len(origData)
	if length == 0 {
		return nil, fmt.Errorf("data is empty")
	}
	// 取得最後一個 byte 獲取填充長度
	unpadding := int(origData[length-1])
	if unpadding > length || unpadding == 0 {
		return nil, fmt.Errorf("invalid padding size")
	}
	return origData[:(length - unpadding)], nil
}

// AES-128-ECB 解密
func EcbDecryptCJ(encryptedStr string, key []byte) (string, error) {
	// 1. 先解開 URL Encode
	decodedURL, err := url.QueryUnescape(encryptedStr)
	if err != nil {
		return "", fmt.Errorf("url unescape error: %v", err)
	}

	// 2. Base64 解碼
	cipherText, err := base64.StdEncoding.DecodeString(decodedURL)
	if err != nil {
		return "", fmt.Errorf("base64 decode error: %v", err)
	}

	// 3. 建立 AES Cipher
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	blockSize := block.BlockSize()
	if len(cipherText)%blockSize != 0 {
		return "", fmt.Errorf("ciphertext length is not a multiple of the block size")
	}

	// 4. 分塊進行 ECB 解密
	plainText := make([]byte, len(cipherText))
	for bs, be := 0, blockSize; bs < len(cipherText); bs, be = bs+blockSize, be+blockSize {
		block.Decrypt(plainText[bs:be], cipherText[bs:be])
	}

	// 5. 去除 PKCS7 Padding
	origBytes, err := pkcs7Unpadding(plainText)
	if err != nil {
		return "", fmt.Errorf("unpadding error: %v", err)
	}

	return string(origBytes), nil
}
