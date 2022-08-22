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

var roundID string = "0000000264W"
var trandID string = "0000000264W"

func Test_SendtoKafka(t *testing.T) {

	config := &kafkatool.KafkaConfig{
		Address:           "192.168.10.151:9092",
		Network:           "tcp",
		NumPartition:      0,
		ReplicationFactor: 1,
	}

	config.CreateTopic("USS-Game", 10)
	c = *config
	count := 100000

	mlist := make([]kafkatool.WriteData, count)
	for i := 0; i < count; i++ {

		t := time.Now().Unix()
		d, h := gotool.DateTimeFromTimeStamp(t)
		trandID = gotool.Base62Increment(trandID)
		roundID = gotool.Base62Increment(roundID)
		var bonus int64 = 0
		quantity := decimal.NewFromInt(bonus).Mul(decimal.NewFromFloat(0.04))
		var fee int64 = gotool.Str2int64(quantity.String())
		var value int64 = 10000 + bonus - fee

		gamebet := &GameBetResult{
			TransactionID:   gotool.Base62Increment(trandID),
			RoundID:         gotool.Base62Increment(roundID),
			TransactionType: "refund",
			GameCode:        "THUSS",
			BetID:           1,
			Country:         "CDN",
			Value:           value,
			FinishTime:      t,
		}

		gamelog := &GameLog{
			TransactionID:   gotool.Base62Increment(trandID),
			RoundID:         gotool.Base62Increment(roundID),
			GameCode:        "THUSS",
			Value:           value,
			Bonus:           bonus,
			Fee:             fee,
			Country:         "THD",
			TransactionType: "refund",
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
		//go c.WriteMessagesKeyValue("USS-Game", ms)
		mlist[i] = *m
	}
	c.WriteMessagesKeyValueList("USS-Game", mlist)

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

type GameLog struct {
	TransactionID   string      `json:"transactionId"`
	RoundID         string      `json:"roundId"`
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
	BetID         int    `json:"betId"`
	Win           string `json:"win"`
	Seake         string `json:"seake"`
	WinMultiplier string `json:"winMultiplier"`
	Insurcnce     string `json:"insurcnce"`
	PL            string `json:"pl"`
}

type GameResult struct {
	GameName      string `json:"gameName"`
	RoundID       string `json:"roundId"`
	Date          string `json:"date"`
	FinishTime    int64  `json:"finishTime"`
	TotalWin      string `json:"totalWin"`
	TotalSeake    string `json:"totalSeake"`
	Totalnsurcnce string `json:"totalInsurcnce"`
	TotalPL       string `json:"totalPL"`
	Lottery       int    `json:"lottery"`
}
