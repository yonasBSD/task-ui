package repository

import (
	"fmt"

	"github.com/go-task/task/v3/taskfile/ast"

	. "github.com/titpetric/task-ui/server/model"
)

// FilterFunc is a signature for task filters.
type FilterFunc func(task *ast.Task) bool

// List tasks lists the names of the defined tasks.
func ListTasks(spec *ast.Taskfile, filters ...FilterFunc) []*TaskInfo {
	result := make([]*TaskInfo, 0, spec.Tasks.Len())
	for _, task := range spec.Tasks.Values() {
		for _, filter := range filters {
			if filter(task) {
				goto next
			}
		}

		result = append(result, NewTaskInfo(task))

	next:
	}
	return result
}

// FilterOutNoDesc removes all tasks that do not contain a description.
func FilterOutNoDesc(task *ast.Task) bool {
	return task.Desc == ""
}

// FilterOutInternal removes all tasks that are marked as internal.
func FilterOutInternal(task *ast.Task) bool {
	return task.Internal
}

// FindTask returns the ast.Task for a given name.
func FindTask(spec *ast.Taskfile, name string) (*ast.Task, error) {
	for _, task := range spec.Tasks.Values() {
		if task.Task == name {
			return task, nil
		}
	}
	return nil, fmt.Errorf("no such task: %s: %w", name, ErrNotFound)
}
