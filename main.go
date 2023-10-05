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
	"github.com/sarmanfatemi/rssagg/internal/handlers"
	"github.com/sarmanfatemi/rssagg/internal/models"

	_ "github.com/lib/pq"
)

// type apiConfig struct {
// 	DB *database.Queries
// }

func main() {
	var err error
	apiCfg := models.ApiConfiguration{}

	useEnvironmentVariables()

	apiCfg.DbQueries, err = useDatabase()
	if err != nil {
		log.Fatal(err)
	}

	// Starting scraper routine
	go startScraping(apiCfg.DbQueries, 10, time.Minute)

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

func useRouter(apiCfg models.ApiConfiguration) (router *chi.Mux) {
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
	v1Router.Get("/healthz", handlers.HandlerReadiness)
	v1Router.Get("/err", handlers.HandlerError)

	v1Router.Post("/users", handlers.Use(handlers.HandlerCreateUser, &apiCfg))
	v1Router.Get("/users", handlers.UseWithGuard(handlers.HandlerGetUserByApiKey, &apiCfg))

	v1Router.Post("/feeds", handlers.UseWithGuard(handlers.HandlerCreateFeed, &apiCfg))
	v1Router.Get("/feeds", handlers.Use(handlers.HandlerGetFeeds, &apiCfg))

	v1Router.Get("/posts", handlers.UseWithGuard(handlers.HandlerGetPostsForUser, &apiCfg))

	v1Router.Post("/feed_follows", handlers.UseWithGuard(handlers.HandlerCreateFeedFollow, &apiCfg))
	v1Router.Get("/feed_follows", handlers.UseWithGuard(handlers.HandlerGetFeedFollows, &apiCfg))
	v1Router.Delete("/feed_follows/{feedFollowID}", handlers.UseWithGuard(handlers.HandlerDeleteFeedFollow, &apiCfg))

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
