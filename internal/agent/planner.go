package agent

import (
	"fmt"
	"strings"
)

// Step represents a single step in a plan
type Step struct {
	ID          string
	Description string
	Action      string
	Parameters  map[string]string
	Completed   bool
	Error       error
}

// Plan represents a complete execution plan
type Plan struct {
	ID          string
	Task        string
	Steps       []*Step
	CurrentStep int
	Completed   bool
}

// Planner creates execution plans
type Planner struct {
	// TODO: Add configuration
}

// NewPlanner creates a new planner
func NewPlanner() *Planner {
	return &Planner{}
}

// Plan creates a plan for the given task
func (p *Planner) Plan(task string) (*Plan, error) {
	plan := &Plan{
		ID:   fmt.Sprintf("plan-%d", 1), // TODO: Generate proper ID
		Task: task,
		Steps: []*Step{
			{
				ID:          "step-1",
				Description: "Analyze task requirements",
				Action:      "analyze",
				Parameters:  map[string]string{"task": task},
			},
			{
				ID:          "step-2",
				Description: "Identify required tools",
				Action:      "identify-tools",
				Parameters:  map[string]string{"task": task},
			},
			{
				ID:          "step-3",
				Description: "Execute plan",
				Action:      "execute",
				Parameters:  map[string]string{"task": task},
			},
		},
	}
	
	// TODO: Implement actual planning logic
	// This is a placeholder that splits the task into basic steps
	if strings.Contains(task, "fix") || strings.Contains(task, "bug") {
		plan.Steps = append(plan.Steps, &Step{
			ID:          "step-4",
			Description: "Verify fix",
			Action:      "verify",
			Parameters:  map[string]string{"task": task},
		})
	}
	
	return plan, nil
}

// NextStep advances to the next step
func (p *Planner) NextStep(plan *Plan) *Step {
	if plan.CurrentStep < len(plan.Steps) {
		step := plan.Steps[plan.CurrentStep]
		plan.CurrentStep++
		return step
	}
	return nil
}

// PreviousStep goes back to the previous step
func (p *Planner) PreviousStep(plan *Plan) *Step {
	if plan.CurrentStep > 0 {
		plan.CurrentStep--
		return plan.Steps[plan.CurrentStep]
	}
	return nil
}