package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"
	"prod/internal/database"
	"prod/internal/handlers"
	"prod/internal/service"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	errar := godotenv.Load()
	if errar != nil {
		slog.Error("Error", "error:", errar)
	}

	databaseURL := os.Getenv("DB_URL")

	slog.Info("Database URL:", "url", databaseURL)
	slog.Info("Начало запуска")

	db, err := database.Connect(databaseURL)
	if err != nil {
		log.Fatal(err)
	}
	// log.Println(db.DB)

	defer db.Close()

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	if err := db.Ping(); err != nil {
		log.Fatalf("failed to ping db: %v", err)
	}

	store := service.NewRepository(db)
	handler := handlers.NewHandler(store, databaseURL)

	mux := http.NewServeMux()

	mux.HandleFunc("/test", handlers.Test)
	mux.HandleFunc("/main", handler.MainPage)
	mux.HandleFunc("/create", handler.CreateTask)
	mux.HandleFunc("/all", handler.GetAll)
	mux.HandleFunc("/get", handler.GetByID)
	mux.HandleFunc("/update", handler.UpdateTask)
	mux.HandleFunc("/delete", handler.DeleteTask)

	log.Println("Server is running on port 8000")
	err = http.ListenAndServe(":8000", mux)
	if err != nil {
		panic(err)
	}
}
