package geo_ip

import (
	"net"

	LogTool "github.com/adimax2953/log-tool"
	"github.com/oschwald/geoip2-golang"
)

const (
	libPath = "lib/ip.mmdb"
)

// 香港、澳门、台湾、菲律宾、柬埔寨
var blockCountryCode map[string]bool = map[string]bool{
	"":   true,
	"HK": true,
	"MO": true,
	"TW": true,
	"PH": true,
	"KH": true,
}

// todo:
// - check third-party enter game
// - check third-party login hall, room
func CheckBlockIp(sourceIp string) (bool, error) {
	db, err := geoip2.Open(libPath)
	if err != nil {
		LogTool.LogError("", err)
		return false, err
	}
	defer db.Close()

	ip := net.ParseIP(sourceIp)
	record, err := db.City(ip)
	if err != nil {
		LogTool.LogError("", err)
		return false, err
	}

	isoCountryCode := record.Country.IsoCode

	_, found := blockCountryCode[isoCountryCode]
	return found, nil
}

func ParseIp(sourceIp string) (*geoip2.City, error) {
	db, err := geoip2.Open(libPath)
	if err != nil {
		LogTool.LogError("", err)
		return nil, err
	}
	defer db.Close()

	ip := net.ParseIP(sourceIp)
	record, err := db.City(ip)
	if err != nil {
		LogTool.LogError("", err)
		return nil, err
	}

	return record, nil
}
