package cache

import (
	"time"

	"github.com/go-redis/redis"
)

type Movie struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type MovieService interface {
	GetMovie(id string) (*Movie, error)
	GetMovies([]*Movie, error)
	CreateMovie(movie *Movie) (*Movie, error)
	UpdateMovie(movie *Movie) (*Movie, error)
	DeleteMovie(id string) error
}

type redisCache struct {
	host string
	db   int
	exp  time.Duration
}

func NewRedisCache(host string, db int, exp time.Duration) *redisCache {
	return &redisCache{
		host: host,
		db:   db,
		exp:  exp,
	}
}

func (cache redisCache) getClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     cache.host,
		Password: "",
		DB:       cache.db,
	})
}

