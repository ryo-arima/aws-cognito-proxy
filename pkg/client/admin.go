package client

import (
	"github.com/ryo-arima/aws-cognito-proxy/pkg/client/controller"
	"github.com/ryo-arima/aws-cognito-proxy/pkg/config"
	"github.com/spf13/cobra"
)

type BaseCmdForAdminUser struct {
	Bootstrap *cobra.Command
	Create    *cobra.Command
	Get       *cobra.Command
	Update    *cobra.Command
	Delete    *cobra.Command
}

func InitRootCmdForAdminUser() *cobra.Command {
	var rootCmdForAdminUser = &cobra.Command{
		Use:   "aws-cognito-proxy-admin",
		Short: "'aws-cognito-proxy' is a CLI tool to manage anniversaries",
		Long:  `''`,
	}
	return rootCmdForAdminUser
}

func InitBaseCmdForAdminUser() BaseCmdForAdminUser {
	bootstrapCmd := &cobra.Command{
		Use:   "bootstrap",
		Short: "bootstrap the value of a key",
		Long:  "bootstrap the value of a key",
	}
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
	baseCmdForAdminUser := BaseCmdForAdminUser{
		Bootstrap: bootstrapCmd,
		Create:    createCmd,
		Get:       getCmd,
		Update:    updateCmd,
		Delete:    deleteCmd,
	}
	return baseCmdForAdminUser
}

func ClientForAdminUser(conf config.BaseConfig) {
	rootCmdForAdminUser := InitRootCmdForAdminUser()
	rootCmdForAdminUser.CompletionOptions.HiddenDefaultCmd = true
	baseCmdForAdminUser := InitBaseCmdForAdminUser()

	//bootstrap
	bootstrapUsersCmdForAdminUser := controller.InitBootstrapUsersCmdForAdminUser(conf)
	baseCmdForAdminUser.Bootstrap.AddCommand(bootstrapUsersCmdForAdminUser)
	bootstrapGroupsCmdForAdminUser := controller.InitBootstrapGroupsCmdForAdminUser(conf)
	baseCmdForAdminUser.Bootstrap.AddCommand(bootstrapGroupsCmdForAdminUser)
	bootstrapRolesCmdForAdminUser := controller.InitBootstrapRolesCmdForAdminUser(conf)
	baseCmdForAdminUser.Bootstrap.AddCommand(bootstrapRolesCmdForAdminUser)
	bootstrapPoliciesCmdForAdminUser := controller.InitBootstrapPoliciesCmdForAdminUser(conf)
	baseCmdForAdminUser.Bootstrap.AddCommand(bootstrapPoliciesCmdForAdminUser)
	rootCmdForAdminUser.AddCommand(baseCmdForAdminUser.Bootstrap)

	//create
	createUsersCmdForAdminUser := controller.InitCreateUsersCmdForAdminUser(conf)
	baseCmdForAdminUser.Create.AddCommand(createUsersCmdForAdminUser)
	createGroupsCmdForAdminUser := controller.InitCreateGroupsCmdForAdminUser(conf)
	baseCmdForAdminUser.Create.AddCommand(createGroupsCmdForAdminUser)
	createRolesCmdForAdminUser := controller.InitCreateRolesCmdForAdminUser(conf)
	baseCmdForAdminUser.Create.AddCommand(createRolesCmdForAdminUser)
	createPoliciesCmdForAdminUser := controller.InitCreatePoliciesCmdForAdminUser(conf)
	baseCmdForAdminUser.Create.AddCommand(createPoliciesCmdForAdminUser)
	rootCmdForAdminUser.AddCommand(baseCmdForAdminUser.Create)

	//get
	getUsersCmdForAdminUser := controller.InitGetUsersCmdForAdminUser(conf)
	baseCmdForAdminUser.Get.AddCommand(getUsersCmdForAdminUser)
	getGroupsCmdForAdminUser := controller.InitGetGroupsCmdForAdminUser(conf)
	baseCmdForAdminUser.Get.AddCommand(getGroupsCmdForAdminUser)
	getRolesCmdForAdminUser := controller.InitGetRolesCmdForAdminUser(conf)
	baseCmdForAdminUser.Get.AddCommand(getRolesCmdForAdminUser)
	getPoliciesCmdForAdminUser := controller.InitGetPoliciesCmdForAdminUser(conf)
	baseCmdForAdminUser.Get.AddCommand(getPoliciesCmdForAdminUser)
	rootCmdForAdminUser.AddCommand(baseCmdForAdminUser.Get)

	//update
	updateUsersCmdForAdminUser := controller.InitUpdateUsersCmdForAdminUser(conf)
	baseCmdForAdminUser.Update.AddCommand(updateUsersCmdForAdminUser)
	updateGroupsCmdForAdminUser := controller.InitUpdateGroupsCmdForAdminUser(conf)
	baseCmdForAdminUser.Update.AddCommand(updateGroupsCmdForAdminUser)
	updateRolesCmdForAdminUser := controller.InitUpdateRolesCmdForAdminUser(conf)
	baseCmdForAdminUser.Update.AddCommand(updateRolesCmdForAdminUser)
	updatePoliciesCmdForAdminUser := controller.InitUpdatePoliciesCmdForAdminUser(conf)
	baseCmdForAdminUser.Update.AddCommand(updatePoliciesCmdForAdminUser)
	rootCmdForAdminUser.AddCommand(baseCmdForAdminUser.Update)
	
	//delete
	deleteUsersCmdForAdminUser := controller.InitDeleteUsersCmdForAdminUser(conf)
	baseCmdForAdminUser.Delete.AddCommand(deleteUsersCmdForAdminUser)
	deleteGroupsCmdForAdminUser := controller.InitDeleteGroupsCmdForAdminUser(conf)
	baseCmdForAdminUser.Delete.AddCommand(deleteGroupsCmdForAdminUser)
	deleteRolesCmdForAdminUser := controller.InitDeleteRolesCmdForAdminUser(conf)
	baseCmdForAdminUser.Delete.AddCommand(deleteRolesCmdForAdminUser)
	deletePoliciesCmdForAdminUser := controller.InitDeletePoliciesCmdForAdminUser(conf)
	baseCmdForAdminUser.Delete.AddCommand(deletePoliciesCmdForAdminUser)
	rootCmdForAdminUser.AddCommand(baseCmdForAdminUser.Delete)

	rootCmdForAdminUser.Execute()
}