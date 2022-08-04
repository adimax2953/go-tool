package kafkatool_test

import (
	"testing"

	gotool "github.com/adimax2953/go-tool"
	"github.com/adimax2953/go-tool/kafkatool"
)

var c kafkatool.KafkaConfig

func Test_SendtoKafka(t *testing.T) {

	config := &kafkatool.KafkaConfig{
		Address:           "192.168.56.1:9092",
		Network:           "tcp",
		NumPartition:      0,
		ReplicationFactor: 1,
	}
	config.CreateTopic("test3", 10)
	c = *config
	// m := map[string]string{}
	// for i := 0; i < 10000; i++ {
	// 	m[gotool.IntToStr(i)+"@player"] = "value " + gotool.IntToStr(i)
	// }
	// config.WriteMessagesKeyValue("test3", m)
	//config.ReadMessages("test3", "1")

	//config.GetTopic()
	//config.DelTopic(config.GetTopic()...)
	wp := gotool.NewWorkPool(1)
	p, _ := wp.NewWorkPoolWithFunc(100, func(i interface{}) {
		test(i)
	})
	for i := 0; i < 10000; i++ {
		p.Invoke(i)
	}
	defer wp.Release()
	defer p.Release()

	// s := make([]string, 10000)
	// for i := 0; i < 10000; i++ {
	// 	s[i] = "value " + gotool.IntToStr(i)
	// }
	// config.WriteMessages("test3", s...)

	//config.CreateTopic("test1112")
	//config.CreateConn("test12")
	//config.WriteMessages("test3", "da", "da", "der", "ma", "te", "sen")
}
func test(i interface{}) {
	m := map[string]string{}

	m[gotool.IntToStr(i.(int))+"@once"] = gotool.IntToStr(i.(int))
	c.WriteMessagesKeyValue("test3", m)

}
