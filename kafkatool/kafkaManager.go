package kafkatool

import (
	"net"
	"strconv"

	logtool "github.com/adimax2953/log-tool"
	"github.com/segmentio/kafka-go"
)

// KafkaConfig - Represents a Configuration
type KafkaConfig struct {
	Network string `yaml:"network"`
	Address string `yaml:"adress"`
}

func InitializeConsumer() {

}

func InitializePublisher() {

}

// CreateTopic -建立topic
func (config *KafkaConfig) CreateTopic(topic string) {
	conn, err := kafka.Dial(config.Network, config.Address)
	if err != nil {
		logtool.LogFatal(err.Error())
	}
	defer conn.Close()

	controller, err := conn.Controller()
	if err != nil {
		logtool.LogFatal(err.Error())
	}
	var controllerConn *kafka.Conn
	controllerConn, err = kafka.Dial("tcp", net.JoinHostPort(controller.Host, strconv.Itoa(controller.Port)))
	if err != nil {
		logtool.LogFatal(err.Error())
	}
	defer controllerConn.Close()

	topicConfigs := []kafka.TopicConfig{
		{
			Topic:             topic,
			NumPartitions:     1,
			ReplicationFactor: 1,
		},
	}

	err = controllerConn.CreateTopics(topicConfigs...)
	if err != nil {
		logtool.LogFatal(err.Error())
	}
}
