package config

import "fmt"

type StaticConfig struct {
}

func NewStaticConfig() *StaticConfig {
	return &StaticConfig{}
}

func (c *StaticConfig) GRPCPort() int {
	return 50051
}

func (c *StaticConfig) DBConn() string {
	host := "localhost"
	port := 5432
	user := "postgres"
	password := "todo"
	dbname := "todo"
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
}
