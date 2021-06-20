package config

import (
	"fmt"
	"os"
	"strconv"
)

type EnvConfig struct {
}

func NewEnvConfig() *EnvConfig {
	return &EnvConfig{}
}

func (c *EnvConfig) GRPCPort() int {
	port, err := strconv.Atoi(os.Getenv("GRPC_PORT"))
	if err != nil {
		panic(err)
	}
	return port
}

func (c *EnvConfig) DBConn() string {
	host := "pgsql"
	port := 5432
	user := "postgres"
	password := "todo"
	dbname := "todo"
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
}
