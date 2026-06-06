package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Manage authentication providers",
	Long: `Manage authentication for AI providers.

Examples:
  agy auth login gemini    # Authenticate with Gemini
  agy auth list            # List configured providers
  agy auth logout gemini   # Remove provider authentication
  agy auth default gemini  # Set default provider
  agy auth connect         # TUI for remote auth`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			return
		}
		
		switch args[0] {
		case "login":
			if len(args) < 2 {
				fmt.Println("Error: provider required")
				return
			}
			// TODO: Implement login
			fmt.Printf("Authenticating with %s...\n", args[1])
		case "list":
			// TODO: List providers
			fmt.Println("Configured providers:")
		case "logout":
			if len(args) < 2 {
				fmt.Println("Error: provider required")
				return
			}
			// TODO: Implement logout
			fmt.Printf("Logging out from %s...\n", args[1])
		case "default":
			if len(args) < 2 {
				fmt.Println("Error: provider required")
				return
			}
			// TODO: Set default
			fmt.Printf("Setting default provider to %s...\n", args[1])
		case "connect":
			// TODO: TUI for remote auth
			fmt.Println("Starting remote auth TUI...")
		default:
			fmt.Printf("Unknown command: %s\n", args[0])
			cmd.Help()
		}
	},
}

func NewAuthCmd() *cobra.Command {
	return authCmd
}