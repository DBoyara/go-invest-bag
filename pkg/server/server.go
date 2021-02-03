package server

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/DBoyara/go-invest-bag/pkg/models"
	"github.com/DBoyara/go-invest-bag/pkg/repository"
	"github.com/go-chi/chi"
	"gorm.io/gorm"
)

type server struct {
	ctx context.Context
	mux chi.Router
	db  *gorm.DB
}

func NewServer(ctx context.Context, mux chi.Router, db *gorm.DB) *server {
	return &server{ctx: ctx, mux: mux, db: db}
}

func (s *server) Init() {
	s.mux.Get("/api/position", s.getPosition)
	s.mux.Post("/api/position", s.addPosition)
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}

func (s *server) getPosition(w http.ResponseWriter, r *http.Request) {

	repository.GetPosition(s.db, s.ctx)

	log.Println()
}

func (s *server) addPosition(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		errModel := models.ErrModel{Err: err.Error()}
		w.WriteHeader(http.StatusServiceUnavailable)
		jsonResponse(w, r, errModel)
		return
	}

	position := &models.Position{}
	err = json.Unmarshal(body, position)
	if err != nil {
		errModel := models.ErrModel{Err: err.Error()}
		w.WriteHeader(http.StatusNotImplemented)
		jsonResponse(w, r, errModel)
		return
	}
	log.Printf("position %v", position)

	position, status := repository.AddPosition(s.db, s.ctx, position)

	w.WriteHeader(status)
	jsonResponse(w, r)
}

func jsonResponse(w http.ResponseWriter, r *http.Request, dtos ...interface{}) {
	respBody, err := json.Marshal(dtos)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadGateway)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	_, err = w.Write(respBody)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadGateway)
		return
	}
}
