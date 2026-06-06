package agent

import (
	"context"
	"fmt"
)

// Harness manages the agent lifecycle
type Harness struct {
	planner     *Planner
	jobManager  *JobManager
	currentJob  *Job
	workspace   string
}

// NewHarness creates a new agent harness
func NewHarness(planner *Planner, jobManager *JobManager, workspace string) *Harness {
	return &Harness{
		planner:    planner,
		jobManager: jobManager,
		workspace:  workspace,
	}
}

// Start begins agent execution
func (h *Harness) Start(ctx context.Context, task string) error {
	fmt.Printf("Starting agent with task: %s\n", task)
	
	// TODO: Implement task planning
	plan, err := h.planner.Plan(task)
	if err != nil {
		return fmt.Errorf("planning failed: %w", err)
	}
	
	// TODO: Execute plan
	for _, step := range plan.Steps {
		fmt.Printf("Executing step: %s\n", step.Description)
		// TODO: Execute step
	}
	
	return nil
}

// Stop stops the current job
func (h *Harness) Stop() error {
	if h.currentJob != nil {
		return h.jobManager.Cancel(h.currentJob.ID)
	}
	return nil
}

// SetWorkspace changes the workspace
func (h *Harness) SetWorkspace(workspace string) {
	h.workspace = workspace
}