package controller

import (
	"fmt"
	"log"

	"github.com/ryo-arima/aws-cognito-proxy/pkg/config"
	"github.com/spf13/cobra"
)

func InitBootstrapPoliciesCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	bootstrapPoliciesCmd := &cobra.Command{
		Use:   "policies",
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
	bootstrapPoliciesCmd.Flags().StringP("key", "k", "", "cache key")
	return bootstrapPoliciesCmd
}

func InitCreatePoliciesCmdForAnonymousUser(conf config.BaseConfig) *cobra.Command {
	createPoliciesCmd := &cobra.Command{
		Use:   "policies",
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
	createPoliciesCmd.Flags().StringP("key", "k", "", "cache key")
	return createPoliciesCmd
}

func InitCreatePoliciesCmdForAppUser(conf config.BaseConfig) *cobra.Command {
	createPoliciesCmd := &cobra.Command{
		Use:   "policies",
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
	createPoliciesCmd.Flags().StringP("key", "k", "", "cache key")
	return createPoliciesCmd
}

func InitCreatePoliciesCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	createPoliciesCmd := &cobra.Command{
		Use:   "policies",
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
	createPoliciesCmd.Flags().StringP("key", "k", "", "cache key")
	return createPoliciesCmd
}

func InitGetPoliciesCmdForAnonymousUser(conf config.BaseConfig) *cobra.Command {
	getPoliciesCmd := &cobra.Command{
		Use:   "policies",
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
	getPoliciesCmd.Flags().StringP("key", "k", "", "cache key")
	return getPoliciesCmd
}

func InitGetPoliciesCmdForAppUser(conf config.BaseConfig) *cobra.Command {
	getPoliciesCmd := &cobra.Command{
		Use:   "policies",
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
	getPoliciesCmd.Flags().StringP("key", "k", "", "cache key")
	return getPoliciesCmd
}

func InitGetPoliciesCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	getPoliciesCmd := &cobra.Command{
		Use:   "policies",
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
	getPoliciesCmd.Flags().StringP("key", "k", "", "cache key")
	return getPoliciesCmd
}

func InitUpdatePoliciesCmdForAnonymousUser(conf config.BaseConfig) *cobra.Command {
	updatePoliciesCmd := &cobra.Command{
		Use:   "policies",
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
	updatePoliciesCmd.Flags().StringP("key", "k", "", "cache key")
	return updatePoliciesCmd
}

func InitUpdatePoliciesCmdForAppUser(conf config.BaseConfig) *cobra.Command {
	updatePoliciesCmd := &cobra.Command{
		Use:   "policies",
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
	updatePoliciesCmd.Flags().StringP("key", "k", "", "cache key")
	return updatePoliciesCmd
}

func InitUpdatePoliciesCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	updatePoliciesCmd := &cobra.Command{
		Use:   "policies",
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
	updatePoliciesCmd.Flags().StringP("key", "k", "", "cache key")
	return updatePoliciesCmd
}

func InitDeletePoliciesCmdForAnonymousUser(conf config.BaseConfig) *cobra.Command {
	deletePoliciesCmd := &cobra.Command{
		Use:   "policies",
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
	deletePoliciesCmd.Flags().StringP("key", "k", "", "cache key")
	return deletePoliciesCmd
}

func InitDeletePoliciesCmdForAppUser(conf config.BaseConfig) *cobra.Command {
	deletePoliciesCmd := &cobra.Command{
		Use:   "policies",
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
	deletePoliciesCmd.Flags().StringP("key", "k", "", "cache key")
	return deletePoliciesCmd
}

func InitDeletePoliciesCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	deletePoliciesCmd := &cobra.Command{
		Use:   "policies",
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
	deletePoliciesCmd.Flags().StringP("key", "k", "", "cache key")
	return deletePoliciesCmd
}