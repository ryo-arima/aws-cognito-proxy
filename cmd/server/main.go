package main

import (
	"github.com/ryo-arima/aws-cognito-proxy/pkg/config"
	"github.com/ryo-arima/aws-cognito-proxy/pkg/server"
)

func main() {
	conf := config.NewBaseConfig()
	server.Main(conf)
}