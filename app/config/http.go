package config

const (
	serverHostEnv = "SERVER_HOST"
	serverPortEnv = "SERVER_PORT"
)

type Http struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

func (s *Http) Init() {
	s.Host = getStringFromEnv(serverHostEnv, "localhost")
	s.Port = getStringFromEnv(serverPortEnv, "8000")
}
