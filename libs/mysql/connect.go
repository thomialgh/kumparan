package mysql

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// DBConfig -
type DBConfig struct {
	Username string
	Password string
	Host     string
	Port     string
}

// ConnectToDB -
func ConnectToDB(conf *DBConfig) (*gorm.DB, error) {
	uri := fmt.Sprintf("%s:%s@tcp(%s:%s)/kumparan?charset=utf8mb4&parseTime=True&loc=Local", conf.Username, conf.Password, conf.Host, conf.Port)
	return gorm.Open("mysql", uri)
}
