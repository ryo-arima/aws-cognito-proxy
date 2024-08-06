package server

import "github.com/ryo-arima/aws-cognito-proxy/pkg/config"

func Main(conf config.BaseConfig) {
	router := InitRouter(conf)
	router.Run(":8000")
}