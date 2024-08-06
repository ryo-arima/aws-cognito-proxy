package controller

import (
	"fmt"
	"log"

	"github.com/ryo-arima/aws-cognito-proxy/pkg/config"
	"github.com/spf13/cobra"
)

func InitBootstrapGroupsCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	bootstrapGroupsCmd := &cobra.Command{
		Use:   "groups",
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
	bootstrapGroupsCmd.Flags().StringP("key", "k", "", "cache key")
	return bootstrapGroupsCmd
}

func InitCreateGroupsCmdForAnonymousUser(conf config.BaseConfig) *cobra.Command {
	createGroupsCmd := &cobra.Command{
		Use:   "groups",
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
	createGroupsCmd.Flags().StringP("key", "k", "", "cache key")
	return createGroupsCmd
}

func InitCreateGroupsCmdForAppUser(conf config.BaseConfig) *cobra.Command {
	createGroupsCmd := &cobra.Command{
		Use:   "groups",
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
	createGroupsCmd.Flags().StringP("key", "k", "", "cache key")
	return createGroupsCmd
}

func InitCreateGroupsCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	createGroupsCmd := &cobra.Command{
		Use:   "groups",
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
	createGroupsCmd.Flags().StringP("key", "k", "", "cache key")
	return createGroupsCmd
}

func InitGetGroupsCmdForAnonymousUser(conf config.BaseConfig) *cobra.Command {
	getGroupsCmd := &cobra.Command{
		Use:   "groups",
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
	getGroupsCmd.Flags().StringP("key", "k", "", "cache key")
	return getGroupsCmd
}

func InitGetGroupsCmdForAppUser(conf config.BaseConfig) *cobra.Command {
	getGroupsCmd := &cobra.Command{
		Use:   "groups",
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
	getGroupsCmd.Flags().StringP("key", "k", "", "cache key")
	return getGroupsCmd
}

func InitGetGroupsCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	getGroupsCmd := &cobra.Command{
		Use:   "groups",
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
	getGroupsCmd.Flags().StringP("key", "k", "", "cache key")
	return getGroupsCmd
}

func InitUpdateGroupsCmdForAnonymousUser(conf config.BaseConfig) *cobra.Command {
	updateGroupsCmd := &cobra.Command{
		Use:   "groups",
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
	updateGroupsCmd.Flags().StringP("key", "k", "", "cache key")
	return updateGroupsCmd
}

func InitUpdateGroupsCmdForAppUser(conf config.BaseConfig) *cobra.Command {
	updateGroupsCmd := &cobra.Command{
		Use:   "groups",
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
	updateGroupsCmd.Flags().StringP("key", "k", "", "cache key")
	return updateGroupsCmd
}

func InitUpdateGroupsCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	updateGroupsCmd := &cobra.Command{
		Use:   "groups",
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
	updateGroupsCmd.Flags().StringP("key", "k", "", "cache key")
	return updateGroupsCmd
}

func InitDeleteGroupsCmdForAnonymousUser(conf config.BaseConfig) *cobra.Command {
	deleteGroupsCmd := &cobra.Command{
		Use:   "groups",
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
	deleteGroupsCmd.Flags().StringP("key", "k", "", "cache key")
	return deleteGroupsCmd
}

func InitDeleteGroupsCmdForAppUser(conf config.BaseConfig) *cobra.Command {
	deleteGroupsCmd := &cobra.Command{
		Use:   "groups",
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
	deleteGroupsCmd.Flags().StringP("key", "k", "", "cache key")
	return deleteGroupsCmd
}

func InitDeleteGroupsCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	deleteGroupsCmd := &cobra.Command{
		Use:   "groups",
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
	deleteGroupsCmd.Flags().StringP("key", "k", "", "cache key")
	return deleteGroupsCmd
}