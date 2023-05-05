package googledrivetool_test

import (
	"testing"

	"github.com/adimax2953/go-tool/googledrivetool"
	LogTool "github.com/adimax2953/log-tool"
)

func TestQQ(t *testing.T) {
	serviceAccountFilePath := "ace-destination-385603-579fe600e78a.json"
	g := googledrivetool.Google{}
	err := g.Init(serviceAccountFilePath)
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
				err := g.DownloadFile("", "ip.mmdb", v.Id)
				if err != nil {
					LogTool.LogError("ReadFile error ", err)
					return
				}
			}
		}
	}

}
