package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type WeatherData struct {
	Name     string
	CityName string
}

// This function should handle the index page
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	// TODO - Handle requests from the index page
	if r.Method == "POST" {
		//TBD
		lan, _ := strconv.Atoi(r.FormValue("lan"))
		lon, _ := strconv.Atoi(r.FormValue("lon"))
		fmt.Println("Receive ajax post data string ", lan, " ", lon)
		url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?lat=%d&lon=%d&APPID=637aa79319844d254e33bad285e23e82", lan, lon)

		res, err := http.Get(url)
		if err != nil {
			Bail(err)
		}

		body, readErr := ioutil.ReadAll(res.Body)
		if readErr != nil {
			log.Fatal(readErr)
		}
		fmt.Println(body)

		w.Header().Set("Content-Type", res.Header.Get("Content-Type"))
		w.Write(body)
		res.Body.Close()
	}
	if r.Method == "GET" {
		//TBD
		p := struct {
			App_id   string
			App_code string
		}{
			App_id:   AppConfig.HereID,
			App_code: AppConfig.HereAppCode,
		}
		// Create template
		t, err := template.ParseFiles("/static/index.html")
		if err != nil {
			Bail(err)
		}
		// Execute template
		err = t.Execute(w, p)
		if err != nil {
			Bail(err)
		}
	}
	//Render index template
	// Define parameters to pass to the index template
}
