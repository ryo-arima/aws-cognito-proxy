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
	createUserCmdForAppUser := controller.InitCreateUserCmdForAppUser(conf)
	baseCmdForAppUser.Create.AddCommand(createUserCmdForAppUser)
	createGroupCmdForAppUser := controller.InitCreateGroupCmdForAppUser(conf)
	baseCmdForAppUser.Create.AddCommand(createGroupCmdForAppUser)
	createRoleCmdForAppUser := controller.InitCreateRoleCmdForAppUser(conf)
	baseCmdForAppUser.Create.AddCommand(createRoleCmdForAppUser)
	createPolicyCmdForAppUser := controller.InitCreatePolicyCmdForAppUser(conf)
	baseCmdForAppUser.Create.AddCommand(createPolicyCmdForAppUser)
	rootCmdForAppUser.AddCommand(baseCmdForAppUser.Create)
	
	//get
	getUserCmdForAppUser := controller.InitGetUserCmdForAppUser(conf)
	baseCmdForAppUser.Get.AddCommand(getUserCmdForAppUser)
	getGroupCmdForAppUser := controller.InitGetGroupCmdForAppUser(conf)
	baseCmdForAppUser.Get.AddCommand(getGroupCmdForAppUser)
	getRoleCmdForAppUser := controller.InitGetRoleCmdForAppUser(conf)
	baseCmdForAppUser.Get.AddCommand(getRoleCmdForAppUser)
	getPolicyCmdForAppUser := controller.InitGetPolicyCmdForAppUser(conf)
	baseCmdForAppUser.Get.AddCommand(getPolicyCmdForAppUser)
	rootCmdForAppUser.AddCommand(baseCmdForAppUser.Get)
	
	//update
	updateUserCmdForAppUser := controller.InitUpdateUserCmdForAppUser(conf)
	baseCmdForAppUser.Update.AddCommand(updateUserCmdForAppUser)
	updateGroupCmdForAppUser := controller.InitUpdateGroupCmdForAppUser(conf)
	baseCmdForAppUser.Update.AddCommand(updateGroupCmdForAppUser)
	updateRoleCmdForAppUser := controller.InitUpdateRoleCmdForAppUser(conf)
	baseCmdForAppUser.Update.AddCommand(updateRoleCmdForAppUser)
	updatePolicyCmdForAppUser := controller.InitUpdatePolicyCmdForAppUser(conf)
	baseCmdForAppUser.Update.AddCommand(updatePolicyCmdForAppUser)
	rootCmdForAppUser.AddCommand(baseCmdForAppUser.Update)
	
	//delete
	deleteUserCmdForAppUser := controller.InitDeleteUserCmdForAppUser(conf)
	baseCmdForAppUser.Delete.AddCommand(deleteUserCmdForAppUser)
	deleteGroupCmdForAppUser := controller.InitDeleteGroupCmdForAppUser(conf)
	baseCmdForAppUser.Delete.AddCommand(deleteGroupCmdForAppUser)
	deleteRoleCmdForAppUser := controller.InitDeleteRoleCmdForAppUser(conf)
	baseCmdForAppUser.Delete.AddCommand(deleteRoleCmdForAppUser)
	deletePolicyCmdForAppUser := controller.InitDeletePolicyCmdForAppUser(conf)
	baseCmdForAppUser.Delete.AddCommand(deletePolicyCmdForAppUser)
	rootCmdForAppUser.AddCommand(baseCmdForAppUser.Delete)

	rootCmdForAppUser.Execute()
}