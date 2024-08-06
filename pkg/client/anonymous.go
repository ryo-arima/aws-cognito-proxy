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
	getUsersCmdForAnonymousUser := controller.InitGetUsersCmdForAnonymousUser(conf)
	baseCmdForAnonymousUser.Get.AddCommand(getUsersCmdForAnonymousUser)
	getGroupsCmdForAnonymousUser := controller.InitGetGroupsCmdForAnonymousUser(conf)
	baseCmdForAnonymousUser.Get.AddCommand(getGroupsCmdForAnonymousUser)
	getRolesCmdForAnonymousUser := controller.InitGetRolesCmdForAnonymousUser(conf)
	baseCmdForAnonymousUser.Get.AddCommand(getRolesCmdForAnonymousUser)
	getPoliciesCmdForAnonymousUser := controller.InitGetPoliciesCmdForAnonymousUser(conf)
	baseCmdForAnonymousUser.Get.AddCommand(getPoliciesCmdForAnonymousUser)
	rootCmdForAnonymousUser.AddCommand(baseCmdForAnonymousUser.Get)
	
	rootCmdForAnonymousUser.Execute()
}