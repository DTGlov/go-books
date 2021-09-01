package main

import (
	"encoding/json"
	"net/http"
)

func (app *application) statusHandler(w http.ResponseWriter,r *http.Request){
	currentStatus := AppStatus{
	Status: "Available",
	Environment: app.config.env,
	Version: version,
}
js,err :=json.MarshalIndent(currentStatus,"","\t")
if err!=nil{
	app.errorLog.Fatal(err)
}
//the js into the browser
w.Header().Set("Content-Type","application/json")
w.WriteHeader(http.StatusOK)
w.Write(js)
}