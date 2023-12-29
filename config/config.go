package config

// Redis configuration related to this specific database
type Redis struct {
	Db   int
	Host string
	Port string
}

// App General configuration related to the Application
type App struct {
	Host string
	Port string
}

// RateLimiter properties for configuration
type RateLimiter struct {
	ByIP int64
}

// Config Final Struct Configuration
type Config struct {
	Redis       Redis
	App         App
	RateLimiter RateLimiter
}

func NewConfig(c IConfig) (*Config, error) {
	if err := c.Validate(); err != nil {
		return nil, err
	}
	return c.Get()
}
