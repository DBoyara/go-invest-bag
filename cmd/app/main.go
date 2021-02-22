package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/DBoyara/go-invest-bag/pkg/models"
	"github.com/DBoyara/go-invest-bag/pkg/server"
)

const (
	defaultPort = "8080"
	defaultHost = "0.0.0.0"
)

func main() {
	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = defaultPort
	}

	host, ok := os.LookupEnv("HOST")
	if !ok {
		host = defaultHost
	}

	if err := execute(net.JoinHostPort(host, port)); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func execute(addr string) (err error) {
	ctx := context.Background()
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Миграция схем
	// db.AutoMigrate(&models.Position{})
	db.Migrator().CreateTable(&models.Position{})

	mux := chi.NewMux()
	application := server.NewServer(ctx, mux, db)
	application.Init()

	s := &http.Server{
		Addr:    addr,
		Handler: application,
	}
	log.Printf("Server run on http://%s", addr)
	return s.ListenAndServe()
}
