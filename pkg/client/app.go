package client

import (
	"github.com/ryo-arima/aws-cognito-proxy/pkg/client/controller"
	"github.com/ryo-arima/aws-cognito-proxy/pkg/config"
	"github.com/spf13/cobra"
)

type BaseCmdForAppUser struct {
	Create    *cobra.Command
	Get       *cobra.Command
	Update    *cobra.Command
	Delete    *cobra.Command
}

func InitRootCmdForAppUser() *cobra.Command {
	var rootCmdForAppUser = &cobra.Command{
		Use:   "aws-cognito-proxy-app",
		Short: "'aws-cognito-proxy' is a CLI tool to manage anniversaries",
		Long:  `''`,
	}
	return rootCmdForAppUser
}

func InitBaseCmdForAppUser() BaseCmdForAppUser {
	createCmd := &cobra.Command{
		Use:   "create",
		Short: "create the value of a key",
		Long:  "create the value of a key",
	}
	getCmd := &cobra.Command{
		Use:   "get",
		Short: "get the value of a key",
		Long:  "get the value of a key",
	}
	updateCmd := &cobra.Command{
		Use:   "update",
		Short: "update the value of a key",
		Long:  "update the value of a key",
	}
	deleteCmd := &cobra.Command{
		Use:   "delete",
		Short: "delete the value of a key",
		Long:  "delete the value of a key",
	}
	baseCmdForAppUser := BaseCmdForAppUser{
		Create:    createCmd,
		Get:       getCmd,
		Update:    updateCmd,
		Delete:    deleteCmd,
	}
	return baseCmdForAppUser
}

func ClientForAppUser(conf config.BaseConfig) {
	rootCmdForAppUser := InitRootCmdForAppUser()
	rootCmdForAppUser.CompletionOptions.HiddenDefaultCmd = true
	baseCmdForAppUser := InitBaseCmdForAppUser()

	//create
	createUsersCmdForAppUser := controller.InitCreateUsersCmdForAppUser(conf)
	baseCmdForAppUser.Create.AddCommand(createUsersCmdForAppUser)
	createGroupsCmdForAppUser := controller.InitCreateGroupsCmdForAppUser(conf)
	baseCmdForAppUser.Create.AddCommand(createGroupsCmdForAppUser)
	createRolesCmdForAppUser := controller.InitCreateRolesCmdForAppUser(conf)
	baseCmdForAppUser.Create.AddCommand(createRolesCmdForAppUser)
	createPoliciesCmdForAppUser := controller.InitCreatePoliciesCmdForAppUser(conf)
	baseCmdForAppUser.Create.AddCommand(createPoliciesCmdForAppUser)
	rootCmdForAppUser.AddCommand(baseCmdForAppUser.Create)
	
	//get
	getUsersCmdForAppUser := controller.InitGetUsersCmdForAppUser(conf)
	baseCmdForAppUser.Get.AddCommand(getUsersCmdForAppUser)
	getGroupsCmdForAppUser := controller.InitGetGroupsCmdForAppUser(conf)
	baseCmdForAppUser.Get.AddCommand(getGroupsCmdForAppUser)
	getRolesCmdForAppUser := controller.InitGetRolesCmdForAppUser(conf)
	baseCmdForAppUser.Get.AddCommand(getRolesCmdForAppUser)
	getPoliciesCmdForAppUser := controller.InitGetPoliciesCmdForAppUser(conf)
	baseCmdForAppUser.Get.AddCommand(getPoliciesCmdForAppUser)
	rootCmdForAppUser.AddCommand(baseCmdForAppUser.Get)
	
	//update
	updateUsersCmdForAppUser := controller.InitUpdateUsersCmdForAppUser(conf)
	baseCmdForAppUser.Update.AddCommand(updateUsersCmdForAppUser)
	updateGroupsCmdForAppUser := controller.InitUpdateGroupsCmdForAppUser(conf)
	baseCmdForAppUser.Update.AddCommand(updateGroupsCmdForAppUser)
	updateRolesCmdForAppUser := controller.InitUpdateRolesCmdForAppUser(conf)
	baseCmdForAppUser.Update.AddCommand(updateRolesCmdForAppUser)
	updatePoliciesCmdForAppUser := controller.InitUpdatePoliciesCmdForAppUser(conf)
	baseCmdForAppUser.Update.AddCommand(updatePoliciesCmdForAppUser)
	rootCmdForAppUser.AddCommand(baseCmdForAppUser.Update)
	
	//delete
	deleteUsersCmdForAppUser := controller.InitDeleteUsersCmdForAppUser(conf)
	baseCmdForAppUser.Delete.AddCommand(deleteUsersCmdForAppUser)
	deleteGroupsCmdForAppUser := controller.InitDeleteGroupsCmdForAppUser(conf)
	baseCmdForAppUser.Delete.AddCommand(deleteGroupsCmdForAppUser)
	deleteRolesCmdForAppUser := controller.InitDeleteRolesCmdForAppUser(conf)
	baseCmdForAppUser.Delete.AddCommand(deleteRolesCmdForAppUser)
	deletePoliciesCmdForAppUser := controller.InitDeletePoliciesCmdForAppUser(conf)
	baseCmdForAppUser.Delete.AddCommand(deletePoliciesCmdForAppUser)
	rootCmdForAppUser.AddCommand(baseCmdForAppUser.Delete)

	rootCmdForAppUser.Execute()
}