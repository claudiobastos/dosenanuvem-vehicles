package config

import "go.uber.org/zap"

type Cfg struct {
	Logger *zap.Logger

	Server ServerSettings `mapstructure:"server"`
}

type ServerSettings struct {
	HTTP HTTPServerSettings `mapstructure:"http"`
}

type HTTPServerSettings struct {
	Endpoint string `mapstructure:"endpoint"`
}

func New() *Cfg {
	cfg := &Cfg{
		Logger: zap.Must(zap.NewDevelopment()),
		Server: ServerSettings{
			HTTP: HTTPServerSettings{},
		},
	}

	return cfg
}
