package repository

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

var Redis *redisRepository

type redisRepository struct {
	client *redis.Client
	ctx    context.Context
}

func NewRedisRepository(addr, password string, db int) *redisRepository {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password, // no password set
		DB:       db,       // use default DB
	})

	return &redisRepository{
		client: rdb,
		ctx:    context.Background(),
	}
}

func (r *redisRepository) AddToZSet(key string, member string, score float64) *redis.IntCmd {

	return r.client.ZAdd(r.ctx, key, redis.Z{
		Score:  score,
		Member: member,
	})
}

func (r *redisRepository) CountZSet(key string) (int64, error) {
	return r.client.ZCard(r.ctx, key).Result()
}

// its assumption that on this redis client there are only segments as zsets
func (r *redisRepository) RemoveBelowScoreAll(minScore float64) (int64, error) {
	var totalRemoved int64
	var cursor uint64
	var err error

	for {
		// SCAN for all keys
		var keys []string
		keys, cursor, err = r.client.Scan(r.ctx, cursor, "*", 100).Result()
		if err != nil {
			return totalRemoved, err
		}

		// check each key if it’s a ZSET
		for _, key := range keys {
			typ, err := r.client.Type(r.ctx, key).Result()
			if err != nil {
				return totalRemoved, err
			}
			if typ == "zset" {
				removed, err := r.client.ZRemRangeByScore(
					r.ctx,
					key,
					"-inf",
					fmt.Sprintf("(%f", minScore),
				).Result()
				if err != nil {
					return totalRemoved, err
				}
				totalRemoved += removed
			}
		}

		if cursor == 0 {
			break
		}
	}

	return totalRemoved, nil
}
