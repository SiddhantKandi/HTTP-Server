package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func ResponsewithError(w http.ResponseWriter,code int, msg string){
	if code > 499 {
		log.Println("Responding with 500 error : ",msg)
	}

	type ErrorResponse struct {
		Error string `json:"error"`
	}

	responsewithJSON(w, code, ErrorResponse{
		Error:msg,
	})
}


func responsewithJSON(w http.ResponseWriter,code int,payload interface{}){
	data,err := json.Marshal(payload)

	if err!=nil{
		log.Printf("failed to marshal JSON response %v",payload)
		w.WriteHeader(500) //header code 500
		return
	}

	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}