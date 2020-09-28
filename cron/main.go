package main

import (
	"kumparan/cron/config"
	"kumparan/cron/pkg"
	"kumparan/libs/kafka"
	msql "kumparan/libs/mysql"
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
	config.KafkaReadConfig = kafka.InitReader("kafka:9092", "kumparan", "cons")
	<-done
	if gerr != nil {
		log.Fatal(gerr)
	}
	pkg.RunCron()
}
