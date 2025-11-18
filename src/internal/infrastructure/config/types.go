package config



type IConfigService interface {
	GetConfig() (*Config, error)
} 

type Config struct {
	Server Server
}

type Server struct {
	Host string
	Port string
	Environment string
}