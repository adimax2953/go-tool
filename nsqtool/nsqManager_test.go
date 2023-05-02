package nsqtool_test

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"testing"

	gotool "github.com/adimax2953/go-tool"
	nsqtool "github.com/adimax2953/go-tool/nsqtool"
	logtool "github.com/adimax2953/log-tool"
	"github.com/nsqio/go-nsq"
)

func Test_SendtoNSQ(t *testing.T) {

	nsqConfig := &nsqtool.NsqConfig{
		Lookups: []string{"192.168.10.184:4161", "192.168.10.185:4161"},
		NSQDs:   []string{"192.168.10.184:4150", "192.168.10.185:4150"},
		NSQD:    "192.168.10.184:4150",
		//NSQD:    "192.168.10.185:4150",
	}

	go nsqtool.InitializeConsumer(nsqConfig, "test", "", NsqunPackTest)
	nsqtool.InitializePublisher(nsqConfig)

	//nsqtool.Send("test", []byte("test山豬"))

	// Graceful shutdown -
	ch := make(chan os.Signal, 1)
	signal.Notify(ch,
		// kill -SIGINT XXXX 或 Ctrl+c
		os.Interrupt,
		syscall.SIGINT, // register that too, it should be ok
		// kill -SIGTERM XXXX
		syscall.SIGTERM,
	)
	s := <-ch
	log.Printf("s...%v\n", s)
}

// NsqunPackTest -
func NsqunPackTest(m *nsq.Message) error {
	defer gotool.RecoverPanic()
	logtool.LogDebug("get", string(m.Body))
	return nil
}
