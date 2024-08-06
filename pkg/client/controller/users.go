package controller

import (
	"fmt"
	"log"

	"github.com/ryo-arima/aws-cognito-proxy/pkg/config"
	"github.com/spf13/cobra"
)

func InitBootstrapUsersCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	bootstrapUsersCmd := &cobra.Command{
		Use:   "users",
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
	bootstrapUsersCmd.Flags().StringP("key", "k", "", "cache key")
	return bootstrapUsersCmd
}

func InitCreateUsersCmdForAnonymousUser(conf config.BaseConfig) *cobra.Command {
	createUsersCmd := &cobra.Command{
		Use:   "users",
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
	createUsersCmd.Flags().StringP("key", "k", "", "cache key")
	return createUsersCmd
}

func InitCreateUsersCmdForAppUser(conf config.BaseConfig) *cobra.Command {
	createUsersCmd := &cobra.Command{
		Use:   "users",
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
	createUsersCmd.Flags().StringP("key", "k", "", "cache key")
	return createUsersCmd
}

func InitCreateUsersCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	createUsersCmd := &cobra.Command{
		Use:   "users",
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
	createUsersCmd.Flags().StringP("key", "k", "", "cache key")
	return createUsersCmd
}

func InitGetUsersCmdForAnonymousUser(conf config.BaseConfig) *cobra.Command {
	getUsersCmd := &cobra.Command{
		Use:   "users",
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
	getUsersCmd.Flags().StringP("key", "k", "", "cache key")
	return getUsersCmd
}

func InitGetUsersCmdForAppUser(conf config.BaseConfig) *cobra.Command {
	getUsersCmd := &cobra.Command{
		Use:   "users",
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
	getUsersCmd.Flags().StringP("key", "k", "", "cache key")
	return getUsersCmd
}

func InitGetUsersCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	getUsersCmd := &cobra.Command{
		Use:   "users",
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
	getUsersCmd.Flags().StringP("key", "k", "", "cache key")
	return getUsersCmd
}

func InitUpdateUsersCmdForAnonymousUser(conf config.BaseConfig) *cobra.Command {
	updateUsersCmd := &cobra.Command{
		Use:   "users",
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
	updateUsersCmd.Flags().StringP("key", "k", "", "cache key")
	return updateUsersCmd
}

func InitUpdateUsersCmdForAppUser(conf config.BaseConfig) *cobra.Command {
	updateUsersCmd := &cobra.Command{
		Use:   "users",
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
	updateUsersCmd.Flags().StringP("key", "k", "", "cache key")
	return updateUsersCmd
}

func InitUpdateUsersCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	updateUsersCmd := &cobra.Command{
		Use:   "users",
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
	updateUsersCmd.Flags().StringP("key", "k", "", "cache key")
	return updateUsersCmd
}

func InitDeleteUsersCmdForAnonymousUser(conf config.BaseConfig) *cobra.Command {
	deleteUsersCmd := &cobra.Command{
		Use:   "users",
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
	deleteUsersCmd.Flags().StringP("key", "k", "", "cache key")
	return deleteUsersCmd
}

func InitDeleteUsersCmdForAppUser(conf config.BaseConfig) *cobra.Command {
	deleteUsersCmd := &cobra.Command{
		Use:   "users",
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
	deleteUsersCmd.Flags().StringP("key", "k", "", "cache key")
	return deleteUsersCmd
}

func InitDeleteUsersCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	deleteUsersCmd := &cobra.Command{
		Use:   "users",
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
	deleteUsersCmd.Flags().StringP("key", "k", "", "cache key")
	return deleteUsersCmd
}