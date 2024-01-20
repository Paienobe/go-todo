package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Paienobe/go-todo/internal/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	godotenv.Load()

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT does not exist in environment")
	}

	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		log.Fatal("DB_URL does not exist in environment")
	}
	dbConn, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal("Failed to connect to DB", err)
	}

	db := database.New(dbConn)
	apiCfg := apiConfig{
		DB: db,
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "DELETE", "PUT"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	router.Get("/", homeHandler)

	router.Post("/register", apiCfg.createUser)
	router.Post("/login", apiCfg.login)

	router.Post("/create", apiCfg.isAuthorized(apiCfg.createTask))
	router.Put("/update", apiCfg.isAuthorized(apiCfg.updateTask))
	router.Delete("/delete/{task_id}", apiCfg.isAuthorized(apiCfg.deleteTask))

	router.Get("/test", apiCfg.isAuthorized(apiCfg.getUser))

	server := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}
	fmt.Println("Server starting on port ", portString)
	log.Fatal(server.ListenAndServe())
}
