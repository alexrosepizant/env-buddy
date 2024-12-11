package cmd

import (
    "fmt"
    "os"
    "github.com/spf13/cobra"
    "envbuddy/encrypt"
)

var RootCmd = &cobra.Command{
    Use:   "envbuddy", package cmd

	import (
		"fmt"
		"os"
		"github.com/spf13/cobra"
		"envbuddy/encrypt"
	)
	
	var RootCmd = &cobra.Command{
		Use:   "envbuddy",
		Short: "EnvBuddy: a tool to manage your .env files",
		Long:  "EnvBuddy helps you manage, version, and secure your .env files.",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("A command must be specified.")
		},
	}
	
	func init() {
		var encryptCmd = &cobra.Command{
			Use:   "encrypt",
			Short: "Encrypt a .env file",
			Run: func(cmd *cobra.Command, args []string) {
				if len(args) < 1 {
					fmt.Println("Please specify a .env file to encrypt.")
					os.Exit(1)
				}
				err := encrypt.EncryptEnvFile(args[0])
				if err != nil {
					fmt.Println("Error:", err)
					os.Exit(1)
				}
				fmt.Println("File encrypted successfully!")
			},
		}
		RootCmd.AddCommand(encryptCmd)
	
		var decryptCmd = &cobra.Command{
			Use:   "decrypt",
			Short: "Decrypt an encrypted .env file",
			Run: func(cmd *cobra.Command, args []string) {
				if len(args) < 1 {
					fmt.Println("Please specify an encrypted .env file to decrypt.")
					os.Exit(1)
				}
				err := encrypt.DecryptEnvFile(args[0])
				if err != nil {
					fmt.Println("Error:", err)
					os.Exit(1)
				}
				fmt.Println("File decrypted successfully!")
			},
		}
		RootCmd.AddCommand(decryptCmd)
	}
	
    Short: "EnvBuddy: a tool to manage your .env files",
    Long:  "EnvBuddy helps you manage, version, and secure your .env files.",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("A command must be specified.")
    },
}

func init() {
    var encryptCmd = &cobra.Command{
        Use:   "encrypt",
        Short: "Encrypt a .env file",
        Run: func(cmd *cobra.Command, args []string) {
            if len(args) < 1 {
                fmt.Println("Please specify a .env file to encrypt.")
                os.Exit(1)
            }
            err := encrypt.EncryptEnvFile(args[0])
            if err != nil {
                fmt.Println("Error:", err)
                os.Exit(1)
            }
            fmt.Println("File encrypted successfully!")
        },
    }
    RootCmd.AddCommand(encryptCmd)

    var decryptCmd = &cobra.Command{
        Use:   "decrypt",
        Short: "Decrypt an encrypted .env file",
        Run: func(cmd *cobra.Command, args []string) {
            if len(args) < 1 {
                fmt.Println("Please specify an encrypted .env file to decrypt.")
                os.Exit(1)
            }
            err := encrypt.DecryptEnvFile(args[0])
            if err != nil {
                fmt.Println("Error:", err)
                os.Exit(1)
            }
            fmt.Println("File decrypted successfully!")
        },
    }
    RootCmd.AddCommand(decryptCmd)
}
