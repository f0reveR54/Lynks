package main

import (
	"stepic-go-basic/cashe/pkg/server"
)

func main() {

	rd := server.InitDB()

	srv := server.New(rd)

	srv.Run()

}
