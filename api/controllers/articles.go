package controllers

import (
	"encoding/json"
	"fmt"
	"kumparan/api/config"
	"kumparan/libs/elastic"
	"kumparan/libs/kafka"
	"kumparan/libs/redis"
	"kumparan/models"
	"log"
	"strconv"
	"sync"
	"time"

	"github.com/labstack/echo/v4"
)

// CreateArticle -
func CreateArticle(c echo.Context) error {
	var data models.Articles
	if err := c.Bind(&data); err != nil {
		return config.BadRequestResp(c, "Invalid data request")
	}

	if data.Author == "" {
		return config.BadRequestResp(c, "Author cannot be empty")
	}
	if data.Body == "" {
		return config.BadRequestResp(c, "body cannot be empty")
	}

	if err := kafka.Publish(config.KafkaWriterConf, []byte("kumparanpost"), data); err != nil {
		if err != nil {
			return config.InternalErrResp(c, err)
		}
	}

	return config.OKResponse(c)
}

type elasticData struct {
	Hits struct {
		Hits []struct {
			Source struct {
				ID      uint64 `json:"id"`
				Created string `json:"created"`
			} `json:"_source"`
		} `json:"hits"`
	} `json:"hits"`
}

// GetArticles -
func GetArticles(c echo.Context) error {
	var page uint64 = 1
	var err error
	if pg := c.QueryParam("page"); pg != "" {
		page, err = strconv.ParseUint(pg, 10, 64)
		if err != nil {
			log.Println(err)
			return config.BadRequestResp(c, "Invalid page value")
		}
	}

	url := "http://es:9200/kumparan/_search"
	query := elastic.Query{
		From: (page - 1) * 10,
		Size: 10,
		Sort: []map[string]interface{}{
			{"created": "desc"},
		},
	}

	data, err := elastic.GetData(url, query)
	if err != nil {
		return config.InternalErrResp(c, err)
	}
	var res elasticData
	if err := json.Unmarshal(data, &res); err != nil {
		return config.InternalErrResp(c, err)
	}
	var wg sync.WaitGroup
	var m sync.RWMutex
	var result []models.Articles
	fmt.Println(res)
	for _, v := range res.Hits.Hits {
		wg.Add(1)
		go getData(&wg, &m, &result, v.Source.ID)
	}
	wg.Wait()
	defer func() {
		data, _ := json.Marshal(result)
		conn := redis.GetConn()
		conn.Set(c.Request().RequestURI, string(data), 5*time.Minute)
	}()
	return config.Data(c, result)
}

func getData(wg *sync.WaitGroup, m *sync.RWMutex, datas *[]models.Articles, ID uint64) {
	defer wg.Done()
	var data models.Articles
	err := models.GetDataArticles(config.DB, &data, ID)
	if err != nil {
		return
	}
	m.Lock()
	*datas = append(*datas, data)
	m.Unlock()
}

// CacheMiddlerwere -
func CacheMiddlerwere(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		conn := redis.GetConn()
		data, err := conn.Get(c.Request().RequestURI)
		if err != nil {
			log.Println(err)
			return next(c)
		}
		if data != "" {
			var d []models.Articles
			err = json.Unmarshal([]byte(data), &d)
			if err == nil {
				log.Println(err)
				return config.Data(c, d)
			}
		}
		return next(c)
	}
}
