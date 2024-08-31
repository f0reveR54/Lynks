package redis

import (
	"context"
	"fmt"
	"stepic-go-basic/cashe/pkg/db"

	"github.com/go-redis/redis/v8"
)

type DB struct {
	rd *redis.Client
}

func (d *DB) SaveURL(ctx context.Context, url db.URL) error {

	err := d.rd.Set(ctx, url.Short, url.Orig, 0).Err()
	if err != nil {
		return err
	}

	return nil

}

func (d *DB) GetOriginal(ctx context.Context, s string) (string, error) {

	cmd, err := d.rd.Get(ctx, "http://localhost:8081/"+s).Result()
	if err != nil {
		return "", err
	}

	return cmd, nil

}

func NewDB(ctx context.Context) (*DB, error) {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // без пароля
		DB:       0,  // БД по умолчанию
	})

	pong, err := redisClient.Ping(ctx).Result()
	fmt.Println(pong, err)
	if err != nil {
		fmt.Println("Redis is not Connect")
	} else {
		fmt.Println("Redis is Connected")
	}

	return &DB{rd: redisClient}, nil
}
