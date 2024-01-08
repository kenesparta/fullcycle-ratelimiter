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
	ByIP LimitValues
}

type LimitValues struct {
	MaxRequests     int
	TimeWindow      int64
	BlockedDuration int64
}

// Config Final Struct Configuration
type Config struct {
	Redis       Redis
	App         App
	RateLimiter RateLimiter
}
