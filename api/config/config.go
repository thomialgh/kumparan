package config

import (
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/jinzhu/gorm"
	"github.com/segmentio/kafka-go"
)

// DB - global variable for db connection
var DB *gorm.DB

// ES - global variable for elasticsearch connection
var ES *elasticsearch.Client

// KafkaWriterConf - global varibale for kafka configuration
var KafkaWriterConf kafka.WriterConfig
