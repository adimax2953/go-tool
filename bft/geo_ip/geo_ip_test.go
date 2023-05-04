package geo_ip

import (
	"net"
	"testing"

	"github.com/adimax2953/go-tool/googledrive"
	LogTool "github.com/adimax2953/log-tool"
	"github.com/oschwald/geoip2-golang"
)

func TestQQ(t *testing.T) {
	serviceAccountFilePath := "ace-destination-385603-579fe600e78a.json"
	g := googledrive.Google{}
	err := g.Init(serviceAccountFilePath)
	fileID := ""
	if err != nil {
		LogTool.LogError("Init error ", err)
		return
	}
	filelist, err := g.GetFileList(10)
	if err != nil {
		LogTool.LogError("GetFileList error ", err)
		return
	}
	if len(filelist) > 0 {
		for _, v := range filelist {
			if v.Name == "GeoLite2-City.mmdb" {
				fileID = v.Id
				LogTool.LogInfo("fileID : ", fileID)
			}
		}
	}
	fileName := "ip.mmdb"
	filePath := "lib/"
	g.DownloadFile(filePath, fileName, fileID)

	db, err := geoip2.Open(filePath + fileName)
	if err != nil {
		LogTool.LogError("open file Error", err)
	}
	defer db.Close()
	// If you are using strings that may be invalid, check that ip is not nil
	ip := net.ParseIP("175.45.20.138")
	record, err := db.City(ip)
	if err != nil {
		LogTool.LogError("City Error", err)
	}

	LogTool.LogInfof("", "Portuguese (BR) city name: %v\n", record.City.Names["pt-BR"])
	if len(record.Subdivisions) > 0 {
		LogTool.LogInfof("", "English subdivision name: %v\n", record.Subdivisions[0].Names["en"])
	}
	LogTool.LogInfof("", "Russian country name: %v\n", record.Country.Names["ru"])
	LogTool.LogInfof("", "ISO country code: %v\n", record.Country.IsoCode)
	LogTool.LogInfof("", "Time zone: %v\n", record.Location.TimeZone)
	LogTool.LogInfof("", "Coordinates: %v, %v\n", record.Location.Latitude, record.Location.Longitude)

	// Output:
	// Portuguese (BR) city name: Londres
	// English subdivision name: England
	// Russian country name: Великобритания
	// ISO country code: GB
	// Time zone: Europe/London
	// Coordinates: 51.5142, -0.0931
}
