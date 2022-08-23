package kafkatool_test

import (
	"encoding/json"
	"testing"
	"time"

	gotool "github.com/adimax2953/go-tool"
	"github.com/adimax2953/go-tool/kafkatool"
	LogTool "github.com/adimax2953/log-tool"
	"github.com/shopspring/decimal"
)

var c kafkatool.KafkaConfig

var roundID string = "00000000000"
var trandID string = "00000000000"

func Test_SendtoKafka(t *testing.T) {

	config := &kafkatool.KafkaConfig{
		Address:           "192.168.10.151:9092",
		Network:           "tcp",
		NumPartition:      0,
		ReplicationFactor: 1,
	}

	config.CreateTopic("USS-Game", 10)
	c = *config
	mlist, id := bet(roundID)
	c.WriteMessagesKeyValueList("USS-Game", mlist)
	mlist, id = bet(roundID)
	c.WriteMessagesKeyValueList("USS-Game", mlist)
	c.WriteMessagesKeyValueList("USS-Game", refund(id, roundID))

	LogTool.LogDebug("", roundID)

	//config.WriteMessagesKeyValue("test-USS-Game", m)

	//config.WriteMessagesKeyValue("test03", m)
	//config.ReadMessages("test02", "1")
	//config.GetTopic()
	//config.DelTopic(config.GetTopic()...)

	//config.CreateTopic("test1112")
	//config.CreateConn("test12")
	//config.WriteMessages("test3", "da", "da", "der", "ma", "te", "sen")
}

func bet(id string) ([]kafkatool.WriteData, string) {

	tid := gotool.Base62Increment(id)
	rid := gotool.Base62Increment(id)
	count := 100000
	mlist := make([]kafkatool.WriteData, count)
	for i := 0; i < count; i++ {

		t := time.Now().Unix()
		d, h := gotool.DateTimeFromTimeStamp(t)
		tid = gotool.Base62Increment(tid)
		rid = gotool.Base62Increment(rid)
		var bonus int64 = 0
		quantity := decimal.NewFromInt(bonus).Mul(decimal.NewFromFloat(0.04))
		var fee int64 = gotool.Str2int64(quantity.String())
		var value int64 = 10000 + bonus - fee

		gamebet := &GameBetResult{
			TransactionID:   tid,
			RoundID:         rid,
			TransactionType: "bet",
			GameCode:        "THUSS",
			BetID:           1,
			Country:         "CDN",
			Value:           value,
			FinishTime:      t,
		}

		gamelog := &GameLog{
			TransactionID:   tid,
			RoundID:         rid,
			GameCode:        "THUSS",
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
			Timestamp:       gotool.TimeStamptoDateTime(t),
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
	tid := gotool.Base62Increment(id)
	rid := gotool.Base62Increment(id)
	count := 100000
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

		gamebet := &GameWinResult{
			TransactionID:   tid,
			RoundID:         rid,
			TransactionType: "win",
			GameCode:        "THUSS",
			BetID:           1,
			Value:           value,
			FinishTime:      t,
		}

		gamelog := &GameLog{
			TransactionID:   tid,
			RoundID:         rid,
			GameCode:        "THUSS",
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
			Timestamp:       gotool.TimeStamptoDateTime(t),
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
	tid := gotool.Base62Increment(id)
	rid := gotool.Base62Increment(id)
	rlateid := gotool.Base62Increment(relate)

	count := 100000
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
			GameCode:        "THUSS",
			Value:           value,
			Bonus:           bonus,
			Fee:             fee,
			Country:         "THD",
			TransactionType: "refund",
			PlayerName:      "15656561",
			SiteCode:        "TG",
			Platform:        "UFA",
			GameResult:      "",
			DateStr:         d,
			TimeStr:         h,
			IsFree:          false,
			UIOrientation:   "vertical",
			Timestamp:       gotool.TimeStamptoDateTime(t),
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
	Timestamp       string      `json:"@timestamp"`
}
type GameBetResult struct {
	TransactionID   string `json:"transactionId"`
	TransactionType string `json:"transactionType"`
	GameCode        string `json:"gameCode"`
	RoundID         string `json:"roundId"`
	Country         string `json:"country"`
	BetID           int    `json:"betId"`
	Value           int64  `json:"value"`
	FinishTime      int64  `json:"finishTime"`
}
type GameWinResult struct {
	TransactionID   string `json:"transactionId"`
	TransactionType string `json:"transactionType"`
	BetID           int    `json:"betId"`
	Win             string `json:"win"`
	Seake           string `json:"seake"`
	WinMultiplier   string `json:"winMultiplier"`
	Insurcnce       string `json:"insurcnce"`
	PL              string `json:"pl"`
	FinishTime      int64  `json:"finishTime"`
	GameCode        string `json:"gameCode"`
	Value           int64  `json:"value"`
	RoundID         string `json:"roundId"`
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
