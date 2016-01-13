package parser

type Config struct {
  volumeIndex string `json:"volume_index"`
freeVolumeIndex string `json:"free_volume_index"`
volumeCacheSize int `json:"volume_cache_size"`
serverId string `json:"server_id"`
index string `json:"index"`
needleMaxSize int `json:"needle_max_size"`
batchMaxNum int `json:"batch_max_num"`

sync int `json:"sync"`

listen string `json:"listen"`
addrs string `json:"addrs"`
}
}

func NewConfig() *Config {
  c := &Config{}
  return c
}
