package main

import "net/http"

func handlerReadiness(w http.ResponseWriter,r *http.Request){
	responsewithJSON(w,200,struct {}{})
}