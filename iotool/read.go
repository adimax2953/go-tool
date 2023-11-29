package iotool

import (
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	LogTool "github.com/adimax2953/log-tool"
	"gitlab.com/bf_tech/template/slot/gametemplate/pkg/slotgames/tools"
	"gitlab.com/bf_tech/template/slot/gametemplate/pkg/slotgames/tools/random"
)

const (
	dataKeyIndex = 0
	gameDir      = "gametemplate"
	dataDir      = "datas"
)

type FilePathType int

const (
	Config FilePathType = iota
	Mapping
	Table
	PayLines
	Probability
)

var (
	mappingPath     string
	payTablePath    string
	payLinesPath    string
	probabilityPath string
)

func init() {
	currentDir, _ := os.Getwd() //當下目錄
	currentDir = strings.ReplaceAll(currentDir, "\\", "/")
	LogTool.LogInfo("slot init", "current dir:", currentDir)
	//計算返回目標資料夾需要幾階".."
	dirSlice := strings.Split(currentDir, "/")
	numDotDot := 0
	for l := len(dirSlice) - 1; l > 0; l-- {
		numDotDot++
		if dirSlice[l] == gameDir {
			break
		}
	}
	elem := generateDotDotSlice(numDotDot)
	elem = append([]string{currentDir}, elem...)

	// 使用..来回到上一级目录，然后继续返回上级目录，直到回到gameDir
	parentDir := filepath.Join(elem...)

	// 產生進入datas的路徑
	targetDir := filepath.Join(parentDir, dataDir)

	//設定路徑
	mappingPath = strings.ReplaceAll(targetDir+"\\*\\config\\mapping.csv", "\\", "/")
	LogTool.LogInfo("slot init", "mappingPath:", mappingPath)

	payTablePath = strings.ReplaceAll(targetDir+"\\*\\pay\\table.csv", "\\", "/")
	LogTool.LogInfo("slot init", "payTablePath:", payTablePath)

	payLinesPath = strings.ReplaceAll(targetDir+"\\*\\pay\\lines.csv", "\\", "/")
	LogTool.LogInfo("slot init", "payLinesPath:", payLinesPath)

	probabilityPath = strings.ReplaceAll(targetDir+"\\*\\probability\\", "\\", "/")
	LogTool.LogInfo("slot init", "probabilityPath:", probabilityPath)
}

func generateDotDotSlice(n int) []string {
	if n <= 0 {
		return nil
	}
	s := make([]string, n)
	for i := range s {
		s[i] = ".."
	}
	return s
}

type Folder string

func (f Folder) getPath(t FilePathType) string {
	var path string
	switch t {
	case Mapping:
		path = mappingPath
	case Table:
		path = payTablePath
	case PayLines:
		path = payLinesPath
	case Probability:
		path = probabilityPath
	default:
		return ""
	}

	return strings.Replace(path, "*", string(f), 1)
}

type DataMap map[string][]string

// CreateDataMap 建一個讀檔格式(橫式的csv檔,機率表...)
func CreateDataMap(path string) (DataMap, error) {
	records, err := tools.OpenCSV(path)
	if err != nil {
		return nil, err
	}
	dataMap := make(DataMap)
	//逐行處理記錄
	for _, record := range records {
		// 逐欄位讀取記錄中的值,整理資料
		var k string
		if len(record) > dataKeyIndex {
			k = record[dataKeyIndex]
		}
		for i := dataKeyIndex + 1; i < len(record); i++ {
			if record[i] == "" {
				dataMap[k] = record[dataKeyIndex+1 : i]
				break
			} else if i == len(record)-1 {
				dataMap[k] = record[dataKeyIndex+1:]
				break
			}
		}
	}
	return dataMap, nil
}

// GetDataByKey 取該row
func (rm DataMap) GetDataByKey(inKey string) ([]string, bool) {
	data, ok := rm[inKey]
	return data, ok
}

// GetDataByIndex 取該row index裡的資料
func (rm DataMap) GetDataByIndex(inKey string, index int) (string, bool) {
	data, ok := rm[inKey]

	if ok && len(data)-1 >= index {
		return data[index], ok
	}
	return "", ok
}

// ParseRowDataToWeight 轉為權重結構
func (rm DataMap) ParseRowDataToWeight(inKey string) (random.ProbWeight, bool) {
	var rslt random.ProbWeight
	data, ok := rm.GetDataByKey(inKey)
	if ok {
		var probSlice []int64
		for i := 0; i < len(data); i++ {
			val, err := strconv.Atoi(data[i])
			if err != nil {
				return rslt, false
			}
			probSlice = append(probSlice, int64(val))
		}
		rslt = random.CreateProb(probSlice...)
	}
	return rslt, ok
}

// ParseRowDataToInt 轉為int型態
func (rm DataMap) ParseRowDataToInt(inKey string) ([]int, bool) {
	data, ok := rm.GetDataByKey(inKey)
	rslt := make([]int, len(data))
	if ok {
		for i := 0; i < len(data); i++ {
			v, _ := strconv.Atoi(data[i])
			rslt[i] = v
		}
	}
	return rslt, ok
}

// ReadPayData 讀取
func (f Folder) ReadPayData(calcStand int) (*PayTable, error) {
	payTable := NewPayTable(calcStand)
	payTableData, err := CreateDataMap(f.getPath(Table))
	if err != nil {
		return payTable, err
	}
	payTable.InitPayInfo(payTableData)

	if calcStand == int(Line) {
		payLinesData, err := CreateDataMap(f.getPath(PayLines))
		if err != nil {
			return payTable, err
		}
		err = payTable.setPayLine(payLinesData)
		if err != nil {
			return payTable, err
		}
	}
	return payTable, err
}

func (f Folder) ReadPayLinesData() (DataMap, error) {
	return CreateDataMap(f.getPath(PayLines))
}

type ProbabilitySetting struct {
	ID   int
	File string
}

// ReadMappingData 讀機率表設定檔
func (f Folder) ReadMappingData() (map[int]string, error) {

	path := f.getPath(Mapping)

	records, err := tools.OpenCSV(path)

	if err != nil {
		return nil, err
	}

	var (
		field      []string
		settingMap map[int]string = make(map[int]string)
	)

	field = records[0]
	for i := 1; i < len(records); i++ {
		setting := ProbabilitySetting{}
		if err := tools.SerializeStructData(field, records[i], &setting); err != nil {
			log.Println(err)
			continue
		} else {
			settingMap[setting.ID] = setting.File
		}
	}

	return settingMap, nil
}

// ReadProbabilityData 讀機率表
func (f Folder) ReadProbabilityData(file string) (DataMap, error) {
	path := f.getPath(Probability) + file
	return CreateDataMap(path)
}
