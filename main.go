package main

import (
	msql "kumparan/libs/mysql"
	"kumparan/models"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	var maxRetry int
	var err error
	var db *gorm.DB
	mysqlConf := msql.DBConfig{
		Username: "root",
		Password: "kumparan-test",
		Host:     "mysql",
		Port:     "3306",
	}
	for {
		db, err = msql.ConnectToDB(&mysqlConf)
		if err != nil {
			if maxRetry == 10 {
				break
			}
			time.Sleep(1 * time.Second)
			maxRetry++
			continue
		} else {
			break
		}
	}

	db.AutoMigrate(models.Articles{})
}
