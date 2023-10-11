package csv

import (
	"bytes"
	"encoding/csv"
	"errors"

	LogTool "github.com/adimax2953/log-tool"
)

func CreateCSVWithHeader(header []string, body [][]string) ([]byte, error) {

	// 创建一个内存中的缓冲区来存储CSV数据
	var buf bytes.Buffer

	// 创建一个 CSV writer，将数据写入到缓冲区中
	writer := csv.NewWriter(&buf)

	if header == nil || len(header) == 0 {
		return nil, errors.New("Header不得為空")
	}
	if body == nil || len(body) == 0 {
		return nil, errors.New("body不得為空")
	}
	// 刷新并关闭 CSV writer
	defer writer.Flush()

	// 写入 CSV 头部
	if err := writer.Write(header); err != nil {
		LogTool.LogError("写入 CSV 头部时出错:", err)
		return nil, err
	}

	// 逐行写入数据
	for _, record := range body {
		if err := writer.Write(record); err != nil {
			LogTool.LogError("写入 CSV 数据时出错:", err)
			return nil, err
		}
	}

	// 检查是否有错误
	if err := writer.Error(); err != nil {
		LogTool.LogError("CSV 写入错误:", err)
		return nil, err
	}

	return buf.Bytes(), nil
}
