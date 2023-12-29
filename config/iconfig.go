package config

type IConfig interface {
	Get() (*Config, error)
	Validate() error
}
