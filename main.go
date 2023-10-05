package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"github.com/sarmanfatemi/rssagg/internal/database"

	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	var err error
	apiCfg := apiConfig{}

	useEnvironmentVariables()

	apiCfg.DB, err = useDatabase()
	if err != nil {
		log.Fatal(err)
	}

	// Starting scraper routine
	go startScraping(apiCfg.DB, 10, time.Minute)

	router := useRouter(apiCfg)

	err = useHttpServer(router)
	if err != nil {
		log.Fatal(err)
	}
}

func useEnvironmentVariables() {
	// load .env file - the name is optional
	godotenv.Load(".env")
}

func useDatabase() (dbQueries *database.Queries, err error) {
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		return nil, errors.New("DB_URL environment variable is not set")
	}

	connection, err := sql.Open("postgres", dbURL)
	if err != nil {
		return nil, err
	}
	dbQueries = database.New(connection)

	return dbQueries, nil
}

func useRouter(apiCfg apiConfig) (router *chi.Mux) {
	// Configure routing
	router = chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	// Define API version 1
	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/err", handlerError)

	v1Router.Post("/users", apiCfg.handlerCreateUser)
	v1Router.Get("/users", apiCfg.middlewareAuth(apiCfg.handlerGetUserByApiKey))

	v1Router.Post("/feeds", apiCfg.middlewareAuth(apiCfg.handlerCreateFeed))
	v1Router.Get("/feeds", apiCfg.handlerGetFeeds)

	v1Router.Get("/posts", apiCfg.middlewareAuth(apiCfg.handlerGetPostsForUser))

	v1Router.Post("/feed_follows", apiCfg.middlewareAuth(apiCfg.handlerCreateFeedFollow))
	v1Router.Get("/feed_follows", apiCfg.middlewareAuth(apiCfg.handlerGetFeedFollows))
	v1Router.Delete("/feed_follows/{feedFollowID}", apiCfg.middlewareAuth(apiCfg.handlerDeleteFeedFollow))

	router.Mount("/v1", v1Router)

	return router
}

func useHttpServer(router *chi.Mux) (err error) {
	port := os.Getenv("PORT")

	if port == "" {
		return errors.New("PORT is not found in th environment")
	}

	// Creating http server
	server := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	fmt.Printf("Server starting on port %v \n", port)
	serverError := server.ListenAndServe()
	if serverError != nil {
		return serverError
	}

	return nil
}
