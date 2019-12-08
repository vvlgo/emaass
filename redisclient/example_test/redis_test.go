package example_test

import (
	"fmt"
	"github.com/vvlgo/emaass/redisclient"
	"testing"
	"time"
)

func TestRedis(t *testing.T) {
	redisclient.RedisInit("127.0.0.1", "6379", "123456")
	reConn, err := redisclient.GetRedisConn(0)
	if err != nil {
		panic(err)
	}
	//_, err = reConn.Redo("SET", "key", "value", "EX", 60)
	err = reConn.Set("test1", "123")
	err = reConn.ExpiresSet("test2", "123", 10)
	if err != nil {
		panic(err)
	}
	//s, err := reConn.Get("test1")
	//if err != nil {
	//	panic(err)
	//}
	s, err := reConn.Get("test2")
	if err != nil {
		panic(err)
	}
	fmt.Println(s)
	time.Sleep(11 * time.Second)
	s1, err := reConn.Get("test2")
	if err != nil {
		panic(err)
	}
	fmt.Println(s1)
}
