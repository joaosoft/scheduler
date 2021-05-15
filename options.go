package scheduler

import (
	logger "github.com/joaosoft/logger"
	"github.com/joaosoft/manager"
)

// SchedulerOption ...
type SchedulerOption func(s *Scheduler)

// Reconfigure ...
func (s *Scheduler) Reconfigure(options ...SchedulerOption) {
	for _, option := range options {
		option(s)
	}
}

// WithConfiguration ...
func WithConfiguration(config *SchedulerConfig) SchedulerOption {
	return func(s *Scheduler) {
		s.config = config
	}
}

// WithLogger ...
func WithLogger(logger logger.ILogger) SchedulerOption {
	return func(s *Scheduler) {
		s.logger = logger
		s.isLogExternal = true
	}
}

// WithLogLevel ...
func WithLogLevel(level logger.Level) SchedulerOption {
	return func(s *Scheduler) {
		s.logger.SetLevel(level)
	}
}

// WithManager ...
func WithManager(mgr *manager.Manager) SchedulerOption {
	return func(s *Scheduler) {
		s.pm = mgr
	}
}
