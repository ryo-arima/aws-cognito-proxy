package controller

import (
	"fmt"
	"log"

	"github.com/ryo-arima/aws-cognito-proxy/pkg/config"
	"github.com/spf13/cobra"
)

func InitBootstrapPolicyCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	bootstrapPolicyCmd := &cobra.Command{
		Use:   "policy",
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
	bootstrapPolicyCmd.Flags().StringP("key", "k", "", "cache key")
	return bootstrapPolicyCmd
}

func InitCreatePolicyCmdForAnonymousUser(conf config.BaseConfig) *cobra.Command {
	createPolicyCmd := &cobra.Command{
		Use:   "policy",
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
	createPolicyCmd.Flags().StringP("key", "k", "", "cache key")
	return createPolicyCmd
}

func InitCreatePolicyCmdForAppUser(conf config.BaseConfig) *cobra.Command {
	createPolicyCmd := &cobra.Command{
		Use:   "policy",
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
	createPolicyCmd.Flags().StringP("key", "k", "", "cache key")
	return createPolicyCmd
}

func InitCreatePolicyCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	createPolicyCmd := &cobra.Command{
		Use:   "policy",
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
	createPolicyCmd.Flags().StringP("key", "k", "", "cache key")
	return createPolicyCmd
}

func InitGetPolicyCmdForAnonymousUser(conf config.BaseConfig) *cobra.Command {
	getPolicyCmd := &cobra.Command{
		Use:   "policy",
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
	getPolicyCmd.Flags().StringP("key", "k", "", "cache key")
	return getPolicyCmd
}

func InitGetPolicyCmdForAppUser(conf config.BaseConfig) *cobra.Command {
	getPolicyCmd := &cobra.Command{
		Use:   "policy",
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
	getPolicyCmd.Flags().StringP("key", "k", "", "cache key")
	return getPolicyCmd
}

func InitGetPolicyCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	getPolicyCmd := &cobra.Command{
		Use:   "policy",
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
	getPolicyCmd.Flags().StringP("key", "k", "", "cache key")
	return getPolicyCmd
}

func InitUpdatePolicyCmdForAnonymousUser(conf config.BaseConfig) *cobra.Command {
	updatePolicyCmd := &cobra.Command{
		Use:   "policy",
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
	updatePolicyCmd.Flags().StringP("key", "k", "", "cache key")
	return updatePolicyCmd
}

func InitUpdatePolicyCmdForAppUser(conf config.BaseConfig) *cobra.Command {
	updatePolicyCmd := &cobra.Command{
		Use:   "policy",
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
	updatePolicyCmd.Flags().StringP("key", "k", "", "cache key")
	return updatePolicyCmd
}

func InitUpdatePolicyCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	updatePolicyCmd := &cobra.Command{
		Use:   "policy",
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
	updatePolicyCmd.Flags().StringP("key", "k", "", "cache key")
	return updatePolicyCmd
}

func InitDeletePolicyCmdForAnonymousUser(conf config.BaseConfig) *cobra.Command {
	deletePolicyCmd := &cobra.Command{
		Use:   "policy",
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
	deletePolicyCmd.Flags().StringP("key", "k", "", "cache key")
	return deletePolicyCmd
}

func InitDeletePolicyCmdForAppUser(conf config.BaseConfig) *cobra.Command {
	deletePolicyCmd := &cobra.Command{
		Use:   "policy",
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
	deletePolicyCmd.Flags().StringP("key", "k", "", "cache key")
	return deletePolicyCmd
}

func InitDeletePolicyCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	deletePolicyCmd := &cobra.Command{
		Use:   "policy",
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
	deletePolicyCmd.Flags().StringP("key", "k", "", "cache key")
	return deletePolicyCmd
}