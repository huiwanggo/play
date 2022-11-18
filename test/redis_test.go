package test

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"testing"
	"time"
)

func TestConn(t *testing.T) {
	rdb := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{":7001", ":7002", ":7003", ":7004", ":7005", ":7006"},
	})

	rdb.Set(context.Background(), "test", "test", time.Hour)
	fmt.Println(rdb.Get(context.Background(), "test").Result())
}
