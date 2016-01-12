package parser

type Config struct {
  IndexRoot string `json:"indexRoot"`
}

func NewConfig() *Config {
  c := &Config{}
  return c
}
