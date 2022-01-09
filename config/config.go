package config

const (
	Environment        = "ENVIRONMENT"
	DBConnectionString = "DB_CONNECTION_STRING"
	RedisAddr          = "REDIS_ADDR"
	RedisPass          = "REDIS_PASS"
)

type Config struct{}

func Init() *Config {
	return &Config{}
}

func (c *Config) Environment() string {
	return getStringOrDefault(Environment, "development")
}

func (c *Config) DBConnectionString() string {
	return getStringOrDefault(DBConnectionString, "")
}

func (c *Config) RedisAddr() string {
	return getStringOrDefault(RedisAddr, "")
}

func (c *Config) RedisPass() string {
	return getStringOrDefault(RedisPass, "")
}
