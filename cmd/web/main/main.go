package main

import "github.com/deli/exp-service/commons/logs"

const port = "8081"

func main() {
	logs.InitDefault()
	srv := NewServer(nil, port)

	srv.Start()
}
