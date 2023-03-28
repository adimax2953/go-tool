package rmqtool

import (
	"context"
	"encoding/json"
	"github.com/apache/rocketmq-client-go/v2/admin"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var (
	nameServers = []string{"103.103.81.12:9876"}
)

func TestSend(t *testing.T) {
	config := new(RmqConfig)
	config.NameServers = nameServers
	m := map[string]string{
		"foo": "bar111",
	}
	body, err := json.Marshal(m)
	assert.Empty(t, err)
	instance := InitializePublisher(config)
	err = instance.Send(&RmqMsg{
		Topic: "test",
		Tag:   "tag_1",
		Keys:  []string{"test123"},
		Body:  body,
	})
	assert.Empty(t, err)
}

func TestConsumer(t *testing.T) {
	config := new(RmqConfig)
	config.NameServers = nameServers
	consumerConfig := new(ConsumerConfig)
	consumerConfig.Topic = "test"
	consumerConfig.Tag = "tag_1"
	consumerConfig.MsgHandler = func(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
		for i := range msgs {
			//fmt.Printf("subscribe callback: %v \n", msgs[i])
			t.Logf("message: %s, \n", msgs[i].Body)
		}
		return consumer.ConsumeSuccess, nil
	}
	InitializeConsumer(config, consumerConfig)
	time.Sleep(10 * time.Second)
}

func TestCreateTopic(t *testing.T) {
	broker := "103.103.81.12:10911"

	checkAdmin, err := admin.NewAdmin(admin.WithResolver(primitive.NewPassthroughResolver(nameServers)))
	assert.Empty(t, err)
	topics := []string{
		"5state",
		"7betzone",
		"4join",
		"5leave",
		"3bet",
		"3win",
		"5other",
		"8sendGame",
		"7sendAll",
		"9askrobots",
		"7records",
		"6record",
		"8opcommit",
		"7ctimout",
		"7dealing",
		"8matching",
		"6commit",
	}

	for _, v := range topics {
		err = checkAdmin.CreateTopic(
			context.TODO(),
			admin.WithBrokerAddrCreate(broker),
			admin.WithTopicCreate(v),
		)
		assert.Empty(t, err)
	}

}
