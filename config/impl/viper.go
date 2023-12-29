package impl

import (
	"fmt"
	"os"

	"github.com/kenesparta/fullcycle-ratelimiter/config"
	"github.com/spf13/viper"
)

const fileExtension = "json"

type Viper struct {
	fileName string
}

func NewViper(fileName string) *Viper {
	return &Viper{
		fileName: fileName,
	}
}

func (v *Viper) prepareViper() {
	viper.SetConfigName(v.fileName)
	viper.SetConfigType(fileExtension)
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("fatal error config file: default \n", err)
		os.Exit(1)
	}
}

func (v *Viper) Validate() error {
	return nil
}

func (v *Viper) Get() (*config.Config, error) {
	v.prepareViper()

	return &config.Config{
		Redis: config.Redis{
			Db:   viper.GetInt("redis.db"),
			Host: viper.GetString("redis.host"),
			Port: viper.GetString("redis.port"),
		},
		App: config.App{
			Host: viper.GetString("app.host"),
			Port: viper.GetString("app.port"),
		},
		RateLimiter: config.RateLimiter{
			ByIP: viper.GetInt64("rate_limiter.by_ip"),
		},
	}, nil
}
