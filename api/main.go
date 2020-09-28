package main

import (
	"kumparan/api/config"
	"kumparan/api/server"
	"kumparan/libs/kafka"
	msql "kumparan/libs/mysql"
	"kumparan/libs/redis"
	"log"
	"time"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	var gerr error
	done := make(chan bool)
	mysqlConf := msql.DBConfig{
		Username: "root",
		Password: "kumparan-test",
		Host:     "mysql",
		Port:     "3306",
	}

	go func() {
		var maxRetry int
		var err error
		for {
			config.DB, err = msql.ConnectToDB(&mysqlConf)
			if err != nil {
				if maxRetry == 10 {
					gerr = err
					break
				}
				time.Sleep(1 * time.Second)
				maxRetry++
				continue
			} else {
				break
			}
		}
		done <- true
	}()
	go func() {
		config.KafkaWriterConf = kafka.InitWriter("kafka:9092", "kumparan", "kumparan")
		done <- true
	}()
	go func() {
		time.Sleep(3 * time.Second)
		if err := redis.InitRedis("redis:6379"); err != nil {
			gerr = err

		}
		done <- true
	}()
	for i := 0; i < 3; i++ {
		<-done
	}
	if gerr != nil {
		log.Fatal(gerr)
	}
	server.RunServer()
}
