package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/SiddhantKandi/HTTPServer/internal/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
    	
)

type apiConfig struct{
	DB * database.Queries
}



func main(){

	godotenv.Load()

	portString := os.Getenv("PORT")

	if portString == ""{
		fmt.Println("Port not found")
		return
	}

	DBString := os.Getenv("DB_URL")

	if DBString == ""{
		fmt.Println("dbURL not found")
		return
	}

	conn,err := sql.Open("postgres", DBString)

	if err != nil{
		log.Fatalf("Database connection failed : %v",err)
	}



	apiCfg := apiConfig{
		DB: database.New(conn),
	}


	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()

	v1Router.Get("/healthz", handlerReadiness)

	v1Router.Get("/err", handlerError)

	v1Router.Post("/users",apiCfg.handlerUser)

	v1Router.Get("/getUser", apiCfg.middlewareAuth(apiCfg.handleGetUser))

	v1Router.Post("/feeds", apiCfg.middlewareAuth(apiCfg.handlerCreateFeed))

	v1Router.Get("/getAllFeeds", apiCfg.handleGetAllFeeds)

	v1Router.Post("/feedFollows", apiCfg.middlewareAuth(apiCfg.handlerCreateFeedFollow))

	v1Router.Get("/getAllFeedFollows", apiCfg.middlewareAuth(apiCfg.handlerGetAllFeedFollow))

	router.Mount("/v1", v1Router)


	server := &http.Server{
		Handler:router,
		Addr : ":" + portString,
	}

	log.Printf("Server running on port :%s",portString)
	err = server.ListenAndServe()

	if err!=nil {
		log.Fatal(err)
	}
}