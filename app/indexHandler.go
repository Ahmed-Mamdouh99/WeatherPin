package main

import (
  "net/http"
  "html/template"
)

// This function should handle the index page
func IndexHandler(w http.ResponseWriter, r *http.Request){
  	// TODO - Handle requests from the index page
  	if r.Method == "POST" {
      //TBD
  	}
  	if r.Method == "GET" {
      //TBD
  	}
    //Render index template
  	// Define parameters to pass to the index template
  	p := struct{
  		App_id string
  		App_code string
  	}{
  		App_id: AppConfig.HereID,
  		App_code: AppConfig.HereAppCode,
  	}
  	// Create template
  	t, err := template.ParseFiles("/static/index.html")
  	if err != nil{
  		Bail(err)
  	}
  	// Execute template
  	err = t.Execute(w, p)
  	if err != nil{
  		Bail(err)
  	}
  }
