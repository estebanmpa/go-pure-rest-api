package config

type PgConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DbName   string
	Dialect  string
}

func GetPort() string {
	return ":3000"
}

func GetDatabaseConfig() PgConfig {
	return PgConfig{Host: "localhost", Port: 5432, User: "godbuser", Password: "godbuser123", DbName: "godb", Dialect: "postgres"}
}
