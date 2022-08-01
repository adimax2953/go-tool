package kafkatool_test

import (
	"testing"

	"github.com/adimax2953/go-tool/kafkatool"
)

func Test_SendtoKafka(t *testing.T) {

	config := &kafkatool.KafkaConfig{
		Address:           "192.168.56.1:9092",
		Network:           "tcp",
		NumPartition:      1,
		ReplicationFactor: 1,
	}
	//config.CreateTopic("test1112")
	//config.CreateConn("test1")
	config.WriteMessages("test3", "da", "da", "der", "ma", "te", "sen")
	//config.ReadMessages("test1")
	config.GetTopic()
	config.DelTopic(config.GetTopic()...)
}
