package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/yueyue000/webook/internal/domain"
	"time"
)

var ErrKeyNotExist = redis.Nil // key不存在，用特定的错误

type UserCache struct {
	client     redis.Cmdable // 面向接口编程。 接口不能用指针，用指针调不了里面的方法。
	expiration time.Duration
}

// NewUserCache 依赖注入，自己不做初始化，要从外面传进来
func NewUserCache(client redis.Cmdable, expiration time.Duration) *UserCache {
	return &UserCache{
		client:     client,
		expiration: expiration,
	}
}

// 只要error为nil，就认为缓存一定有数据。如果没有数据，返回特定的error。
func (cache *UserCache) Get(ctx context.Context, id int64) (domain.User, error) {
	key := cache.Key(id)
	val, err := cache.client.Get(ctx, key).Bytes()
	if err != nil {
		return domain.User{}, err
	}
	var u domain.User
	err = json.Unmarshal(val, &u)
	return u, err
}

func (cache *UserCache) Set(ctx context.Context, u domain.User) error {
	val, err := json.Marshal(u)
	if err != nil {
		return err
	}
	key := cache.Key(u.ID)
	return cache.client.Set(ctx, key, val, cache.expiration).Err()
}

func (cache *UserCache) Key(id int64) string {
	return fmt.Sprintf("user:info:%d", id)
}
