package controller

import (
	"fmt"
	"log"

	"github.com/ryo-arima/aws-cognito-proxy/pkg/config"
	"github.com/spf13/cobra"
)

func InitBootstrapRoleCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	bootstrapRoleCmd := &cobra.Command{
		Use:   "role",
		Short: "bootstrap the value of a key",
		Long:  "bootstrap the value of a key",
		Run: func(cmd *cobra.Command, args []string) {
			option, err := cmd.Flags().GetString("key")
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(option)
		},
	}
	bootstrapRoleCmd.Flags().StringP("key", "k", "", "cache key")
	return bootstrapRoleCmd
}

func InitCreateRoleCmdForAnonymousUser(conf config.BaseConfig) *cobra.Command {
	createRoleCmd := &cobra.Command{
		Use:   "role",
		Short: "create the value of a key",
		Long:  "create the value of a key",
		Run: func(cmd *cobra.Command, args []string) {
			option, err := cmd.Flags().GetString("key")
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(option)
		},
	}
	createRoleCmd.Flags().StringP("key", "k", "", "cache key")
	return createRoleCmd
}

func InitCreateRoleCmdForAppUser(conf config.BaseConfig) *cobra.Command {
	createRoleCmd := &cobra.Command{
		Use:   "role",
		Short: "create the value of a key",
		Long:  "create the value of a key",
		Run: func(cmd *cobra.Command, args []string) {
			option, err := cmd.Flags().GetString("key")
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(option)
		},
	}
	createRoleCmd.Flags().StringP("key", "k", "", "cache key")
	return createRoleCmd
}

func InitCreateRoleCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	createRoleCmd := &cobra.Command{
		Use:   "role",
		Short: "create the value of a key",
		Long:  "create the value of a key",
		Run: func(cmd *cobra.Command, args []string) {
			option, err := cmd.Flags().GetString("key")
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(option)
		},
	}
	createRoleCmd.Flags().StringP("key", "k", "", "cache key")
	return createRoleCmd
}

func InitGetRoleCmdForAnonymousUser(conf config.BaseConfig) *cobra.Command {
	getRoleCmd := &cobra.Command{
		Use:   "role",
		Short: "get the value of a key",
		Long:  "get the value of a key",
		Run: func(cmd *cobra.Command, args []string) {
			option, err := cmd.Flags().GetString("key")
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(option)
		},
	}
	getRoleCmd.Flags().StringP("key", "k", "", "cache key")
	return getRoleCmd
}

func InitGetRoleCmdForAppUser(conf config.BaseConfig) *cobra.Command {
	getRoleCmd := &cobra.Command{
		Use:   "role",
		Short: "get the value of a key",
		Long:  "get the value of a key",
		Run: func(cmd *cobra.Command, args []string) {
			option, err := cmd.Flags().GetString("key")
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(option)
		},
	}
	getRoleCmd.Flags().StringP("key", "k", "", "cache key")
	return getRoleCmd
}

func InitGetRoleCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	getRoleCmd := &cobra.Command{
		Use:   "role",
		Short: "get the value of a key",
		Long:  "get the value of a key",
		Run: func(cmd *cobra.Command, args []string) {
			option, err := cmd.Flags().GetString("key")
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(option)
		},
	}
	getRoleCmd.Flags().StringP("key", "k", "", "cache key")
	return getRoleCmd
}

func InitUpdateRoleCmdForAnonymousUser(conf config.BaseConfig) *cobra.Command {
	updateRoleCmd := &cobra.Command{
		Use:   "role",
		Short: "update the value of a key",
		Long:  "udpate the value of a key",
		Run: func(cmd *cobra.Command, args []string) {
			option, err := cmd.Flags().GetString("key")
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(option)
		},
	}
	updateRoleCmd.Flags().StringP("key", "k", "", "cache key")
	return updateRoleCmd
}

func InitUpdateRoleCmdForAppUser(conf config.BaseConfig) *cobra.Command {
	updateRoleCmd := &cobra.Command{
		Use:   "role",
		Short: "update the value of a key",
		Long:  "udpate the value of a key",
		Run: func(cmd *cobra.Command, args []string) {
			option, err := cmd.Flags().GetString("key")
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(option)
		},
	}
	updateRoleCmd.Flags().StringP("key", "k", "", "cache key")
	return updateRoleCmd
}

func InitUpdateRoleCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	updateRoleCmd := &cobra.Command{
		Use:   "role",
		Short: "update the value of a key",
		Long:  "update the value of a key",
		Run: func(cmd *cobra.Command, args []string) {
			option, err := cmd.Flags().GetString("key")
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(option)
		},
	}
	updateRoleCmd.Flags().StringP("key", "k", "", "cache key")
	return updateRoleCmd
}

func InitDeleteRoleCmdForAnonymousUser(conf config.BaseConfig) *cobra.Command {
	deleteRoleCmd := &cobra.Command{
		Use:   "role",
		Short: "delete the value of a key",
		Long:  "delete the value of a key",
		Run: func(cmd *cobra.Command, args []string) {
			option, err := cmd.Flags().GetString("key")
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(option)
		},
	}
	deleteRoleCmd.Flags().StringP("key", "k", "", "cache key")
	return deleteRoleCmd
}

func InitDeleteRoleCmdForAppUser(conf config.BaseConfig) *cobra.Command {
	deleteRoleCmd := &cobra.Command{
		Use:   "role",
		Short: "delete the value of a key",
		Long:  "delete the value of a key",
		Run: func(cmd *cobra.Command, args []string) {
			option, err := cmd.Flags().GetString("key")
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(option)
		},
	}
	deleteRoleCmd.Flags().StringP("key", "k", "", "cache key")
	return deleteRoleCmd
}

func InitDeleteRoleCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	deleteRoleCmd := &cobra.Command{
		Use:   "role",
		Short: "delete the value of a key",
		Long:  "delete the value of a key",
		Run: func(cmd *cobra.Command, args []string) {
			option, err := cmd.Flags().GetString("key")
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(option)
		},
	}
	deleteRoleCmd.Flags().StringP("key", "k", "", "cache key")
	return deleteRoleCmd
}