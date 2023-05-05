package gotool_test

import (
	"encoding/json"
	"testing"
	"time"

	gjson "github.com/adimax2953/go-tool/json"
	LogTool "github.com/adimax2953/log-tool"
)

func Test_Json(t *testing.T) {

	type PingResult struct {
		// in: body
		// required: true
		// example: Success
		Status string `json:"status,omitempty"` //狀態
		// in: body
		// required: true
		// example: 1663557101001
		RequestTime int64 `json:"requestTime,omitempty"` //請求時間
		// in: body
		// min length: 1
		// max length: 16
		// example: pong
		Value string `json:"value,omitempty"` //數值
	}
	args := &PingResult{}
	result := &PingResult{ //<---postResult
		Status:      "Success",
		RequestTime: time.Now().UnixMilli(),
	}

	body, err := gjson.JsonMarshal(result)
	if err != nil {
		result.Status = "result Marshal fail"
	}
	LogTool.LogDebug("result", body)

	if err := gjson.JsonUnmarshal(body, &args); err != nil {
		result.Status = "args Unmarshal fail"
	}
	LogTool.LogDebug("args", *args)
}

func Test_Json2(t *testing.T) {

	type PingResult struct {
		// in: body
		// required: true
		// example: Success
		Status string `json:"status,omitempty"` //狀態
		// in: body
		// required: true
		// example: 1663557101001
		RequestTime int64 `json:"requestTime,omitempty"` //請求時間
		// in: body
		// min length: 1
		// max length: 16
		// example: pong
		Value string `json:"value,omitempty"` //數值
	}
	args := &PingResult{}
	result := &PingResult{ //<---postResult
		Status:      "Success",
		RequestTime: time.Now().UnixMilli(),
	}

	body, err := json.Marshal(result)
	if err != nil {
		result.Status = "result Marshal fail"
	}
	LogTool.LogDebug("result", body)

	if err := json.Unmarshal(body, &args); err != nil {
		result.Status = "args Unmarshal fail"
	}
	LogTool.LogDebug("args", *args)
}
