package util

import (
	"gopkg.in/redis.v3"
	"time"
)

type RedisClient interface {
	Get(key string) StringCmd
	Set(key string, value interface{}, expiration time.Duration) StatusCmd
}

type Command interface {
	Val() string
	Result() (string, error)
	String() string
}

type StringCmd interface {
	Command
	Bytes() ([]byte, error)
	Float64() (float64, error)
	Int64() (int64, error)
	Scan(val interface{}) error
	Uint64() (uint64, error)
}

type StatusCmd interface {
	Command
}

type RedisClientWrap struct {
	R *redis.Client
}

func (r *RedisClientWrap) Get(key string) StringCmd {
	return r.R.Get(key)
}

func (r *RedisClientWrap) Set(key string, value interface{}, expiration time.Duration) StatusCmd {
	return r.R.Set(key, value, expiration)
}
