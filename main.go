package main

import (
	"github.com/wupyweb/realworld-kit/internal/config"
	"github.com/wupyweb/realworld-kit/internal/server"
)

func main() {
	// client, err := ent.Open("sqlite3", "test.sqlite?_fk=1")
	// if err != nil {
	// 	log.Fatalf("failed opening connection to sqlite: %v", err)
	// }
	// defer client.Close()

	config := &config.Config{}
	server, _ := server.NewServer(config)

	server.Serve()
}

