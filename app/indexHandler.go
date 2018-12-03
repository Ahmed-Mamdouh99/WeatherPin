package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

type WeatherData struct {
	Name     string
	CityName string
}

// This function should handle the index page
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	// Handle requests from the index page
	if r.Method == "POST" {
		lan, _ := strconv.Atoi(r.FormValue("lan"))
		lon, _ := strconv.Atoi(r.FormValue("lon"))
		fmt.Println("Receive ajax post data string ", lan, " ", lon)

		data, _ := GetJson(lan, lon)

		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(data))
	}
	if r.Method == "GET" {
		//Render index template
		// Define parameters to pass to the index template
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
}
