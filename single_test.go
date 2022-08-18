package redis_glance

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v9"
	uuid "github.com/satori/go.uuid"
	"sync"
	"testing"
)

func setAndDelete(ctx context.Context, rdb *redis.Client) {
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

func setAndDeleteRepeat(wg *sync.WaitGroup, ctx context.Context, rdb *redis.Client, repeat int) {
	defer wg.Done()

	for i := 1; i <= repeat; i++ {
		setAndDelete(ctx, rdb)
	}
}

func BenchmarkSingle(b *testing.B) {
	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	defer func(rdb *redis.Client) {
		err := rdb.Close()
		if err != nil {
			panic(err)
		}
	}(rdb)

	var wg sync.WaitGroup

	goroutines := 100
	for i := 0; i < goroutines; i++ {
		wg.Add(1)
		go setAndDeleteRepeat(&wg, ctx, rdb, 1000)
	}

	wg.Wait()
}
