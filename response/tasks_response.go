package response

import (
	"time"

	"github.com/agusbasari29/simple-xjx-tasking.git/entity"
)

type TasksResponse struct {
	ID       uint              `json:"id"`
	Task     string            `json:"task"`
	Assignee string            `json:"assignee"`
	Status   entity.TaskStatus `json:"status"`
	Deadline time.Time         `json:"deadline"`
}

func TaskFormatter(tasks entity.Tasks) TasksResponse {
	taskFormatter := TasksResponse{}
	taskFormatter.ID = tasks.ID
	taskFormatter.Task = tasks.Task
	taskFormatter.Assignee = tasks.Assignee
	taskFormatter.Status = tasks.Status
	taskFormatter.Deadline = tasks.Deadline
	return taskFormatter
}
