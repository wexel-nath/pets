package database

import (
	"database/sql"
	"fmt"

	"github.com/wexel-nath/pets/pkg/config"
	"github.com/wexel-nath/pets/pkg/logger"

	_ "github.com/lib/pq"
)

var (
	connection *sql.DB
)

type PostgresConfig struct {
	Host string
	User string
	Pass string
	Name string
}

func NewPostgresConnection(config PostgresConfig) (*sql.DB, error) {
	source := fmt.Sprintf(
		"host='%s' user='%s' password='%s' dbname='%s' sslmode='disable'",
		config.Host,
		config.User,
		config.Pass,
		config.Name,
	)

	logger.Info("Creating new postgres connection [%s]", source)
	return sql.Open("postgres", source)
}

func GetConnection() *sql.DB {
	if connection == nil || connection.Ping() != nil {
		db, err := NewPostgresConnection(PostgresConfig{
			Host: config.GetDBHost(),
			User: config.GetDBUser(),
			Pass: config.GetDBPass(),
			Name: config.GetDBName(),
		})
		if err != nil {
			panic(err)
		}
		connection = db
	}
	return connection
}
