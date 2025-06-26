package main

import (
	"fmt"
	"net/http"
	"github.com/SiddhantKandi/HTTPServer/internal/auth"
	"github.com/SiddhantKandi/HTTPServer/internal/database"
)

type authHandler func(http.ResponseWriter,*http.Request,database.User)


func (apiCfg *apiConfig) middlewareAuth(handler authHandler) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		APIKey,err := auth.GetAPIKey(r.Header)

		if err != nil {
			ResponsewithError(w, 403, fmt.Sprintf("Couldn't get APIKey: %v",err))
			return
		}

		user,err := apiCfg.DB.GetUserByAPIKey(r.Context(),APIKey)

		if err != nil {
			ResponsewithError(w, 400, fmt.Sprintf("Couldn't get user: %v",err))
			return
		}

		handler(w,r,user)
	}	
}