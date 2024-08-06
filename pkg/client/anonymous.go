package client

import (
	"github.com/ryo-arima/aws-cognito-proxy/pkg/client/controller"
	"github.com/ryo-arima/aws-cognito-proxy/pkg/config"
	"github.com/spf13/cobra"
)

type BaseCmdForAnonymousUser struct {
	Get       *cobra.Command
}

func InitRootCmdForAnonymousUser() *cobra.Command {
	var rootCmdForAnonymousUser = &cobra.Command{
		Use:   "aws-cognito-proxy-anonymous",
		Short: "'aws-cognito-proxy' is a CLI tool to manage anniversaries",
		Long:  `''`,
	}
	return rootCmdForAnonymousUser
}

func InitBaseCmdForAnonymousUser() BaseCmdForAnonymousUser {
	getCmd := &cobra.Command{
		Use:   "get",
		Short: "get the value of a key",
		Long:  "get the value of a key",
	}
	baseCmdForAnonymousUser := BaseCmdForAnonymousUser{
		Get:       getCmd,
	}
	return baseCmdForAnonymousUser
}

func ClientForAnonymousUser(conf config.BaseConfig) {
	rootCmdForAnonymousUser := InitRootCmdForAnonymousUser()
	rootCmdForAnonymousUser.CompletionOptions.HiddenDefaultCmd = true
	baseCmdForAnonymousUser := InitBaseCmdForAnonymousUser()

	//get
	getUserCmdForAnonymousUser := controller.InitGetUserCmdForAnonymousUser(conf)
	baseCmdForAnonymousUser.Get.AddCommand(getUserCmdForAnonymousUser)
	getGroupCmdForAnonymousUser := controller.InitGetGroupCmdForAnonymousUser(conf)
	baseCmdForAnonymousUser.Get.AddCommand(getGroupCmdForAnonymousUser)
	getRoleCmdForAnonymousUser := controller.InitGetRoleCmdForAnonymousUser(conf)
	baseCmdForAnonymousUser.Get.AddCommand(getRoleCmdForAnonymousUser)
	getPolicyCmdForAnonymousUser := controller.InitGetPolicyCmdForAnonymousUser(conf)
	baseCmdForAnonymousUser.Get.AddCommand(getPolicyCmdForAnonymousUser)
	rootCmdForAnonymousUser.AddCommand(baseCmdForAnonymousUser.Get)
	
	rootCmdForAnonymousUser.Execute()
}