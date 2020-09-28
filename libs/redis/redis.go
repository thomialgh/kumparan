package redis

import (
	"time"

	"github.com/gomodule/redigo/redis"
)

var pool *redis.Pool

type RedisConn struct {
	Conn redis.Conn
}

func createPool(address string) {
	pool = redis.NewPool(func() (redis.Conn, error) {
		return redis.Dial("tcp", address)
	}, 10)
}

// InitRedis -
func InitRedis(address string) error {
	createPool(address)
	conn := GetConn()
	defer conn.Conn.Close()
	_, err := conn.Conn.Do("PING")
	return err
}

// GetConn -
func GetConn() RedisConn {
	return RedisConn{
		Conn: pool.Get(),
	}
}

// Set -
func (r RedisConn) Set(key, value string, duration time.Duration) error {
	_, err := r.Conn.Do("SET", key, value, "EX", duration.Seconds())
	return err
}

// Get -
func (r RedisConn) Get(key string) (string, error) {
	rep, err := r.Conn.Do("GET", key)
	if rep == nil {
		return "", nil
	}
	return string(rep.([]byte)), err
}
