package kafkatool

import (
	"context"
	"errors"
	"net"
	"strconv"
	"strings"
	"time"

	gotool "github.com/adimax2953/go-tool"
	logtool "github.com/adimax2953/log-tool"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/compress"
)

// KafkaConfig - Represents a Configuration
type KafkaConfig struct {
	Network           string `yaml:"network"`
	Address           string `yaml:"adress"`
	NumPartition      int    `yaml:"numPartition"`
	ReplicationFactor int    `yaml:"replicationFactor"`
	Conn              *kafka.Conn
}

type WriteData struct {
	Key   string
	Value string
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
		logtool.LogFatal("CreateTopic Dial Error", err.Error())
	}
	defer conn.Close()

	/*controller, err := conn.Controller()
	if err != nil {
		logtool.LogFatal(err.Error())
	}*/
	var controllerConn *kafka.Conn
	addr := strings.Split(config.Address, ":")
	controllerConn, err = kafka.Dial(config.Network, net.JoinHostPort(addr[0], addr[1]))
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

	err = conn.CreateTopics(topicConfigs...)
	if err != nil {
		logtool.LogFatal(err.Error())
	}
	config.Conn = controllerConn

	logtool.LogInfo("Kafka CreateTopic ", topic)
}

// DelTopic - 刪除Topic的列表
func (config *KafkaConfig) DelTopic(topic ...string) {

	conn, err := kafka.Dial(config.Network, config.Address)
	if err != nil {
		logtool.LogFatal(err.Error())
	}
	defer conn.Close()

	controller, err := conn.Controller()
	if err != nil {
		logtool.LogError("DelTopic controller Error ", err)
	}

	conn, err = kafka.Dial(config.Network, net.JoinHostPort(controller.Host, strconv.Itoa(controller.Port)))
	if err != nil {
		logtool.LogError("DelTopic Dial Error", err)
	}

	conn.SetDeadline(time.Now().Add(10 * time.Second))

	if err := conn.DeleteTopics(topic...); err != nil {
		logtool.LogError("DelTopic Delete Error ", err)
	}
	logtool.LogInfo("Kafka DelTopic ", topic)

}

// GetTopic - 取得Topic的列表
func (config *KafkaConfig) GetTopic() []string {

	conn, err := kafka.Dial(config.Network, config.Address)
	if err != nil {
		logtool.LogFatal(err.Error())
	}

	defer conn.Close()

	partitions, err := conn.ReadPartitions()
	if err != nil {
		logtool.LogFatal(err.Error())
	}
	count := len(partitions)

	m := make([]string, count)

	for i, p := range partitions {
		m[i] = p.Topic
	}
	for _, v := range m {
		logtool.LogInfo(v)
	}
	logtool.LogInfo("Kafka GetTopic ", m)

	return m
}

// WriteMessagesKeyValueList - 發送訊息到Topic
func (config *KafkaConfig) WriteMessagesKeyValueList(topic string, value []WriteData) {
	count := len(value)
	if count == 0 {
		logtool.LogError("WriteMessagesKeyValueList value is nil")
		return
	}
	mlist := make([]kafka.Message, count)

	w := &kafka.Writer{
		Addr:                   kafka.TCP(config.Address),
		Topic:                  topic,
		AllowAutoTopicCreation: true,
		Balancer:               &kafka.Murmur2Balancer{},
		RequiredAcks:           -1,
		BatchSize:              1048576,
		BatchBytes:             1048576,
		Compression:            compress.None,
	}
	sum := 0
	for _, mv := range value {
		mlist[sum] = kafka.Message{
			Key:   []byte(mv.Key),
			Value: []byte(mv.Value)}
		sum++
	}

	var err error
	const retries = 1
	for i := 0; i < retries; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err = w.WriteMessages(ctx, mlist...); err != nil {
			if errors.Is(err, kafka.LeaderNotAvailable) || errors.Is(err, context.DeadlineExceeded) {
				time.Sleep(time.Millisecond * 100)
				continue
			}
			logtool.LogError("WriteMessages unexpected error %v", err)
		}
	}

	if err := w.Close(); err != nil {
		logtool.LogError("failed to close writer:", err)
	}
	//logtool.LogInfo("Kafka WriteMessages ", mlist)
}

// WriteMessagesKeyValue - 發送訊息到Topic
func (config *KafkaConfig) WriteMessagesKeyValue(topic string, value map[string]string) {
	count := len(value)
	if count == 0 {
		logtool.LogError("WriteMessagesKeyValue value is nil")
		return
	}
	mlist := make([]kafka.Message, count)

	w := &kafka.Writer{
		Addr:                   kafka.TCP(config.Address),
		Topic:                  topic,
		AllowAutoTopicCreation: true,
		Balancer:               &kafka.Murmur2Balancer{},
		RequiredAcks:           -1,
		BatchSize:              1048576,
		BatchBytes:             1048576,
		Compression:            compress.None,
	}
	sum := 0
	for k, v := range value {
		mlist[sum] = kafka.Message{
			Key:   []byte(k),
			Value: []byte(v)}
		sum++
	}

	var err error
	const retries = 1
	for i := 0; i < retries; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err = w.WriteMessages(ctx, mlist...); err != nil {
			if errors.Is(err, kafka.LeaderNotAvailable) || errors.Is(err, context.DeadlineExceeded) {
				time.Sleep(time.Millisecond * 100)
				continue
			}
			logtool.LogError("WriteMessages unexpected error %v", err)
		}
	}

	if err := w.Close(); err != nil {
		logtool.LogError("failed to close writer:", err)
	}
	//logtool.LogInfo("Kafka WriteMessages ", mlist)
}

// WriteMessages - 發送訊息到Topic
func (config *KafkaConfig) WriteMessages(topic string, value ...string) {
	count := len(value)
	if count == 0 {
		logtool.LogError("WriteMessages value is nil")
		return
	}
	mlist := make([]kafka.Message, count)

	w := &kafka.Writer{
		Addr:                   kafka.TCP(config.Address),
		Topic:                  topic,
		AllowAutoTopicCreation: true,
		Compression:            compress.None,
		BatchSize:              1048576,
		BatchBytes:             1048576,
		Balancer:               &kafka.Murmur2Balancer{},
		RequiredAcks:           -1,
	}

	for k, v := range value {
		mlist[k] = kafka.Message{
			Key:   []byte(v),
			Value: []byte(v)}
	}

	var err error
	const retries = 1
	for i := 0; i < retries; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err = w.WriteMessages(ctx, mlist...); err != nil {
			if errors.Is(err, kafka.LeaderNotAvailable) || errors.Is(err, context.DeadlineExceeded) {
				time.Sleep(time.Millisecond * 250)
				continue
			}
			logtool.LogError("WriteMessages unexpected error %v", err)
		}
	}

	if err := w.Close(); err != nil {
		logtool.LogError("failed to close writer:", err)
	}
	logtool.LogInfo("Kafka WriteMessages ", mlist)
}

// ReadMessages - 接收Topic的訊息
func (config *KafkaConfig) ReadMessages(topic, groupid string) {

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:        []string{config.Address},
		GroupID:        groupid,
		Topic:          topic,
		MinBytes:       10e3,        // 10KB
		MaxBytes:       10e6,        // 10MB
		CommitInterval: time.Second, // flushes commits to Kafka every second
	})

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	i := 0
	for {
		m, err := r.ReadMessage(ctx)
		if err != nil {
			break
		}
		logtool.LogInfo("message at", gotool.IntToStr(i), string(m.Value))
		i++
	}

	if err := r.Close(); err != nil {
		logtool.LogError("failed to close reader:", err)
	}
}
