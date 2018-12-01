/*
  This file is responsible for handling redis requests and dealing with the OpenWeatherMap API
  Exported methods:
    -GetJson
    -FlushCache
*/
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/go-redis/redis"
)

// Initializing a new client
var client = redis.NewClient(&redis.Options{
	Addr:     AppConfig.RedisHost + ":" + AppConfig.RedisPort,
	Password: AppConfig.RedisPasswd,
	DB:       AppConfig.RedisDB,
})

// Call this function at the beginning to clear the cache
func FlushCache() {
	client.FlushAll()
}

// This is the function that gets called from the outside
func GetJson(lan int, lon int) (string, error) {
	// Check if the weather data exists in the cache
	key := fmt.Sprintf("%d %d", lan, lon)
	val, err := client.Get(key).Result()
	if err == redis.Nil { // Data wasn't cached
		// Get the data from the API method
		val, err = getFromOwm(lan, lon)
		if err != nil {
			return "", nil
		}
		// Add the data to the cache
		client.Set(key, val, time.Duration(AppConfig.RedisTimeout)*time.Hour)
		// Return result
		//fmt.println("hello from cache")
		return val, err
	}
	// The data was either cached or an error occured
	return val, err
}

func getFromOwm(lan int, lon int) (string, error) {
	// TODO - Get the json data from OpenWeatherMap using the API key in AppConfig
	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?lat=%d&lon=%d&APPID=%s&units=metric", lan, lon, AppConfig.OWMKey)
	res, err := http.Get(url)
	if err != nil {
		Bail(err)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		Bail(readErr)
	}

	return string(body), nil
}
