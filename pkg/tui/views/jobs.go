package views

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbletea"
)

// JobStatus represents job status
type JobStatus int

const (
	JobStatusPending JobStatus = iota
	JobStatusRunning
	JobStatusCompleted
	JobStatusFailed
	JobStatusCancelled
)

// Job represents a job in the jobs view
type Job struct {
	ID      string
	Task    string
	Status  JobStatus
	Progress int
}

// JobsView displays and manages background jobs
type JobsView struct {
	jobs    []Job
	cursor  int
	quitting bool
}

// NewJobsView creates a new jobs view
func NewJobsView() *JobsView {
	return &JobsView{
		jobs: make([]Job, 0),
	}
}

// Init initializes the jobs view
func (jv *JobsView) Init() tea.Model {
	return jv
}

// Update handles messages
func (jv *JobsView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			jv.quitting = true
			return jv, tea.Quit
		case "up":
			if jv.cursor > 0 {
				jv.cursor--
			}
		case "down":
			if jv.cursor < len(jv.jobs)-1 {
				jv.cursor++
			}
		case "c":
			// Cancel job
			if len(jv.jobs) > 0 {
				jv.jobs[jv.cursor].Status = JobStatusCancelled
			}
		case "r":
			// Refresh
			// TODO: Refresh job list
		}
	}
	
	return jv, nil
}

// View renders the jobs view
func (jv *JobsView) View() string {
	var sb strings.Builder
	
	// Header
	sb.WriteString("=== Background Jobs ===\n\n")
	
	// Jobs list
	for i, job := range jv.jobs {
		cursor := "  "
		if i == jv.cursor {
			cursor = "> "
		}
		
		status := "Pending"
		switch job.Status {
		case JobStatusRunning:
			status = "Running"
		case JobStatusCompleted:
			status = "Completed"
		case JobStatusFailed:
			status = "Failed"
		case JobStatusCancelled:
			status = "Cancelled"
		}
		
		progress := ""
		if job.Progress > 0 {
			progress = fmt.Sprintf(" [%d%%]", job.Progress)
		}
		
		sb.WriteString(fmt.Sprintf("%s%s: %s%s\n", cursor, job.ID, job.Task, progress))
		sb.WriteString(fmt.Sprintf("   Status: %s\n", status))
		sb.WriteString("\n")
	}
	
	if len(jv.jobs) == 0 {
		sb.WriteString("No jobs running\n")
	}
	
	// Help
	sb.WriteString("Press q to quit, up/down to navigate, c to cancel, r to refresh\n")
	
	return sb.String()
}

// Name returns the view name
func (jv *JobsView) Name() string {
	return "jobs"
}

// AddJob adds a job to the list
func (jv *JobsView) AddJob(job Job) {
	jv.jobs = append(jv.jobs, job)
}

// UpdateJob updates a job in the list
func (jv *JobsView) UpdateJob(id string, status JobStatus, progress int) {
	for i := range jv.jobs {
		if jv.jobs[i].ID == id {
			jv.jobs[i].Status = status
			jv.jobs[i].Progress = progress
			break
		}
	}
}

// RemoveJob removes a job from the list
func (jv *JobsView) RemoveJob(id string) {
	for i := range jv.jobs {
		if jv.jobs[i].ID == id {
			jv.jobs = append(jv.jobs[:i], jv.jobs[i+1:]...)
			if jv.cursor >= i {
				jv.cursor--
			}
			break
		}
	}
}