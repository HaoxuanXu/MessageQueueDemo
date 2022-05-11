package main

import (
	"fmt"
	"time"

	"github.come/HaoxuanXu/MessageQueueDemo/config"
	"github.come/HaoxuanXu/MessageQueueDemo/consumer"
	"github.come/HaoxuanXu/MessageQueueDemo/producer"
	"github.come/HaoxuanXu/MessageQueueDemo/queue"
)

func main() {

	// setup message queue
	config := queue.QueueConnectionConfig{
		Host:     config.Host,
		Port:     config.Port,
		User:     config.User,
		Password: config.Password,
		Dbname:   config.Dbname,
	}
	// setup 1 producer
	producerBob := producer.GetProducer("Bob", config)

	// setup 3 consumers
	consumerJohn := consumer.GetConsumer("John", config)
	consumerMike := consumer.GetConsumer("Mike", config)
	consumerDan := consumer.GetConsumer("Dan", config)

	// run the pipeline for 10 minutes
	startTime := time.Now()

	go func() {
		for time.Since(startTime) < 3*time.Minute {
			producerBob.PushWork()
			time.Sleep(5 * time.Second)
		}
	}()

	go func() {
		for time.Since(startTime) < 3*time.Minute {
			id := consumerJohn.TakeWork()
			isSuccess := consumerJohn.Work(id, 15)
			consumerJohn.ReportWorkFinished(id, isSuccess)
		}
	}()

	go func() {
		for time.Since(startTime) < 3*time.Minute {
			id := consumerMike.TakeWork()
			isSuccess := consumerMike.Work(id, 15)
			consumerMike.ReportWorkFinished(id, isSuccess)
		}
	}()

	go func() {
		for time.Since(startTime) < 3*time.Minute {
			id := consumerDan.TakeWork()
			isSuccess := consumerDan.Work(id, 15)
			consumerDan.ReportWorkFinished(id, isSuccess)
		}
	}()

	fmt.Println("Keep main routine alive")
	time.Sleep(3 * time.Minute)

}
