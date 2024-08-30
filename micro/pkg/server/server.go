package server

import (
	"context"
	"log"
	"net/http"

	"stepic-go-basic/micro/pkg/api"
	"stepic-go-basic/micro/pkg/db"
	"stepic-go-basic/micro/pkg/db/pgsql"
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
	http.ListenAndServe(":8080", s.api.Router())
}

func InitDB() *pgsql.DB {
	ctx := context.Background()

	connString := "postgres://admin:password@localhost:5432/test"

	db, err := pgsql.NewDB(ctx, connString)
	if err != nil {
		log.Fatalf("Error creating DB instance: %v\n", err)
	}

	return db
}
