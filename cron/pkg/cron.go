package pkg

import (
	"encoding/json"
	"kumparan/cron/config"
	"kumparan/libs/elastic"
	"kumparan/libs/kafka"
	"kumparan/models"
	"log"
	"time"
)

// ReadMessage -
func ReadMessage() (*models.Articles, error) {
	var arc *models.Articles
	msg, err := kafka.Cousume(config.KafkaReadConfig)
	if err != nil {
		return arc, err
	}

	if err := json.Unmarshal(msg, &arc); err != nil {
		return arc, err
	}

	return arc, nil
}

// InsertData -
func InsertData() error {
	data, err := ReadMessage()
	if err != nil {
		return err
	}
	tx := config.DB.Begin()
	if err := models.InsertArticles(tx, data); err != nil {
		tx.Rollback()
		return err
	}

	ins := struct {
		ID      uint64     `json:"id"`
		Created *time.Time `json:"created"`
	}{
		data.ID,
		data.Created,
	}
	if err := elastic.PutData("http://es:9200/kumparan/_doc", ins); err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	log.Println("data succesfully insert")
	return nil
}

// RunCron -
func RunCron() {
	for {
		if err := InsertData(); err != nil {
			log.Println(err)
		}
	}
}
