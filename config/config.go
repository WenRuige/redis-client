package config

const (
	DefaultHost = "127.0.0.1"
	DefaultPort = 6379
)

type Config struct {
	Host string
	Port int
}
