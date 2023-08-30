package main

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
	"rest-api/controller"
	"rest-api/repository"
	"rest-api/service"
)

type App struct{}

func loadEnvVariables() {
	err := godotenv.Load("app.conf")
	if err != nil {
		log.Fatalf("Error loading app.conf file: %s", err)
	}
}

func (App) Run() {
	loadEnvVariables()
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	conInfo := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", dbUser, dbPassword, dbHost, dbName)

	db, err := sql.Open("postgres", conInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	userRepo := repository.NewUserRepositoryImpl(db)

	userService := service.UserService{UserRepository: userRepo}

	userController := controller.UserController{UserService: userService}

	router := mux.NewRouter()
	router.HandleFunc("/users", userController.GetAllUsersHandler).Methods("GET")
	router.HandleFunc("/users", userController.AddUserHandler).Methods("POST")
	router.HandleFunc("/users/{id:[0-9]+}", userController.GetUserByIdHandler).Methods("GET")
	router.HandleFunc("/users/{id:[0-9]+}", userController.DeleteUserByIdHandler).Methods("DELETE")
	router.HandleFunc("/users/{login}", userController.GetUserByLoginHandler).Methods("GET")

	serverAddr := host + ":" + port
	log.Printf("Server is starting at %s...\n", serverAddr)
	err = http.ListenAndServe(serverAddr, router)
	if err != nil {
		log.Fatalf("Server error: %s", err)
	}
}
