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

// Config form the config file
var AppConfig = readConfig()

// Reads the file config.yml
func readConfig() Config {
	//Defining output variable
	var config Config
	//Working directory
	/*var pwd, err = os.Getwd()
	if err != nil{
		Bail(err)
	}*/
	url := "/static/conf.yml"
	// Load the file
	yamlFile, err := ioutil.ReadFile(url)
	if err != nil {
		Bail(err)
	}
	// parse yaml file
	yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		Bail(err)
	}
	return config
}

func Bail(err error){
	log.Printf("Error: %v.", err)
	log.Fatal("Bailing! You're on your own now :(")
}

func cleanUp(){
	//VV TODO - Clean up goes here VV
	fmt.Println("Exiting Weather Pin.")
}

func main() {
	// Listening on the web port
	defer http.ListenAndServe(":"+AppConfig.WebPort, nil)
	// Making sure the config file was read
	if AppConfig.WebPort == ""{
		log.Fatal("Weatherpin-Main: Failed to start the app.")
	}
	// Graceful exit
	var gracefulStop = make(chan os.Signal)
	signal.Notify(gracefulStop, syscall.SIGTERM)
	signal.Notify(gracefulStop, syscall.SIGINT)
	go func(){
		sig := <-gracefulStop
		fmt.Printf("\nCaught sig: %+v.\n", sig)
		cleanUp()
		fmt.Println("Goodbye!")
		os.Exit(0)
	}()
	// Handling the calls to the index
	http.HandleFunc("/", IndexHandler)
	//TODO - Add redis handlers
	log.Println("Serving port: ", AppConfig.WebPort)
}
