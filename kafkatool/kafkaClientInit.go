package kafkatool

import (
	"strings"
	"sync"
	"time"

	LogTool "github.com/adimax2953/log-tool"
	"github.com/segmentio/kafka-go"
)

var (
	kafkaConnManager     *KafkaConnManager
	kafkaConnManagerOnce sync.Once
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
		Brokers:  host,
		GroupID:  groupID,
		Topic:    topic,
		MaxBytes: 1e6,             // 1MB,  kafka-go v0.4.46 版本這個功能會無效
		MinBytes: 1e5,             // 100KB
		MaxWait:  5 * time.Second, // 5秒拉一次
	})
}

// NewDLQReader - 建立NewDLQReader
// hosts - kafka host list, 可多組, 用逗號分開.
// e.g.:100.0.0.1:9092,100.0.0.2:9092
func NewDLQReader(hosts, topic, groupID string) *kafka.Reader {
	LogTool.LogInfof("newReader", "topic: %s, groupID: %s", topic, groupID)
	host := strings.Split(hosts, ",")
	for i := range host {
		host[i] = strings.TrimSpace(host[i])
	}
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:  host,
		GroupID:  groupID,
		Topic:    topic,
		MaxBytes: 1e6,              // 1M
		MinBytes: 1e6,              // 1M
		MaxWait:  10 * time.Second, // 10秒拉一次
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

func GetKafkaConnManager() *KafkaConnManager {
	kafkaConnManagerOnce.Do(func() {
		kafkaConnManager = &KafkaConnManager{
			readers: make([]*kafka.Reader, 0),
			writers: make([]*kafka.Writer, 0),
		}
	})
	return kafkaConnManager
}

type KafkaConnManager struct {
	readers []*kafka.Reader
	writers []*kafka.Writer
	mu      sync.Mutex
}

func (kch *KafkaConnManager) AddReader(reader *kafka.Reader) {
	kch.mu.Lock()
	defer kch.mu.Unlock()
	kch.readers = append(kch.readers, reader)
}

func (kch *KafkaConnManager) AddWriter(writer *kafka.Writer) {
	kch.mu.Lock()
	defer kch.mu.Unlock()
	kch.writers = append(kch.writers, writer)
}

func (kch *KafkaConnManager) CloseReader() {
	kch.mu.Lock()
	defer kch.mu.Unlock()
	var wg sync.WaitGroup
	for _, r := range kch.readers {
		if r == nil {
			continue
		}
		wg.Add(1)
		go func(r *kafka.Reader) {
			defer wg.Done()
			err := r.Close()
			if err != nil {
				LogTool.LogError("CloseReaders", "reader close error:", err)
			}
		}(r)
	}
	wg.Wait()
}

func (kch *KafkaConnManager) CloseWriter() {
	kch.mu.Lock()
	defer kch.mu.Unlock()
	var wg sync.WaitGroup
	for _, w := range kch.writers {
		if w == nil {
			continue
		}
		wg.Add(1)
		go func(w *kafka.Writer) {
			defer wg.Done()
			err := w.Close()
			if err != nil {
				LogTool.LogError("CloseWriters", "writer close error:", err)
			}
		}(w)
	}
	wg.Wait()
}
