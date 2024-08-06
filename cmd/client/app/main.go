package main

import (
	"github.com/ryo-arima/aws-cognito-proxy/pkg/client"
	"github.com/ryo-arima/aws-cognito-proxy/pkg/config"
)

func main() {
	conf := config.NewBaseConfig()
	client.ClientForAppUser(conf)
}