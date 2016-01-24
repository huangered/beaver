package model

type Config struct {
}

func NewConfig(file string) (c *Config) {
	c = &Config{}
	return c
}
