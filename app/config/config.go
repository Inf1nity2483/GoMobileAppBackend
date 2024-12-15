package config

type Config struct {
	Http Http
	DB   DB
}

func New() *Config {
	return &Config{
		Http: Http{},
		DB:   DB{},
	}
}

func (c *Config) Init() {
	c.Http.Init()
	c.DB.Init()
}
