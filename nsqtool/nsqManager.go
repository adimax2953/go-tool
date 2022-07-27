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

// Config - Represents a Configuration
type Config struct {
	NSQ struct {
		Lookups []string `yaml:"nsqlookups"`
		NSQDs   []string `yaml:"nsqds"`
		NSQD    string   `yaml:"nsqd"`
	} `yaml:"nsq"`
}

func InitializeConsumer(nsqconfig *Config, topic, channel string, back func(m *nsq.Message) error) {
	config := nsq.NewConfig()
	{
		config.MaxInFlight = 8
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

	err = c.ConnectToNSQDs(nsqconfig.NSQ.NSQDs)
	if err != nil {
		LogTool.LogFatal("ConnectToNSQDs Error ", err)
	}

	err = c.ConnectToNSQLookupds(nsqconfig.NSQ.Lookups)
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
	LogTool.LogInfo("NSQ consumer started: %v", topic)
	<-c.StopChan
}

func InitializePublisher(nsqconfig *Config) {
	var err error
	config := nsq.NewConfig()
	{
		config.MaxInFlight = 8
		config.HeartbeatInterval = 10
		config.DefaultRequeueDelay = 0
		config.MaxBackoffDuration = time.Millisecond * 50
	}

	producer, err = nsq.NewProducer(nsqconfig.NSQ.NSQD, config)
	if err != nil {
		LogTool.LogFatal("NewProducer Error ", err)
	}
}

func Send(topic string, msg []byte) error {
	if producer == nil {
		return nil
	}
	return producer.Publish(topic+"#ephemeral", msg)
}
