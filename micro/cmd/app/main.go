package main

import (
	"os"
	"stepic-go-basic/micro/pkg/logger"
	"stepic-go-basic/micro/pkg/metrics"
	"stepic-go-basic/micro/pkg/server"

	"github.com/rs/zerolog"

	"github.com/prometheus/client_golang/prometheus"
)

func main() {

	prometheus.MustRegister(metrics.RequestCounter)
	prometheus.MustRegister(metrics.RequestDuration)

	logger.Logger = zerolog.New(os.Stdout).With().Timestamp().Logger()

	db := server.InitDB()

	defer db.Close()

	srv := server.New(db)

	srv.Run()

}
