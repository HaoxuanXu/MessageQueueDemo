package queue

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/BurntSushi/toml"
	_ "github.com/lib/pq"
	"github.come/HaoxuanXu/MessageQueueDemo/util"
)

func GetQueue(queueConfig QueueConnectionConfig) DemoMessageQueue {
	queue := DemoMessageQueue{
		config: queueConfig,
	}
	queue.init()
	return queue
}

type QueueConnectionConfig struct {
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	User     string `toml:"user"`
	Password string `toml:"password"`
	Dbname   string `toml:"dbname"`
}

func LoadConfig(path string, config *QueueConnectionConfig) {
	if _, err := toml.DecodeFile(path, config); err != nil {
		log.Fatalln("Reading config failed", err)
	}
}

type DemoMessageQueue struct {
	config    QueueConnectionConfig
	dbHandler *sql.DB
}

func (queue *DemoMessageQueue) setupDB() {
	setupSQL := util.LoadSQL("setup.sql")
	_, err := queue.dbHandler.Exec(setupSQL)
	util.CheckError(err)

}

func (queue *DemoMessageQueue) init() {
	connectionString := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		queue.config.Host,
		queue.config.Port,
		queue.config.User,
		queue.config.Password,
		queue.config.Dbname,
	)
	dbHandler, err := sql.Open("postgres", connectionString)
	util.CheckError(err)
	err = dbHandler.Ping()
	util.CheckError(err)
	queue.dbHandler = dbHandler

	queue.setupDB()
}

func (queue *DemoMessageQueue) Select() *sql.Rows {
	queryString := util.LoadSQL("pull.sql")
	tx, err := queue.dbHandler.Begin()
	util.CheckError(err)
	rows, err := queue.dbHandler.Query(queryString)
	if err != nil {
		tx.Rollback()
	} else {
		_ = tx.Commit()
	}

	return rows
}

func (queue *DemoMessageQueue) Update(id int, consumer_name, work_status, end_work_status string) error {
	isolationString := util.LoadSQL("isolation.sql")
	updateString := util.LoadSQL("update.sql")
	tx, err := queue.dbHandler.Begin()
	util.CheckError(err)

	_, err = tx.Exec(isolationString)
	util.CheckError(err)
	_, err = tx.Exec(updateString, work_status, consumer_name, id, end_work_status)
	if err != nil {
		tx.Rollback()
	} else {
		err = tx.Commit()
	}
	return err
}

func (queue *DemoMessageQueue) Ingest(producerSource string) (sql.Result, error) {
	ingestString := util.LoadSQL("ingest.sql")
	tx, err := queue.dbHandler.Begin()
	util.CheckError(err)

	res, err := tx.Exec(ingestString, producerSource, "open")
	if err != nil {
		fmt.Println(err)
		tx.Rollback()
	} else {
		err = tx.Commit()
	}
	return res, err
}
