package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
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
	// Here maps
	HereID      string
	HereAppCode string
	// OWM OpenWeatherMap
	OWMKey string
}

var (
	// Config form the config file
	appConfig = readConfig()
	// Redis client
	redis_client = redis.NewClient(&redis.Options{
		Addr:     appConfig.RedisHost + ":" + appConfig.RedisPort,
		Password: appConfig.RedisPasswd,
		DB:       appConfig.RedisDB,
	})
)

// Reads the file config.yml
func readConfig() Config {
	var config Config
	url := "conf.yml"
	pwd, _ := os.Getwd()
	// Load the file
	yamlFile, err := ioutil.ReadFile(pwd + url)
	if err != nil {
		bail(err)
	}
	// parse yaml file
	yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		bail(err)
	}
	return config
}

// This function should handle the index page
func indexHandler(w http.ResponseWriter, r *http.Request) {
	// TODO - Handle requests from the index page.
	if r.Method == "POST" {
	}
	if r.Method == "GET" {
	}
	fmt.Fprintf(w, "Hi")
}

func redis_handler(w http.ResponseWriter, r *http.Request) {
	// TODO - Handle requests from the front end... Maybe?
}

func bail(err error){
	log.Printf("Error: %v.", err)
	log.Fatal("Bailing out. You're on your own now :(")
}

func gracefulStop(){
	//VV TODO - Clean up goes here VV
}

func main() {
	// Listening on the web port
	defer http.ListenAndServe(":"+appConfig.WebPort, nil)
	// Making sure the config file was read
	if appConfig.WebPort == ""{
		log.Fatal("Weatherpin-Main: Failed to start the app.")
	}
	// Graceful exit
	var gracefulStop = make(chan os.Signal)
	signal.Notify(gracefulStop, syscall.SIGTERM)
	signal.Notify(gracefulStop, syscall.SIGINT)
	go func(){
		sig := <-gracefulStop
		fmt.Printf("Caught sig: %+v.", sig)
		gracefulStop()
		fmt.Println("Good night!")
		os.Exit(0)
	}()
	// Handling the calls to the index
	http.HandleFunc("/", indexHandler)
	log.Println("Serving port: ", appConfig.WebPort)
}
