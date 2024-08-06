package controller

import (
	"fmt"
	"log"

	"github.com/ryo-arima/aws-cognito-proxy/pkg/config"
	"github.com/spf13/cobra"
)

func InitBootstrapRolesCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	bootstrapRolesCmd := &cobra.Command{
		Use:   "roles",
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
	bootstrapRolesCmd.Flags().StringP("key", "k", "", "cache key")
	return bootstrapRolesCmd
}

func InitCreateRolesCmdForAnonymousUser(conf config.BaseConfig) *cobra.Command {
	createRolesCmd := &cobra.Command{
		Use:   "roles",
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
	createRolesCmd.Flags().StringP("key", "k", "", "cache key")
	return createRolesCmd
}

func InitCreateRolesCmdForAppUser(conf config.BaseConfig) *cobra.Command {
	createRolesCmd := &cobra.Command{
		Use:   "roles",
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
	createRolesCmd.Flags().StringP("key", "k", "", "cache key")
	return createRolesCmd
}

func InitCreateRolesCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	createRolesCmd := &cobra.Command{
		Use:   "roles",
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
	createRolesCmd.Flags().StringP("key", "k", "", "cache key")
	return createRolesCmd
}

func InitGetRolesCmdForAnonymousUser(conf config.BaseConfig) *cobra.Command {
	getRolesCmd := &cobra.Command{
		Use:   "roles",
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
	getRolesCmd.Flags().StringP("key", "k", "", "cache key")
	return getRolesCmd
}

func InitGetRolesCmdForAppUser(conf config.BaseConfig) *cobra.Command {
	getRolesCmd := &cobra.Command{
		Use:   "roles",
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
	getRolesCmd.Flags().StringP("key", "k", "", "cache key")
	return getRolesCmd
}

func InitGetRolesCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	getRolesCmd := &cobra.Command{
		Use:   "roles",
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
	getRolesCmd.Flags().StringP("key", "k", "", "cache key")
	return getRolesCmd
}

func InitUpdateRolesCmdForAnonymousUser(conf config.BaseConfig) *cobra.Command {
	updateRolesCmd := &cobra.Command{
		Use:   "roles",
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
	updateRolesCmd.Flags().StringP("key", "k", "", "cache key")
	return updateRolesCmd
}

func InitUpdateRolesCmdForAppUser(conf config.BaseConfig) *cobra.Command {
	updateRolesCmd := &cobra.Command{
		Use:   "roles",
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
	updateRolesCmd.Flags().StringP("key", "k", "", "cache key")
	return updateRolesCmd
}

func InitUpdateRolesCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	updateRolesCmd := &cobra.Command{
		Use:   "roles",
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
	updateRolesCmd.Flags().StringP("key", "k", "", "cache key")
	return updateRolesCmd
}

func InitDeleteRolesCmdForAnonymousUser(conf config.BaseConfig) *cobra.Command {
	deleteRolesCmd := &cobra.Command{
		Use:   "roles",
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
	deleteRolesCmd.Flags().StringP("key", "k", "", "cache key")
	return deleteRolesCmd
}

func InitDeleteRolesCmdForAppUser(conf config.BaseConfig) *cobra.Command {
	deleteRolesCmd := &cobra.Command{
		Use:   "roles",
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
	deleteRolesCmd.Flags().StringP("key", "k", "", "cache key")
	return deleteRolesCmd
}

func InitDeleteRolesCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	deleteRolesCmd := &cobra.Command{
		Use:   "roles",
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
	deleteRolesCmd.Flags().StringP("key", "k", "", "cache key")
	return deleteRolesCmd
}