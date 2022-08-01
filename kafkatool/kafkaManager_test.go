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
	config.CreateTopic("test111")
	config.CreateConn("test111")
	config.WriteMessages("test111", "one", "two", "three")

}
