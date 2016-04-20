package redis_test

import (
  "testing"
  "auth/core/redis"
)

func TestConnection(t *testing.T) {
  redisConn := redis.Connect();
  if redisConn == nil {
    t.Error("redis not initialised")
  }
}
