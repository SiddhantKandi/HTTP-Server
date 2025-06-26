package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"github.com/SiddhantKandi/HTTPServer/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig)handlerCreateFeed(w http.ResponseWriter,r *http.Request,user database.User){

	type parameter struct {
		Name string `json:"name"`
		Url string  `json:"url"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameter{}

	err := decoder.Decode(&params)

	if err!=nil{
		ResponsewithError(w, 400, fmt.Sprintf("Error parsing JSON : %v",err))
		return
	}

	feed,err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name: params.Name,
		Url: params.Url,
		UserID: user.ID,
	})

	if err!=nil{
		ResponsewithError(w, 400, fmt.Sprintf("Couldn't create feed: %s",err))
		return
	}
	

	responsewithJSON(w,201,databaseFeedtoFeed(feed))
}


func (apiCfg *apiConfig)handleGetAllFeeds(w http.ResponseWriter,r *http.Request){	
	feeds,err := apiCfg.DB.GetAllFeeds(r.Context())

	if err!=nil{
		ResponsewithError(w, 400, fmt.Sprintf("Couldn't get feeds: %v",err))
		return
	}
	

	responsewithJSON(w,201,databaseFeedstoFeeds(feeds))
}