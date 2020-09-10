package main

import (
	// "fmt"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)
type HealthResponse struct {
	Status string `json:"status"`
	Code int `json:"code"`
}

func healthCheck (w http.ResponseWriter, r *http.Request){
	healthResp := HealthResponse{"ok", 200} 
	response, err := json.Marshal(healthResp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	  }
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
	
}

func main()  {
	route := mux.NewRouter()
	route.HandleFunc("/api", healthCheck)
	http.Handle("/", route)
	http.ListenAndServe(":8080", nil)
}