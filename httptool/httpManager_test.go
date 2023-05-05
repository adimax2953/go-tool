package httptool_test

import (
	"testing"

	"github.com/adimax2953/go-tool/httptool"
	LogTool "github.com/adimax2953/log-tool"
)

func Test_http(t *testing.T) {
	gconfig := &httptool.Option{
		URL:                       "http://103.103.81.12:9090/",
		Name:                      "LocalTest",
		MaxConnsPerHost:           100,
		MaxIdemponentCallAttempts: 100,
		ReadTimeout:               15,
		WriteTimeout:              15,
	}

	api, err := httptool.NewClient(gconfig)
	if err != nil {
		LogTool.LogError("", err)
	}
	s, err := api.GetString("ping", "")
	if err != nil {
		LogTool.LogError("", err)
	}
	LogTool.LogInfo(s)

}
