package agent

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// JobStatus represents the status of a job
type JobStatus int

const (
	JobStatusPending JobStatus = iota
	JobStatusRunning
	JobStatusCompleted
	JobStatusFailed
	JobStatusCancelled
)

// Job represents a background job
type Job struct {
	ID          string
	Task        string
	Status      JobStatus
	CreatedAt   time.Time
	StartedAt   *time.Time
	CompletedAt *time.Time
	Error       error
	Result      string
	Workspace   string
}

// JobManager manages background jobs
type JobManager struct {
	jobs      map[string]*Job
	queue     []string
	current   string
	mu        sync.RWMutex
	ctx       context.Context
	cancel    context.CancelFunc
	wg        sync.WaitGroup
}

// NewJobManager creates a new job manager
func NewJobManager(ctx context.Context) *JobManager {
	ctx, cancel := context.WithCancel(ctx)
	return &JobManager{
		jobs:   make(map[string]*Job),
		queue:  make([]string, 0),
		ctx:    ctx,
		cancel: cancel,
	}
}

// CreateJob creates a new job
func (jm *JobManager) CreateJob(task, workspace string) *Job {
	jm.mu.Lock()
	defer jm.mu.Unlock()
	
	job := &Job{
		ID:        fmt.Sprintf("job-%d", len(jm.jobs)+1),
		Task:      task,
		Status:    JobStatusPending,
		CreatedAt: time.Now(),
		Workspace: workspace,
	}
	
	jm.jobs[job.ID] = job
	jm.queue = append(jm.queue, job.ID)
	
	return job
}

// StartJob starts a job
func (jm *JobManager) StartJob(jobID string) error {
	jm.mu.Lock()
	defer jm.mu.Unlock()
	
	job, exists := jm.jobs[jobID]
	if !exists {
		return fmt.Errorf("job not found: %s", jobID)
	}
	
	if job.Status != JobStatusPending {
		return fmt.Errorf("job not in pending state: %s", jobID)
	}
	
	now := time.Now()
	job.StartedAt = &now
	job.Status = JobStatusRunning
	jm.current = jobID
	
	// TODO: Start actual job execution in goroutine
	jm.wg.Add(1)
	go func() {
		defer jm.wg.Done()
		// TODO: Execute job
		job.Result = fmt.Sprintf("Completed: %s", job.Task)
		job.Status = JobStatusCompleted
		completedAt := time.Now()
		job.CompletedAt = &completedAt
	}()
	
	return nil
}

// CancelJob cancels a job
func (jm *JobManager) Cancel(jobID string) error {
	jm.mu.Lock()
	defer jm.mu.Unlock()
	
	job, exists := jm.jobs[jobID]
	if !exists {
		return fmt.Errorf("job not found: %s", jobID)
	}
	
	if job.Status != JobStatusRunning && job.Status != JobStatusPending {
		return fmt.Errorf("job not cancellable: %s", jobID)
	}
	
	job.Status = JobStatusCancelled
	if job.StartedAt != nil {
		completedAt := time.Now()
		job.CompletedAt = &completedAt
	}
	
	if jm.current == jobID {
		jm.current = ""
	}
	
	return nil
}

// GetJob returns a job by ID
func (jm *JobManager) GetJob(jobID string) (*Job, bool) {
	jm.mu.RLock()
	defer jm.mu.RUnlock()
	
	job, exists := jm.jobs[jobID]
	return job, exists
}

// ListJobs returns all jobs
func (jm *JobManager) ListJobs() []*Job {
	jm.mu.RLock()
	defer jm.mu.RUnlock()
	
	jobs := make([]*Job, 0, len(jm.jobs))
	for _, job := range jm.jobs {
		jobs = append(jobs, job)
	}
	return jobs
}

// Stop stops the job manager
func (jm *JobManager) Stop() {
	jm.cancel()
	jm.wg.Wait()
}