package main

import (
	"context"
	"github.com/LeakIX/l9format"
	"github.com/Shopify/sarama"
	"log"
	"net"
	"time"
)

type KafkaOpenPlugin struct {
	l9format.ServicePluginBase
}

func main() {}
func New() l9format.ServicePluginInterface {
	return KafkaOpenPlugin{}
}

func (KafkaOpenPlugin) GetVersion() (int, int, int) {
	return 0, 0, 1
}

func (KafkaOpenPlugin) GetProtocols() []string {
	return []string{"kafka"}
}

func (KafkaOpenPlugin) GetName() string {
	return "KafkaOpenPlugin"
}

func (KafkaOpenPlugin) GetStage() string {
	return "open"
}

// Get info
func (plugin KafkaOpenPlugin) Run(ctx context.Context, event *l9format.L9Event, pluginOptions map[string]string) (leak l9format.L9LeakEvent, hasLeak bool) {
	config := sarama.NewConfig()
	deadline, hasDeadline := ctx.Deadline()
	if hasDeadline {
		config.Net.DialTimeout = deadline.Sub(time.Now())
	} else {
		config.Net.DialTimeout = 5 * time.Second
	}
	config.Consumer.Return.Errors = true

	//kafka end point
	brokers := []string{net.JoinHostPort(event.Ip, event.Port)}
	cluster, err := sarama.NewConsumer(brokers, config)
	if err != nil {
		log.Println(err)
		return leak, false
	}
	defer cluster.Close()
	topics, err := cluster.Topics()
	if err != nil || len(topics) < 1 {
		log.Println(err)
		return leak, false
	}
	leak.Data = "NoAuth\n"
	for _, topic := range topics {
		leak.Data += "Found topic " + topic + "\n"
	}
	return leak, true
}