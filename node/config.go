package node

type Config struct {
	DataDir  string
	HttpPort string
}

func NewConfig(configPath string) *Config {

	return nil
}

func defaultConfig() *Config {
	return &Config{
		DataDir:  "./tmp",
		HttpPort: "8088",
	}
}
