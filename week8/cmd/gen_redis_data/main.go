package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"strconv"
	"strings"

	redis "github.com/go-redis/redis/v8"
	"github.com/thanhpk/randstr"
)

func main() {
	redisHost := flag.String("h", "localhost:6379", "redis address")
	valueSize := flag.Int("size", 10, "value size in bytes")
	keyCount := flag.Int("count", 100000, "key count")
	flag.Parse()

	rdb := redis.NewClient(&redis.Options{
		Addr: *redisHost,
	})
	ctx := context.Background()
	result, err := checkMemory(ctx, rdb, *keyCount, *valueSize)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Memory usage for %d keys with %s value\n%s\n", *keyCount, prettySize(*valueSize), statistics(result, *keyCount))
}

func statistics(values map[string]int64, count int) string {
	sb := strings.Builder{}
	for k, v := range values {
		sb.WriteString(fmt.Sprintf("\t%s: total %s, each key %s\n", k, prettySize(int(v)), prettySize(int(v)/count)))
	}
	return sb.String()
}

func prettySize(size int) string {
	units := []string{"%.2f B", "%.2f KB", "%.2f MB", "%.2f GB"}
	v := float64(size)
	i := 0
	for ; i < len(units); i++ {
		if v < 1024 {
			break
		}
		v = v / 1024
	}
	return fmt.Sprintf(units[i]+" (%d)", v, size)
}

func checkMemory(ctx context.Context, rdb *redis.Client, keyCount, valueSize int) (map[string]int64, error) {
	resultBeforeInsert, err := checkMemoryUsage(ctx, rdb)
	if err != nil {
		return nil, err
	}
	// fmt.Println("Memory usage before insert: ", resultBeforeInsert)
	for i := 0; i < keyCount; i++ {
		key := fmt.Sprintf("key%06d", i)
		rdb.Set(ctx, key, randstr.Bytes(valueSize), 0)
	}
	resultAfterInsert, err := checkMemoryUsage(ctx, rdb)
	if err != nil {
		return nil, err
	}
	// fmt.Println("Memory usage after insert: ", resultAfterInsert)
	rdb.FlushAll(ctx)
	// resultAfterFlush, err := checkMemoryUsage(ctx, rdb)
	// if err != nil {
	// 	return nil, err
	// }
	// fmt.Println("Memory usage after insert: ", resultAfterFlush)
	return delta(resultBeforeInsert, resultAfterInsert), nil
}

func checkMemoryUsage(ctx context.Context, rdb *redis.Client) (map[string]int64, error) {
	value, err := rdb.Info(ctx, "memory").Result()
	if err != nil {
		return nil, err
	}
	result := make(map[string]int64)
	scanner := bufio.NewScanner(strings.NewReader(value))
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, ":")
		switch split[0] {
		case "used_memory", "used_memory_rss":
			v, err := strconv.ParseInt(split[1], 10, 64)
			if err != nil {
				return nil, err
			}
			result[split[0]] = v
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return result, nil
}

func delta(before, after map[string]int64) map[string]int64 {
	result := make(map[string]int64)
	for k, valueBefore := range before {
		if valueAfter, ok := after[k]; ok {
			result[k] = valueAfter - valueBefore
		}
	}
	return result
}
