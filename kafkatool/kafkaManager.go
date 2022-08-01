package kafkatool

import (
	"net"
	"strconv"
	"time"

	logtool "github.com/adimax2953/log-tool"
	"github.com/segmentio/kafka-go"
)

// KafkaConfig - Represents a Configuration
type KafkaConfig struct {
	Network           string `yaml:"network"`
	Address           string `yaml:"adress"`
	NumPartition      int    `yaml:"numPartition"`
	ReplicationFactor int    `yaml:"replicationFactor"`
	Client            *kafka.Conn
}

func InitializeConsumer() {

}

func InitializePublisher() {

}

// CreateTopic -建立topic 1.topic 2.NumPartition 3.ReplicationFactor
func (config *KafkaConfig) CreateTopic(topic string, num ...int) {

	//初始數值
	numPartition := config.NumPartition
	replicationFactor := config.ReplicationFactor
	if len(num) > 0 {
		numPartition = num[0]
	}
	if len(num) > 1 {
		replicationFactor = num[1]
	}

	conn, err := kafka.Dial(config.Network, config.Address)
	if err != nil {
		logtool.LogFatal("CreateTopic0", err.Error())
	}
	defer conn.Close()

	controller, err := conn.Controller()
	if err != nil {
		logtool.LogFatal(err.Error())
	}
	var controllerConn *kafka.Conn

	controllerConn, err = kafka.Dial(config.Network, net.JoinHostPort(controller.Host, strconv.Itoa(controller.Port)))
	if err != nil {
		logtool.LogFatal(err.Error())
	}
	defer controllerConn.Close()

	topicConfigs := []kafka.TopicConfig{
		{
			Topic:             topic,
			NumPartitions:     numPartition,
			ReplicationFactor: replicationFactor,
		},
	}

	err = controllerConn.CreateTopics(topicConfigs...)
	if err != nil {
		logtool.LogFatal(err.Error())
	}
	config.Client = controllerConn

	logtool.LogInfo("Kafka CreateTopic ", topic)
}

// CreateConn - 建立對Topic的連線
func (config *KafkaConfig) CreateConn(topic string, num ...int) *kafka.Conn {
	logtool.LogInfo("Kafka CreateConn ", config.Network, config.Address)

	conn, err := kafka.Dial(config.Network, config.Address)
	if err != nil {
		logtool.LogFatal(err.Error())
	}
	defer conn.Close()
	controller, err := conn.Controller()
	if err != nil {
		logtool.LogFatal(err.Error())
	}
	var connLeader *kafka.Conn
	connLeader, err = kafka.Dial(config.Network, net.JoinHostPort(controller.Host, strconv.Itoa(controller.Port)))
	if err != nil {
		logtool.LogFatal(err.Error())
	}
	config.Client = connLeader
	defer connLeader.Close()
	return config.Client
}

// WriteMessages - 發送訊息到Topic
func (config *KafkaConfig) WriteMessages(topic string, value ...string) {
	count := len(value)
	if count == 0 {
		return
	}
	mlist := make([]kafka.Message, count)

	for k, v := range value {
		mlist[k] = kafka.Message{Value: []byte(v)}
	}

	config.Client.SetWriteDeadline(time.Now().Add(10 * time.Second))

	_, err := config.Client.WriteMessages(mlist...)
	if err != nil {
		logtool.LogError("failed to write messages ", err)
	}
	logtool.LogInfo("Kafka WriteMessages ", mlist)
}
