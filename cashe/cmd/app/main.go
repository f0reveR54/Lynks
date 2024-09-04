package main

import (
	"os"
	"stepic-go-basic/cashe/pkg/logger"
	"stepic-go-basic/cashe/pkg/metrics"
	"stepic-go-basic/cashe/pkg/server"
)

func main() {

	logger.Logger = zerolog.New(os.Stdout).With().Timestamp().Logger()

	metrics.New()

	rd := server.InitDB()

	srv := server.New(rd)

	srv.Run()

}
