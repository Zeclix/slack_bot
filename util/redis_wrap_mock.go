package util

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	redis "gopkg.in/redis.v3"
)

type StringCmdMock struct {
	*redis.StringCmd
	value []byte
	e     error
}

type RedisClientMock struct {
	data map[string][]byte
}

func (cmd *StringCmdMock) Val() string {
	return string(cmd.value)
}

func (cmd *StringCmdMock) Result() (string, error) {
	return cmd.Val(), cmd.e
}

func (cmd *StringCmdMock) Bytes() ([]byte, error) {
	return cmd.value, cmd.e
}

func (cmd *StringCmdMock) Int64() (int64, error) {
	if cmd.e != nil {
		return 0, cmd.e
	}
	return strconv.ParseInt(cmd.Val(), 10, 64)
}

func (cmd *StringCmdMock) Uint64() (uint64, error) {
	if cmd.e != nil {
		return 0, cmd.e
	}
	return strconv.ParseUint(cmd.Val(), 10, 64)
}

func (cmd *StringCmdMock) Float64() (float64, error) {
	if cmd.e != nil {
		return 0, cmd.e
	}
	return strconv.ParseFloat(cmd.Val(), 64)
}

func (cmd *StringCmdMock) Scan(val interface{}) error {
	return nil
}

func (cmd *StringCmdMock) String() string {
	return fmt.Sprint(cmd, cmd.value)
}

func (c *RedisClientMock) Set(key string, value interface{}, expiration time.Duration) StatusCmd {
	ret := new(redis.StatusCmd)
	switch value.(type) {
	case string:
		c.data[key] = []byte(value.(string))
	case []byte:
		c.data[key] = value.([]byte)
	default:
		c.data[key] = []byte(fmt.Sprintf("%v", value))
	}
	return ret
}

func (c *RedisClientMock) Get(key string) StringCmd {
	ret := &StringCmdMock{new(redis.StringCmd), []byte{}, nil}
	if data, ok := c.data[key]; ok {
		ret.value = data
	} else {
		ret.e = errors.New("Value not found")
	}
	return ret
}

func NewRedisMock() RedisClient {
	return &RedisClientMock{map[string][]byte{}}
}
