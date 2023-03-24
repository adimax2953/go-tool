package rmqtool

import (
	"context"
	LogTool "github.com/adimax2953/log-tool"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"github.com/apache/rocketmq-client-go/v2/rlog"
	"sync"
)

var (
	p    rocketmq.Producer
	once sync.Once
)

type RmqConfig struct {
	NameServers []string
}

type Rmq struct {
	Producer *rocketmq.Producer
}

type RmqMsg struct {
	Topic string
	Tag   string
	Keys  []string
	Body  []byte
}

func InitializePublisher(config *RmqConfig) *Rmq {
	once.Do(func() {
		rlog.SetLogLevel("error")
	})
	p, _ = rocketmq.NewProducer(
		producer.WithNameServer(config.NameServers),
		//producer.WithNameServer([]string{"0.0.0.0:9876"}),
		producer.WithRetry(2),
	)

	err := p.Start()
	if err != nil {
		if err != nil {
			LogTool.LogFatal("NewProducer Error ", err)
		}
	}
	return &Rmq{
		Producer: &p,
	}
}

func (rmq *Rmq) Send(msg *RmqMsg) error {
	message := primitive.NewMessage(msg.Topic, msg.Body)
	message.WithTag(msg.Tag)
	message.WithKeys(msg.Keys)
	_, err := p.SendSync(context.TODO(), message)
	if err != nil {
		LogTool.LogErrorf("RocketMQ", "send message error: %s", err)
	}
	return err
}
