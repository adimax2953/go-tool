package kafkatool_test

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	gotool "github.com/adimax2953/go-tool"
	"github.com/adimax2953/go-tool/kafkatool"
	"github.com/adimax2953/go-tool/randtool"
	"github.com/shopspring/decimal"
)

var c kafkatool.KafkaConfig

var roundID string = "0000000000"
var trandID string = "0000000000"

func Test_SendtoKafka(t *testing.T) {

	config := &kafkatool.KafkaConfig{
		Address:           "192.168.10.151:9092",
		Network:           "tcp",
		NumPartition:      0,
		ReplicationFactor: 1,
	}
	config.CreateTopic("USS-Game", 10)
	c = *config
	y, w := gotool.GetWeek()
	roundID = fmt.Sprintf("%s%s%06d", gotool.Encode10To62(int64(y))+gotool.Encode10To62(int64(w)), "01", 0)

	mlist, id := bet(roundID)
	c.WriteMessagesKeyValueList("USS-Game", mlist)
	mlist, id = win(id)
	c.WriteMessagesKeyValueList("USS-Game", mlist)
	c.WriteMessagesKeyValueList("USS-Game", refund(id, roundID))
	//LogTool.LogDebug("", roundID)

	//config.WriteMessagesKeyValue("test-USS-Game", m)

	//config.WriteMessagesKeyValue("test03", m)
	//config.ReadMessages("test02", "1")
	//config.GetTopic()
	//config.DelTopic(config.GetTopic()...)
}

func bet(id string) ([]kafkatool.WriteData, string) {

	tid := gotool.Base62Increment(id)
	rid := gotool.Base62Increment(id)
	count := 20
	mlist := make([]kafkatool.WriteData, count)
	for i := 0; i < count; i++ {

		t := time.Now().Unix()
		d, h := gotool.DateTimeFromTimeStamp(t)
		tid = gotool.Base62Increment(tid) + gotool.Encode10To62(int64(randtool.GetRandom(62)))
		rid = gotool.Base62Increment(rid) + gotool.Encode10To62(int64(randtool.GetRandom(62)))

		var bonus int64 = 0
		quantity := decimal.NewFromInt(bonus).Mul(decimal.NewFromFloat(0.04))
		var fee int64 = gotool.Str2int64(quantity.String())
		var value int64 = 10000 + bonus - fee

		gamebet := &[]GameBetResult{{
			BetID:      1,
			Value:      value,
			FinishTime: t,
		}, {
			BetID:      2,
			Value:      value,
			FinishTime: t,
		}}

		gamelog := &GameLog{
			TransactionID:   tid,
			RoundID:         rid,
			GameCode:        "01",
			Value:           value,
			Bonus:           bonus,
			Fee:             fee,
			Country:         "THD",
			TransactionType: "bet",
			PlayerName:      "15656561",
			SiteCode:        "TG",
			Platform:        "UFA",
			GameResult:      gamebet,
			DateStr:         d,
			TimeStr:         h,
			IsFree:          false,
			UIOrientation:   "vertical",
			Timestamp:       time.Now().UnixMilli(),
		}
		jsonBytes, err := json.Marshal(gamelog)
		if err != nil {
		}
		m := &kafkatool.WriteData{
			Key:   gamelog.SiteCode,
			Value: string(jsonBytes),
		}
		ms := map[string]string{}
		ms["tg"] = string(jsonBytes)
		mlist[i] = *m
	}
	return mlist, tid
}
func win(id string) ([]kafkatool.WriteData, string) {
	tid := gotool.Base62Increment(id) + gotool.Encode10To62(int64(randtool.GetRandom(62)))
	rid := gotool.Base62Increment(id) + gotool.Encode10To62(int64(randtool.GetRandom(62)))
	count := 20
	mlist := make([]kafkatool.WriteData, count)
	for i := 0; i < count; i++ {

		t := time.Now().Unix()
		d, h := gotool.DateTimeFromTimeStamp(t)
		tid = gotool.Base62Increment(tid)
		rid = gotool.Base62Increment(rid)

		var bonus int64 = 30000
		quantity := decimal.NewFromInt(bonus).Mul(decimal.NewFromFloat(0.04))
		var fee int64 = gotool.Str2int64(quantity.String())
		var value int64 = 10000 + bonus - fee

		gamebet := &[]GameWinResult{{
			BetID:      1,
			Value:      value,
			FinishTime: t,
			PL:         "20",
		}, {
			BetID:      1,
			Value:      value,
			FinishTime: t,
			PL:         "20",
		}}

		gamelog := &GameLog{
			TransactionID:   tid,
			RoundID:         rid,
			GameCode:        "01",
			Value:           value,
			Bonus:           bonus,
			Fee:             fee,
			Country:         "THD",
			TransactionType: "win",
			PlayerName:      "15656561",
			SiteCode:        "TG",
			Platform:        "UFA",
			GameResult:      gamebet,
			DateStr:         d,
			TimeStr:         h,
			IsFree:          false,
			UIOrientation:   "vertical",
			Timestamp:       time.Now().UnixMilli(),
		}
		jsonBytes, err := json.Marshal(gamelog)
		if err != nil {
		}
		m := &kafkatool.WriteData{
			Key:   gamelog.SiteCode,
			Value: string(jsonBytes),
		}
		ms := map[string]string{}
		ms["tg"] = string(jsonBytes)
		mlist[i] = *m
	}
	return mlist, tid
}
func refund(id, relate string) []kafkatool.WriteData {
	tid := gotool.Base62Increment(id) + gotool.Encode10To62(int64(randtool.GetRandom(62)))
	rid := gotool.Base62Increment(id) + gotool.Encode10To62(int64(randtool.GetRandom(62)))
	rlateid := gotool.Base62Increment(relate) + gotool.Encode10To62(int64(randtool.GetRandom(62)))

	count := 20
	mlist := make([]kafkatool.WriteData, count)
	for i := 0; i < count; i++ {

		t := time.Now().Unix()
		d, h := gotool.DateTimeFromTimeStamp(t)
		tid = gotool.Base62Increment(tid)
		rid = gotool.Base62Increment(rid)
		rlateid = gotool.Base62Increment(rlateid)

		var bonus int64 = 0
		quantity := decimal.NewFromInt(bonus).Mul(decimal.NewFromFloat(0.04))
		var fee int64 = gotool.Str2int64(quantity.String())
		var value int64 = 10000 + bonus - fee

		gamelog := &GameLog{
			TransactionID:   tid,
			RoundID:         rid,
			RelateID:        rlateid,
			GameCode:        "01",
			Value:           value,
			Bonus:           bonus,
			Fee:             fee,
			Country:         "THD",
			TransactionType: "refund",
			PlayerName:      "15656561",
			SiteCode:        "TG",
			Platform:        "UFA",
			DateStr:         d,
			TimeStr:         h,
			IsFree:          false,
			UIOrientation:   "vertical",
			Timestamp:       time.Now().UnixMilli(),
		}
		jsonBytes, err := json.Marshal(gamelog)
		if err != nil {
		}
		m := &kafkatool.WriteData{
			Key:   gamelog.SiteCode,
			Value: string(jsonBytes),
		}
		ms := map[string]string{}
		ms["tg"] = string(jsonBytes)
		mlist[i] = *m
	}
	return mlist
}

