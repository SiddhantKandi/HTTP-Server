package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/SiddhantKandi/HTTPServer/internal/database"
	"github.com/go-chi/chi"
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

func (apiCfg *apiConfig)handlerGetAllFeedFollow(w http.ResponseWriter,r *http.Request,user database.User){

	feedFollows,err := apiCfg.DB.GetAllFeedFollows(r.Context(),user.ID)

	if err!=nil{
		ResponsewithError(w, 400, fmt.Sprintf("Couldn't get all feed follow: %v",err))
		return
	}
	

	responsewithJSON(w,201,databaseGetAllFeedFollowstoFeedFollows(feedFollows))
}


func (apiCfg *apiConfig)handlerDeleteFeedFollow(w http.ResponseWriter,r *http.Request,user database.User){

	feedFollowSTR := chi.URLParam(r, "feed_follow_id")

	feedFollowId,err := uuid.Parse(feedFollowSTR)

	if err!=nil{
		ResponsewithError(w, 400, fmt.Sprintf("Couldn't parse feed follow ID: %v",err))
		return
	}

	err = apiCfg.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		ID:feedFollowId,
		UserID: user.ID,
	})

	if err!=nil{
		ResponsewithError(w, 400, fmt.Sprintf("Couldn't delete feed follow : %v",err))
		return
	}

	responsewithJSON(w,201,struct{}{})

}