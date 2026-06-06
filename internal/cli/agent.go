package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var agentCmd = &cobra.Command{
	Use:   "agent",
	Short: "Run agentic coding tasks",
	Long: `Run agentic coding tasks with the configured AI provider.

Examples:
  agy agent                      # Start agent mode
  agy agent --task "fix bug"     # Run specific task
  agy agent --workspace ./src    # Use specific workspace`,
	Run: func(cmd *cobra.Command, args []string) {
		task, _ := cmd.Flags().GetString("task")
		workspace, _ := cmd.Flags().GetString("workspace")
		
		fmt.Printf("Starting agent with task: %s, workspace: %s\n", task, workspace)
		// TODO: Implement agent TUI
	},
}

func init() {
	agentCmd.Flags().StringP("task", "t", "", "Task description")
	agentCmd.Flags().StringP("workspace", "w", ".", "Workspace directory")
}

func NewAgentCmd() *cobra.Command {
	return agentCmd
}