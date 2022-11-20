package test

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"testing"
)

func TestConn(t *testing.T) {
	ctx := context.Background()
	rdb := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{":7001", ":7002", ":7003", ":7004", ":7005", ":7006"},
	})

	result, err := rdb.ClusterNodes(ctx).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(result)

	// ForEachShard ForEachMaster ForEachSlave
	err = rdb.ForEachShard(ctx, func(ctx context.Context, client *redis.Client) error {
		s, err := client.Ping(ctx).Result()
		fmt.Println(s)
		return err
	})
	if err != nil {
		panic(err)
	}
}
