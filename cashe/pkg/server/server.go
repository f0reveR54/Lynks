package server

import (
	"context"
	"log"
	"net/http"

	"stepic-go-basic/cashe/pkg/api"
	"stepic-go-basic/cashe/pkg/db"
	"stepic-go-basic/cashe/pkg/db/redis"
)

type Server struct {
	api *api.API
}

func New(db db.Interface) *Server {

	srv := Server{}

	srv.api = api.New(db)

	return &srv
}

func (s *Server) Run() {
	http.ListenAndServe(":8081", s.api.Router())
}

func InitDB() *redis.DB {

	ctx := context.Background()

	db, err := redis.NewDB(ctx)
	if err != nil {
		log.Fatalf("Error creating DB instance: %v\n", err)
	}

	return db
}
