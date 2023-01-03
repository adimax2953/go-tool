package nsqtool

import (
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	LogTool "github.com/adimax2953/log-tool"
	"github.com/nsqio/go-nsq"
)

var (
	producer *nsq.Producer
)

// NsqConfig - Represents a Configuration
type NsqConfig struct {
	Lookups []string `yaml:"nsqlookups"`
	NSQDs   []string `yaml:"nsqds"`
	NSQD    string   `yaml:"nsqd"`
}
type Nsq struct {
	Producer *nsq.Producer
}

func InitializeConsumer(nsqconfig *NsqConfig, topic, channel string, back func(m *nsq.Message) error) {
	config := nsq.NewConfig()
	{
		config.MaxInFlight = 80
		config.HeartbeatInterval = 10
		config.DefaultRequeueDelay = 0
		config.MaxBackoffDuration = time.Millisecond * 50
	}

	chars := []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	buf := make([]byte, 10)
	for i := 0; i < 10; i++ {
		buf[i] = chars[rand.Intn(len(chars))]
	}
	if channel == "" {
		channel = string(buf)
	}
	// test#ephemeral 後綴#ephemeral為臨時Topic或Channel
	c, err := nsq.NewConsumer(topic+"#ephemeral", channel+"#ephemeral", config)
	if err != nil {
		LogTool.LogFatal("NewConsumer Error ", err)
	}
	c.SetLoggerLevel(nsq.LogLevelError)

	defer func() {
		c.Stop()
		LogTool.LogInfo(topic+" Stats: %+v\n", c.Stats())
		LogTool.LogInfo(topic+"IsStarved: %+v\n", c.IsStarved())
	}()

	c.AddConcurrentHandlers(nsq.HandlerFunc(back), 10)

	err = c.ConnectToNSQDs(nsqconfig.NSQDs)
	if err != nil {
		LogTool.LogFatal("ConnectToNSQDs Error ", err)
	}

	err = c.ConnectToNSQLookupds(nsqconfig.Lookups)
	if err != nil {
		LogTool.LogFatal("ConnectToNSQLookupds Error ", err)
	}

	// Graceful shutdown
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	//register for interupt (Ctrl+C) and SIGTERM (docker)
	go func() {
		<-signalChan
		c.Stop()
	}()

	// start listening
	LogTool.LogInfo("NSQ consumer started topic ", topic)
	<-c.StopChan
}

func InitializePublisher(nsqconfig *NsqConfig) *Nsq {
	var err error
	config := nsq.NewConfig()
	{
		config.MaxInFlight = 8
		config.HeartbeatInterval = 10
		config.DefaultRequeueDelay = 0
		config.MaxBackoffDuration = time.Millisecond * 50
	}

	producer, err = nsq.NewProducer(nsqconfig.NSQD, config)
	if err != nil {
		LogTool.LogFatal("NewProducer Error ", err)
	}

	return &Nsq{Producer: producer}
}

func (n *Nsq) Send(topic string, msg []byte) error {

	if n == nil || n.Producer == nil {
		LogTool.LogError("producer nil")
		return nil
	}
	if err := producer.Ping(); err != nil {
		LogTool.LogError("producer ping error:%+v", err)
		return nil
	}
	// return producer.Publish(topic+"#ephemeral", msg)
	// msgCount := 1
	// responseChan := make(chan *nsq.ProducerTransaction, msgCount)
	// return producer.PublishAsync(topic+"#ephemeral", msg, responseChan)
	return producer.PublishAsync(topic+"#ephemeral", msg, nil)
}
func (n *Nsq) SendSync(topic string, msg []byte) error {
	if producer == nil {
		LogTool.LogError("producer nil")
		return nil
	}
	return producer.Publish(topic+"#ephemeral", msg)
}
