package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/SiddhantKandi/HTTPServer/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig)handlerUser(w http.ResponseWriter,r *http.Request){

	type parameter struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameter{}

	err := decoder.Decode(&params)

	if err!=nil{
		ResponsewithError(w, 400, fmt.Sprintf("Error parsing JSON : %v",err))
		return
	}

	user,err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name: params.Name,
	})

	if err!=nil{
		ResponsewithError(w, 400, fmt.Sprintf("Couldn't create User: %s",err))
		return
	}
	

	responsewithJSON(w,200,databaseUsertoUser(user))
}