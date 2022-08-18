package redis_glance

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v9"
	uuid "github.com/satori/go.uuid"
	"sync"
	"testing"
)

func clusterSetAndDelete(ctx context.Context, rdb *redis.ClusterClient) {
	key := uuid.NewV4().String()
	rdb.Do(ctx, "SET", key, "cluster_test_"+key)
	_, err := rdb.Do(ctx, "get", key).Result()
	if err != nil {
		if err == redis.Nil {
			fmt.Println("key does not exists")
			return
		}
		panic(err)
	}
	rdb.Do(ctx, "DEL", key)
}

func clusterSetAndDeleteRepeat(wg *sync.WaitGroup, ctx context.Context, rdb *redis.ClusterClient, repeat int) {
	defer wg.Done()

	for i := 1; i <= repeat; i++ {
		clusterSetAndDelete(ctx, rdb)
	}
}

func BenchmarkCluster(b *testing.B) {
	ctx := context.Background()

	rdb := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{":7000", ":7001", ":7002", ":7003", ":7004", ":7005"},

		// To route commands by latency or randomly, enable one of the following.
		//RouteByLatency: true,
		//RouteRandomly: true,
	})

	defer func(rdb *redis.ClusterClient) {
		err := rdb.Close()
		if err != nil {
			panic(err)
		}
	}(rdb)

	var wg sync.WaitGroup

	goroutines := 100
	for i := 0; i < goroutines; i++ {
		wg.Add(1)
		go clusterSetAndDeleteRepeat(&wg, ctx, rdb, 1000)
	}

	wg.Wait()
}
