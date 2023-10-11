package csv_test

import (
	"testing"

	"github.com/adimax2953/go-tool/bfttool/csv"
	LogTool "github.com/adimax2953/log-tool"
)

func TestQQ(t *testing.T) {
	header := []string{"Name", "Age", "City"}

	data := [][]string{
		{"Alice", "30", "New York"},
		{"Bob", "25", "San Francisco"},
		{"Charlie", "35", "Los Angeles"},
	}

	csvBytes, err := csv.CreateCSVWithHeader(header, data)
	if err != nil {
		panic(err)
	}
	LogTool.LogInfo(string(csvBytes))
}
