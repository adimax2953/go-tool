package encrypttool_test

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net/url"
	"strconv"
	"testing"
	"time"

	"github.com/adimax2953/go-tool/encrypttool"
)

func TestEncrypet(t *testing.T) {
	// 模擬環境變數
	aesKey := "" // AES Key 長度需為 16 bytes (AES-128)
	md5Key := ""
	vendor := ""
	encryptedDataParam := ""

	parm := "account=max01&country=CNY&gameId=2087&ip=127.0.0.1"
	key := []byte(aesKey)

	// 1. AES 加密
	encryptedDataCJ, err := encrypttool.EcbEncryptCJ(parm, key)
	if err != nil {
		fmt.Printf("Encryption error: %v\n", err)
		return
	}

	// 2. URL Encode (對齊 encodeURIComponent)
	encryptedData := url.QueryEscape(encryptedDataCJ)
	fmt.Println("Encrypted Data (param):", encryptedData)

	// 3. 取得當前時間戳（秒）
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	fmt.Println("Timestamp:", timestamp)

	// 4. 計算 MD5 簽名 (vendor + timestamp + md5_key)
	md5Input := vendor + timestamp + md5Key
	hash := md5.Sum([]byte(md5Input))
	md5String := hex.EncodeToString(hash[:])
	fmt.Println("Key (MD5):", md5String)

	// 模擬傳入已加密且經過 URL Encode 的字串
	// (填入加密端產生的 encryptedData 即可測試)
	encryptedDataParam = encryptedData

	// 執行解密
	decryptedText, err := encrypttool.EcbDecryptCJ(encryptedDataParam, key)
	if err != nil {
		fmt.Printf("解密失敗: %v\n", err)
		return
	}

	fmt.Println("解密後的原始資料:", decryptedText)
	// 預期輸出: account=max01&country=CNY&gameId=2087&ip=127.0.0.1

}
