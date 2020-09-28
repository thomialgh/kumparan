package config

import (
	"github.com/jinzhu/gorm"
	"github.com/segmentio/kafka-go"
)

// KafkaReadConfig -
var KafkaReadConfig kafka.ReaderConfig

// DB -
var DB *gorm.DB
