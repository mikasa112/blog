package pkg

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

type ServerSetting struct {
	Mode        string
	Port        string
	PageSize    int
	MaxPageSize int
}

type DatabasesSetting struct {
	DBName   string
	Username string
	Password string
}

type JwtSetting struct {
	Secret string
	Issuer string
	Expire time.Duration
}

var Sc ServerSetting

var Ds DatabasesSetting

var Js JwtSetting

func ReadConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("configs")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	//解析key到struct
	err = viper.UnmarshalKey("Server", &Sc)
	if err != nil {
		Log.Sugar().Fatalf("unable to decode into struct, %v", err)
	}
	err = viper.UnmarshalKey("Database", &Ds)
	if err != nil {
		Log.Sugar().Fatalf("unable to decode into struct, %v", err)
	}
	err = viper.UnmarshalKey("JWT", &Js)
	if err != nil {
		Log.Sugar().Fatalf("unable to decode into struct, %v", err)
	}
	Js.Expire *= time.Second
}
