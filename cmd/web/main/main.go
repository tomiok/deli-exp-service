package main

import (
	"fmt"
	"github.com/deli/exp-service/commons/logs"
	"github.com/deli/exp-service/engine"
)

const (
	port   = "8081"
	dbPath = "%s:%s@tcp(%s:3306)/deli_experiences?parseTime=true"
)

func main() {
	logs.InitDefault()
	e := engine.New(fmt.Sprintf(dbPath, "root", "root", "localhost"))

	srv := NewServer(&specHandler{e}, port)

	srv.Start()
}
