package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	// WeatherPin
	WebHost 			string
	WebPort 			string
	// Redis
	RedisPort   	string
	RedisPasswd 	string
	RedisDB     	int
	RedisHost   	string
	RedisTimeout	int
	// Here maps
	HereID      	string
	HereAppCode 	string
	// OWM OpenWeatherMap
	OWMKey 				string
}

// Config form the config file
const AppConfig = readConfig()

// Reads the file config.yml
func readConfig() Config {
	// Defining output variable
	var config Config
	// Load the file
	url := "/static/conf.yml"
	yamlFile, err := ioutil.ReadFile(url)
	if err != nil {
		Bail(err)
	}
	// parse yaml file
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		Bail(err)
	}
	return config
}

func Bail(err error){
	log.Printf("Error: %v.", err)
	log.Println("Bailing! You're on your own now :(")
	cleanUp()
}

func cleanUp(){
	// VV TODO - Clean up goes here VV
	// Flushing cache before exiting
	FlushCache()

	// Pretty messages
	fmt.Println("Exiting Weather Pin.")
	fmt.Println("Goodbye!")
	// System exit
	os.Exit(0)
}

func main() {
	//Flsuh Redis cache
	FlushCache()
	// Listening on the web port
	defer http.ListenAndServe(":"+AppConfig.WebPort, nil)
	// Making sure the config file was read
	if AppConfig.WebPort == ""{
		log.Fatal("Weatherpin-Main: Invalid config file. Failed to start the app.")
	}
	// Graceful exit
	var gracefulStop = make(chan os.Signal)
	signal.Notify(gracefulStop, syscall.SIGTERM)
	signal.Notify(gracefulStop, syscall.SIGINT)
	go func(){
		sig := <-gracefulStop
		log.Printf("\nCaught sig: %+v.\n", sig)
		cleanUp()
	}()
	// Handling the calls to the index
	http.HandleFunc("/", IndexHandler)
	//TODO - Add redis handlers
	log.Println("Serving port: ", AppConfig.WebPort)
}
