package rmqtool

import (
	"context"
	LogTool "github.com/adimax2953/log-tool"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"os"
	"os/signal"
	"syscall"
)

type ConsumerMode int

const (
	SingleMode ConsumerMode = iota // 同樣group中的的consumer 只會有一個consumer收到訊息
	PubSubMode                     // pub/sub mode 用於水平擴展 group 中的所有consumer 都會收到同一條消息
)

type ConsumerConfig struct {
	Topic        string
	Tag          string
	Group        string // consumer group
	Order        bool   // fifo message
	ConsumerMode ConsumerMode
	MsgHandler   func(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error)
}

func InitializeConsumer(config *RmqConfig, consumerConfig *ConsumerConfig) {
	var err error
	c, err := rocketmq.NewPushConsumer(
		consumer.WithGroupName(gerGroupName(consumerConfig)),
		consumer.WithNsResolver(primitive.NewPassthroughResolver(config.NameServers)),
		consumer.WithConsumerOrder(consumerConfig.Order), // 是否啟用有序消費
		consumer.WithConsumeConcurrentlyMaxSpan(10),
		consumer.WithRetry(2),
		consumer.WithConsumerModel(getConsumerMode(consumerConfig.ConsumerMode)),
	)

	if err != nil {
		LogTool.LogFatalf("RockerMQ", "NewPushConsumer error: %s", err)
	}

	err = c.Subscribe(consumerConfig.Topic, getMessageSelector(consumerConfig), consumerConfig.MsgHandler)
	if err != nil {
		LogTool.LogFatalf("RockerMQ", "subscribe error: %s", err)
	}

	err = c.Start()
	if err != nil {
		LogTool.LogFatalf("RockerMQ", "start consumer error: %s", err)
	}

	// Graceful shutdown
	signalChan := make(chan os.Signal)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	//register for interupt (Ctrl+C) and SIGTERM (docker)
	<-signalChan
	c.Unsubscribe(consumerConfig.Topic)
	c.Shutdown()
}

func gerGroupName(config *ConsumerConfig) string {
	if config.Group == "" {
		LogTool.LogFatal("consumer group name is required")
	}
	return config.Group
}

func getMessageSelector(config *ConsumerConfig) consumer.MessageSelector {
	result := consumer.MessageSelector{}
	if config.Tag != "" {
		result.Type = consumer.TAG
		result.Expression = config.Tag
	}
	return result
}

func getConsumerMode(mode ConsumerMode) consumer.MessageModel {
	switch mode {
	case PubSubMode:
		return consumer.BroadCasting
	default:
		return consumer.Clustering
	}
}
