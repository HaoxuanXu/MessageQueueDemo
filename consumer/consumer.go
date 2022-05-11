package consumer

import (
	"errors"
	"fmt"
	"time"

	"github.come/HaoxuanXu/MessageQueueDemo/queue"
	"github.come/HaoxuanXu/MessageQueueDemo/util"
)

func GetConsumer(name string, config queue.QueueConnectionConfig) DemoConsumer {
	return DemoConsumer{
		name:          name,
		workStatusMap: util.GetWorkStatusMapping(),
		queue:         queue.GetQueue(config),
	}
}

type DemoConsumer struct {
	name          string
	queue         queue.DemoMessageQueue
	workStatusMap *util.WorkStatus
}

func (consumer DemoConsumer) TakeWork() int {
	rows := consumer.queue.Select()
	var id int
	var err error
	for rows.Next() {
		err = rows.Scan(&id)
		util.CheckError(err)

		// attempt to grab the work
		err = consumer.queue.Update(id, consumer.name, consumer.workStatusMap.Occupied, consumer.workStatusMap.Open)
		if err != nil {
			continue
		} else {
			fmt.Printf("%s successfully grab job %d\n", consumer.name, id)
			break
		}
	}
	if err != nil {
		id = 0
	}
	return id
}

func (consumer DemoConsumer) Work(id, timeToFinished int) bool {
	var success bool
	if id != 0 {
		fmt.Printf("%s working on work %d\n", consumer.name, id)
		time.Sleep(time.Duration(timeToFinished) * time.Second)
		fmt.Printf("%s finished Work %d\n", consumer.name, id)
		success = true
	} else {
		success = false
	}
	return success
}

func (consumer DemoConsumer) ReportWorkFinished(id int, isSuccess bool) error {
	var err error = errors.New("error template")
	if isSuccess {
		for err != nil {
			err = consumer.queue.Update(id, consumer.name, consumer.workStatusMap.Finished, consumer.workStatusMap.Occupied)
			if err != nil {
				fmt.Printf("%s failed to report work %d finished\n", consumer.name, id)
			} else {
				fmt.Printf("%s successfully reported work %d finished\n", consumer.name, id)
			}
		}
	}
	return err
}
