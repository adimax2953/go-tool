package iotool

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"reflect"
	"strconv"
	"strings"
	"unicode/utf8"

	LogTool "github.com/adimax2953/log-tool"
)

// OpenCS 開csv檔
func OpenCSV(path string) ([][]string, error) {
	// 開啟 CSV 檔案
	if !strings.Contains(path, ".csv") {
		path += ".csv"
	}

	file, err := os.Open(path)
	if err != nil {
		LogTool.LogError("OpenCSV", "無法開啟檔案:%v", err)
		return [][]string{}, err
	}
	defer file.Close()

	//處理UTF-8 編碼的 CSV 文件中，因為UTF-8 編碼的用來標識的字節
	err = handleBOM(file)
	if err != nil {
		return [][]string{}, err
	}

	// 建立 CSV Reader
	reader := csv.NewReader(file)

	// 讀取 CSV 檔案中的所有記錄
	records, err := reader.ReadAll()
	if err != nil {
		LogTool.LogError("OpenCSV", "無法開啟檔案:%v", err)
		return [][]string{}, err
	}

	return records, nil
}

// handleBOM 處理掉在UTF-8編碼的檔案的開首加入一段位元組串EF BB BF
func handleBOM(file *os.File) error {
	// 讀取文件內容
	content := make([]byte, 3)
	_, err := file.Read(content)
	if err != nil {
		return fmt.Errorf("error reading file: %s", err)
	}

	// 檢查是否存在 BOM(在UTF-8編碼的檔案的開首加入一段位元組串EF BB BF)
	if utf8.Valid(content) && content[0] == 0xEF && content[1] == 0xBB && content[2] == 0xBF {
		// 存在，跳過三個字節
		_, err = file.Seek(3, io.SeekStart)
		if err != nil {
			return fmt.Errorf("error seeking file: %s", err)
		}
	} else {
		// 不存在，將文件重置到開頭
		_, err = file.Seek(0, io.SeekStart)
		if err != nil {
			return fmt.Errorf("error seeking file: %s", err)
		}
	}

	return nil
}

// SerializeStructData 檔案序列化為結構(結構欄位要與檔案欄位名稱一致)
func SerializeStructData(field, row []string, docs interface{}) error {
	var errMsg string
	docsV := reflect.ValueOf(docs).Elem()

	for i := 0; i < len(field); i++ {
		fieldName := field[i]

		fieldV := docsV.FieldByName(fieldName)

		if !fieldV.CanSet() {
			continue
		}

		fieldT := fieldV.Type().Kind()

		if len(row) <= 1 {
			continue
		}

		switch fieldT {
		case reflect.String:
			fieldV.SetString(row[i])
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			if v, err := strconv.Atoi(row[i]); err == nil {
				fieldV.SetInt(int64(v))
			}
		case reflect.Float32:
			if f, err := strconv.ParseFloat(row[i], 32); err == nil {
				fieldV.SetFloat(f)
			}
		case reflect.Float64:
			if f, err := strconv.ParseFloat(row[i], 64); err == nil {
				fieldV.SetFloat(f)
			}
		case reflect.Bool:
			if b, err := strconv.ParseBool(row[i]); err == nil {
				fieldV.SetBool(b)
			}
		default:
			errStr := fmt.Sprintf("rwo_%d_%s, unknown type.\n", i, row[i])
			errMsg += errStr
			continue
		}
	}

	if errMsg != "" {
		return errors.New(errMsg)
	}

	return nil
}
