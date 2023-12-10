package config

import (
	"fmt"

	"github.com/takahiromitsui/go-tasks/utils"
)

func GetPostgresURI() string {
	host := utils.GetEnvWithDefault("PG_HOST", "localhost")
	user := utils.GetEnvWithDefault("PG_USER", "postgres")
	password := utils.GetEnvWithDefault("PG_PASSWORD", "password")
	dbname := utils.GetEnvWithDefault("PG_DBNAME", "go_tasks")
	port := utils.GetEnvWithDefault("PG_PORT", "5432")

	// Format the PostgreSQL URI
	uri := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, port)

	return uri
}