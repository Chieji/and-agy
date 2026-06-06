package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var chatCmd = &cobra.Command{
	Use:   "chat",
	Short: "Start an interactive chat session",
	Long: `Start an interactive chat session with the configured AI provider.

Examples:
  agy chat                    # Start chat with default provider
  agy chat --provider gemini # Use specific provider
  agy chat --model gemini-1.5-pro # Use specific model`,
	Run: func(cmd *cobra.Command, args []string) {
		provider, _ := cmd.Flags().GetString("provider")
		model, _ := cmd.Flags().GetString("model")
		
		fmt.Printf("Starting chat with provider: %s, model: %s\n", provider, model)
		// TODO: Implement chat TUI
	},
}

func init() {
	chatCmd.Flags().StringP("provider", "p", "", "AI provider to use (gemini, openai, etc.)")
	chatCmd.Flags().StringP("model", "m", "", "Model to use")
}

func NewChatCmd() *cobra.Command {
	return chatCmd
}