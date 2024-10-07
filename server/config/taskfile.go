package config

import (
	"time"

	"github.com/go-task/task/v3/taskfile"
	"github.com/go-task/task/v3/taskfile/ast"
)

func NewTaskfileLoader() *TaskfileLoader {
	return &TaskfileLoader{}
}

type TaskfileLoader struct{}

func (*TaskfileLoader) Load(folder, filename string) (*ast.Taskfile, error) {
	node, err := taskfile.NewRootNode(nil, filename, folder, false, 60)
	if err != nil {
		return nil, err
	}

	reader := taskfile.NewReader(node, false, false, false, time.Duration(0), "/tmp", nil)
	graph, err := reader.Read()

	if err != nil {
		return nil, err
	}

	t, err := graph.Merge()

	if err != nil {
		return nil, err
	}

	return t, err
}
