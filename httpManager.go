package gotool

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	LogTool "github.com/adimax2953/log-tool"
)

//发送GET请求
//url:请求地址
//response:请求返回的内容
func Get(url string) (response string) {
	client := http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		LogTool.LogDebug(fmt.Sprintf("Get-err-1 %+v", err))
		return
	}
	defer resp.Body.Close() //必須調用否則可能產生記憶體洩漏

	var buffer [512]byte
	result := bytes.NewBuffer(nil)
	for {
		n, err := resp.Body.Read(buffer[0:])
		result.Write(buffer[0:n])
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			LogTool.LogDebug(fmt.Sprintf("Get-err-2 %+v", err))
			return
		}
	}

	response = result.String()
	return
}

//发送POST请求
//url:请求地址, data:POST请求提交的数据,contentType:请求体格式, 如：application/json
//content:请求放回的内容
func Post(url string, data interface{}, contentType string) (content string) {
	jsonStr, _ := json.Marshal(data)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Add("content-type", contentType)
	if err != nil {
		LogTool.LogDebug(fmt.Sprintf("Post-Err-1 %+v", err))
		return
	}
	defer req.Body.Close() //必須調用否則可能產生記憶體洩漏

	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		LogTool.LogDebug(fmt.Sprintf("Post-Err-2 %+v", err))
		return
	}
	defer resp.Body.Close()

	result, _ := ioutil.ReadAll(resp.Body)
	content = string(result)

	return

}
