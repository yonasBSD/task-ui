package model

import (
	"github.com/go-task/task/v3/taskfile"
)

// TaskInfo holds information available about a task, without
// coupling to taskfile structures too much. We need a subset
// of data available for the UI, we can still use the rest
// internally.
//
// This is also the point where exposing vars and cmds becomes
// a security issue. We still asume that users may want to
// run it in less trusted environments, in which case we
// explicitly don't want any information about the contents
// of the taskfile beyond the target names and descriptions.
type TaskInfo struct {
	Task        string // Task holds the task name, e.g. `lint`.
	Description string // Description holds the task description.

	// Flags holds some flags
	Flags TaskFlags

	// Stats holds some counters
	Stats TaskStats

	// History holds history data
	History []HistoryRecord
}

// TaskStats holds counters for sensitive data.
type TaskStats struct {
	Cmds int
	Vars int
}

// TaskFlags holds some task flags.
type TaskFlags struct {
	Internal    bool
	Interactive bool
}

// NewTaskInfo converts the taskfile record into our own.
func NewTaskInfo(spec *taskfile.Task) *TaskInfo {
	return &TaskInfo{
		Task:        spec.Task,
		Description: spec.Desc,
		Flags: TaskFlags{
			Internal:    spec.Internal,
			Interactive: spec.Interactive,
		},
		Stats: TaskStats{
			Cmds: len(spec.Cmds),
			Vars: spec.Vars.Len(),
		},
	}
}
