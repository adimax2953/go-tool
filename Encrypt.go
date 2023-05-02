package gotool

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/base64"
	"fmt"
)

// MD5
func MD5Encryption(str string) string {
	hash := md5.New()
	hash.Write([]byte(str))
	return fmt.Sprintf("%x", hash.Sum(nil))
}

type Aes struct {
	BlockSize int
	Key       string
	Iv        string
}

// 開始加密
func (a *Aes) Encode(data string) (string, error) {
	_data := []byte(data)
	_key := []byte(a.Key)
	_iv := []byte(a.Iv)

	_data = a.PKCS7Padding(_data)
	block, err := aes.NewCipher(_key)
	if err != nil {
		return "", err
	}
	mode := cipher.NewCBCEncrypter(block, _iv)
	mode.CryptBlocks(_data, _data)
	return base64.StdEncoding.EncodeToString(_data), nil
}

// 開始解密
func (a *Aes) Decode(data string) (string, error) {
	_data, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return "", err
	}
	_key := []byte(a.Key)
	_iv := []byte(a.Iv)

	block, err := aes.NewCipher(_key)
	if err != nil {
		return "", err
	}
	mode := cipher.NewCBCDecrypter(block, _iv)
	mode.CryptBlocks(_data, _data)
	_data = a.PKCS7UnPadding(_data)

	return string(_data), nil
}
func (a *Aes) PKCS7Padding(data []byte) []byte {
	padding := a.BlockSize - len(data)%a.BlockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padtext...)
}
func (a *Aes) PKCS7UnPadding(data []byte) []byte {
	length := len(data)
	unpadding := int(data[length-1])
	return data[:(length - unpadding)]
}
