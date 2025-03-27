package rmqtool

import (
	"context"
	"time"

	LogTool "github.com/adimax2953/log-tool"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"github.com/apache/rocketmq-client-go/v2/rlog"
)

var (
	p rocketmq.Producer
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

func SetRLogLevelToError() {
	rlog.SetLogLevel("error")
}

func InitializePublisher(config *RmqConfig) *Rmq {
	var err error
	p, err = rocketmq.NewProducer(
		producer.WithNameServer(config.NameServers),
		producer.WithRetry(2),
		producer.WithSendMsgTimeout(10*time.Second),
	)
	if err != nil {
		LogTool.LogFatal("NewProducer Error ", err)
	}

	err = p.Start()
	if err != nil {
		LogTool.LogFatal("start produce Error ", err)
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

func (rmq *Rmq) SendAsync(msg *RmqMsg) error {
	message := primitive.NewMessage(msg.Topic, msg.Body)
	message.WithTag(msg.Tag)
	message.WithKeys(msg.Keys)
	err := p.SendAsync(context.TODO(), func(ctx context.Context, result *primitive.SendResult, err error) {
		if err != nil {
			LogTool.LogErrorf("RocketMQ", "send async message error: %s", err)
		}
	}, message)
	return err
}

func (rmq *Rmq) SendOneWay(msg *RmqMsg) error {
	message := primitive.NewMessage(msg.Topic, msg.Body)
	message.WithTag(msg.Tag)
	message.WithKeys(msg.Keys)
	err := p.SendOneWay(context.TODO(), message)
	return err
}
