package store

import (
	"os"
	"strings"
	"time"
)

type TaskActions interface {
	AddTask(task Task) (int, error)
	GetTask(taskId int) (*Task, error)
	DeleteTask(taskId int) error
	GetTasks() ([]Task, error)
}

type Task struct {
	Id   int       `json:"id"`
	Text string    `json:"text"`
	Due  time.Time `json:"due"`
	Tags []string  `json:"tags"`
}

var TasksStore TaskActions

func NewStore() error {
	isTestEnv := os.Getenv("IS_TEST_ENV")
	var err error
	if isTestEnv == "" {
		TasksStore, err = newDBStore()
	}
	// return NewTestStore()
	return err
}

var delimiter string = ","

func stringifyTags(tags []string) string {

	return strings.Join(tags, delimiter)
}

func slicifyTags(tags string) []string {
	if len(tags) == 0 {
		return []string{}
	}

	splittedTags := strings.Split(tags, delimiter)
	return splittedTags
}
