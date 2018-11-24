package main

import (
  "github.com/go-redis/redis"
)

var(
  // Redis client
	redis_client = redis.NewClient(&redis.Options{
		Addr:     AppConfig.RedisHost + ":" + AppConfig.RedisPort,
		Password: AppConfig.RedisPasswd,
		DB:       AppConfig.RedisDB,
	})
)

//TBD
