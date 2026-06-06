package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "agy",
	Short: "Antigravity CLI-style agentic coding tool for Android/Termux",
	Long: `and-agy is a Termux-native Go application that provides an Antigravity-style
agentic coding experience on Android devices. It features a Bubble Tea TUI,
multi-provider AI support, and secure authentication.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(
		NewChatCmd(),
		NewAgentCmd(),
		NewConfigCmd(),
		NewAuthCmd(),
		NewVersionCmd(),
	)
}

func Execute() error {
	return rootCmd.Execute()
}

// NewVersionCmd returns the version command
func NewVersionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print the version information",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("and-agy version dev")
		},
	}
}