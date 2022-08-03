package kafkatool_test

import (
	"testing"

	gotool "github.com/adimax2953/go-tool"
	"github.com/adimax2953/go-tool/kafkatool"
)

func Test_SendtoKafka(t *testing.T) {

	config := &kafkatool.KafkaConfig{
		Address:           "192.168.56.1:9092",
		Network:           "tcp",
		NumPartition:      0,
		ReplicationFactor: 1,
	}
	config.CreateTopic("test3", 10)

	//config.NewClient("test100")
	//config.WriteMessages2("test100")
	/*config.WriteMessages("test100", "99")
	config.WriteMessages("test100", "de")
	config.WriteMessages("test100", "do")
	config.WriteMessages("test100", "dc")
	config.WriteMessages("test100", "oh")*/

	// m := map[string]string{}
	// for i := 0; i < 10000; i++ {
	// 	m[gotool.IntToStr(i)+"@player"] = "value " + gotool.IntToStr(i)
	// }
	// config.WriteMessagesKeyValue("test001", m)

	// for i := 0; i < 10000; i++ {
	// 	m := map[string]string{}

	// 	m[gotool.IntToStr(i)+"@once"] = gotool.IntToStr(i)
	// 	config.WriteMessagesKeyValue("test001", m)
	// }

	s := make([]string, 10000)
	for i := 0; i < 10000; i++ {
		s[i] = "value " + gotool.IntToStr(i)
	}
	config.WriteMessages("test3", s...)

	//config.CreateTopic("test1112")
	//config.CreateConn("test12")
	//config.WriteMessages("test3", "da", "da", "der", "ma", "te", "sen")
	//config.ReadMessages("test3")
	//config.GetTopic()
	//config.DelTopic(config.GetTopic()...)
}
