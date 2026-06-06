package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage configuration",
	Long: `Manage application configuration.

Examples:
  agy config list          # List all configuration
  agy config get key       # Get specific configuration value
  agy config set key value # Set configuration value
  agy config edit          # Edit configuration in editor`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			return
		}
		
		switch args[0] {
		case "list":
			// TODO: List configuration
			fmt.Println("Configuration:")
		case "get":
			if len(args) < 2 {
				fmt.Println("Error: key required")
				return
			}
			// TODO: Get configuration value
			fmt.Printf("%s: <value>\n", args[1])
		case "set":
			if len(args) < 3 {
				fmt.Println("Error: key and value required")
				return
			}
			// TODO: Set configuration value
			fmt.Printf("Set %s = %s\n", args[1], args[2])
		case "edit":
			// TODO: Edit configuration in editor
			fmt.Println("Opening editor...")
		default:
			fmt.Printf("Unknown command: %s\n", args[0])
			cmd.Help()
		}
	},
}

func NewConfigCmd() *cobra.Command {
	return configCmd
}