package producer

import (
	"fmt"

	"github.come/HaoxuanXu/MessageQueueDemo/queue"
)

func GetProducer(name string, config queue.QueueConnectionConfig) DemoProducer {
	return DemoProducer{
		name:  name,
		queue: queue.GetQueue(config),
	}
}

type DemoProducer struct {
	name  string
	queue queue.DemoMessageQueue
}

func (producer DemoProducer) PushWork() {
	_, err := producer.queue.Ingest(producer.name)
	if err != nil {
		fmt.Printf("Work pushing for producer %s failed: %s\n", producer.name, err)
	} else {
		fmt.Printf("%s pushed a work\n", producer.name)
	}
}
