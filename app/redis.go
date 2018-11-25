/*
  This file is responsible for handling redis requests and dealing with the OpenWeatherMap API
  Exported methods:
    -GetJson
    -FlushCache
*/
package main

import (
  "github.com/go-redis/redis"
  "time"
)

// Initializing a new client
const client = redis.NewClient(&redis.Options{
  Addr: AppConfig.RedisHost+":"+AppConfig.RedisPort,
  Password: AppConfig.RedisPasswd,
  DB: AppConfig.RedisDB,
})

// Call this function at the beginning to clear the cache
func FlushCache(){
  client.FlushAll()
}

// This is the function that gets called from the outside
func GetJson(cityCode string) (string, error) {
  // Check if the weather data exists in the cache
  val, err := client.Get(cityCode).Result()
  if err == redis.Nil{ // Data wasn't cached
    // Get the data from the API method
    val, err = getFromOwm(cityCode)
    if err != nil{
      return "", nil
    }
    // Add the data to the cache
    client.Set(cityCode, val, time.Duration(AppConfig.RedisTimeout) * time.Hour)
    // Return result
    return val, err
  }
  // The data was either cached or an error occured
  return val, err
}

func getFromOwm(cityCode string) (string, error) {
  // TODO - Get the json data from OpenWeatherMap using the API key in AppConfig
  return "Success", nil
}
