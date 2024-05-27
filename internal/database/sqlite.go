package database

import (
	//"fmt"

	"github.com/wupyweb/realworld-kit/ent"
	"github.com/wupyweb/realworld-kit/internal/config"

	_ "github.com/mattn/go-sqlite3"
)


func NewSqlite(conf *config.Config) (*ent.Client, error) {
	// dsn := fmt.Sprintf("%s.sqlite?_fk=1", "test")
	return ent.Open("sqlite3", "test.sqlite?_fk=1")
}