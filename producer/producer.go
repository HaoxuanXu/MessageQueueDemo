package producer

import (
	"fmt"

	"github.come/HaoxuanXu/MessageQueueDemo/queue"
)

type Producer interface {
	PushWork()
}

type DemoProducer struct {
	name  string
	queue queue.DemoMessageQueue
}

func GetProducer(name string, config queue.QueueConnectionConfig) Producer {
	return DemoProducer{
		name:  name,
		queue: queue.GetQueue(config),
	}
}

func (producer DemoProducer) PushWork() {
	_, err := producer.queue.Ingest(producer.name)
	if err != nil {
		fmt.Printf("Work pushing for producer %s failed: %s\n", producer.name, err)
	} else {
		fmt.Printf("%s pushed a work\n", producer.name)
	}
}
