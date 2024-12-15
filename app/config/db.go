package config

const (
	dbHostEnv     = "DB_HOST"
	dbPortEnv     = "DB_PORT"
	dbNameEnv     = "DB_NAME"
	dbUserEnv     = "DB_USER"
	dbPasswordEnv = "DB_PASSWORD"
)

type DB struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Name     string `json:"name"`
	User     string `json:"user"`
	Password string `json:"password"`
}

func (p *DB) Init() {
	p.Host = getStringFromEnv(dbHostEnv, "localhost")

	p.Port = getStringFromEnv(dbPortEnv, "5439")

	p.Name = getStringFromEnv(dbNameEnv, "postgres")

	p.User = getStringFromEnv(dbUserEnv, "postgres")

	p.Password = getStringFromEnv(dbPasswordEnv, "postgres")
}
