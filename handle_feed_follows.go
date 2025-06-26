package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"github.com/SiddhantKandi/HTTPServer/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig)handlerCreateFeedFollow(w http.ResponseWriter,r *http.Request,user database.User){

	type parameter struct {
		Name string `json:"name"`
		FeedID uuid.UUID  `json:"feed_id"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameter{}

	err := decoder.Decode(&params)

	if err!=nil{
		ResponsewithError(w, 400, fmt.Sprintf("Error parsing JSON : %v",err))
		return
	}

	feedFollow,err := apiCfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID: user.ID,
		FeedID: params.FeedID,
	})

	if err!=nil{
		ResponsewithError(w, 400, fmt.Sprintf("Couldn't create feed follow: %v",err))
		return
	}
	

	responsewithJSON(w,201,databaseFeedFollowtoFeedFollow(feedFollow))
}
