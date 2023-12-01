package kafkatool

import (
	LogTool "github.com/adimax2953/log-tool"
	"github.com/segmentio/kafka-go"
	"strings"
	"time"
)

// NewReader - 建立Reader
// hosts - kafka host list, 可多組, 用逗號分開.
// e.g.:100.0.0.1:9092,100.0.0.2:9092
func NewReader(hosts, topic, groupID string) *kafka.Reader {
	LogTool.LogInfof("newReader", "topic: %s, groupID: %s", topic, groupID)
	host := strings.Split(hosts, ",")
	for i := range host {
		host[i] = strings.TrimSpace(host[i])
	}
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:          host,
		GroupID:          groupID,
		Topic:            topic,
		MaxBytes:         1e6,             // 1MB,  kafka-go v0.4.46 版本這個功能會無效
		MinBytes:         1e5,             // 100KB
		ReadBatchTimeout: 5 * time.Second, // 5秒拉一次
	})
}

// NewWriter - 建立Writer, 可多組, 用逗號分開
// e.g.:100.0.0.1:9092,100.0.0.2:9092
func NewWriter(hosts, topic string) *kafka.Writer {
	LogTool.LogInfof("newWriter", "topic: %s", topic)
	host := strings.Split(hosts, ",")
	for i := range host {
		host[i] = strings.TrimSpace(host[i])
	}
	return &kafka.Writer{
		Addr:     kafka.TCP(host...),
		Topic:    topic,
		Balancer: &kafka.Murmur2Balancer{},
	}
}
