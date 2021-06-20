package config

type Config interface {
	GRPCPort() int
	DBConn() string
}
