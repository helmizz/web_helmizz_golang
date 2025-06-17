package middlewares

import (
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/boj/redistore"
)

var RedisClient *redis.Client

func InitRedisClient() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   0,
	})
}

func NewRedisSessionStore() *redistore.RediStore {
	store, err := redistore.NewRediStore(10, "tcp", "localhost:6379", "", []byte("secret-key"))
	if err != nil {
		panic(err)
	}
	store.SetMaxAge(int(30 * time.Minute.Seconds()))
	return store
}