type GameLog struct {
	TransactionID   string      `json:"transactionId"`
	RoundID         string      `json:"roundId"`
	RelateID        string      `json:"relateId"`
	GameCode        string      `json:"gameCode"`
	Value           int64       `json:"value"`
	Bonus           int64       `json:"bonus"`
	Fee             int64       `json:"fee"`
	Country         string      `json:"country"`
	TransactionType string      `json:"transactionType"`
	PlayerID        string      `json:"playerId"`
	PlayerName      string      `json:"playerName"`
	SiteCode        string      `json:"siteCode"`
	Platform        string      `json:"platform"`
	GameResult      interface{} `json:"gameResult"`
	DateStr         string      `json:"dateStr"`
	TimeStr         string      `json:"timeStr"`
	IsFree          bool        `json:"isFree"`
	UIOrientation   string      `json:"uiOrientation"`
	Timestamp       int64       `json:"@timestamp"`
}
type GameBetResult struct {
	BetID      int   `json:"betId"`
	Value      int64 `json:"value"`
	FinishTime int64 `json:"finishTime"`
}
type GameWinResult struct {
	BetID         int    `json:"betId"`
	Win           int64  `json:"win"`
	Stake         string `json:"stake"`
	WinMultiplier string `json:"winMultiplier"`
	Insurcnce     string `json:"insurcnce"`
	PL            string `json:"pl"`
	FinishTime    int64  `json:"finishTime"`
	Value         int64  `json:"value"`
}

type GameRankResult struct {
	GameCode     string              `json:"gameCode"`
	Gap          string              `json:"gap"`
	Country      string              `json:"country"`
	GameRankList []GameSubRankResult `json:"gameRankList"`
	FinishTime   int64               `json:"finishTime"`
}
type GameSubRankResult struct {
	No       string `json:"no"`
	Name     string `json:"name"`
	Fraction int64  `json:"fraction"`
}
