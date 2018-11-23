package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Config struct {
	// WeatherPin
	WebHost string
	WebPort string
	// Redis
	RedisPort   string
	RedisPasswd string
	RedisDB     int
	RedisHost   string
	//Here maps
	HereID      string
	HereAppCode string
	//OWM OpenWeatherMap
	OWMKey string
}

var (
	//Config form the config file
	appConfig = readConfig()
	//readConfig()
	//Redis client
	redis_client = redis.NewClient(&redis.Options{
		Addr:     appConfig.RedisHost + ":" + appConfig.RedisPort,
		Password: appConfig.RedisPasswd,
		DB:       appConfig.RedisDB,
	})
)

//Reads the file config.yml
func readConfig() Config {
	url := "conf.yml"
	pwd, _ := os.Getwd()
	yamlFile, err := ioutil.ReadFile(pwd + url)
	var config Config

	if err != nil {
		log.Fatal("Could not read config file: ", err)
		log.Println("Falling back to default config")
		getDefaultConfig(&config)
		return config
	}

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
		log.Println("Falling back to default config.")
		getDefaultConfig(&config)
	}
	return config
}

//Sets the config to the default
func getDefaultConfig(c *Config) {
	*c = Config{
		"", "8080", "6379", "", 0,
		"redis-service", "ivpuD5BHEfdq49qdXWIS",
		"jrVubtlb8B0giad6hzo4YA", "8998549a7a2993add254b22a8c84cc8c"}
}

//This function should handle the index page
func indexHandler(w http.ResponseWriter, r *http.Request) {
	//TODO - Handle requests from the index page.
	if r.Method == "POST" {
	}
	if r.Method == "GET" {
	}
	fmt.Fprintf(w, "Hi")
}

func redis_handler(w http.ResponseWriter, r *http.Request) {
	//TODO - Handle requests from the front end... Maybe?
}

func main() {
	defer http.ListenAndServe(":"+appConfig.WebPort, nil)
	http.HandleFunc("/", indexHandler)
	log.Println("Serving port: ", appConfig.WebPort)
}
