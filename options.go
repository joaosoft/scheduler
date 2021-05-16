package scheduler

import (
	logger "github.com/joaosoft/logger"
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
	}
}

// WithLogLevel ...
func WithLogLevel(level logger.Level) SchedulerOption {
	return func(s *Scheduler) {
		s.logger.SetLevel(level)
	}
}
