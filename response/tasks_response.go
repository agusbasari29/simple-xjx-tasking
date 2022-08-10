package response

import (
	"time"

	"github.com/agusbasari29/simple-xjx-tasking.git/entity"
	"gorm.io/gorm"
)

type TasksResponse struct {
	ID        uint              `json:"id"`
	Task      string            `json:"task"`
	Assignee  string            `json:"assignee"`
	Status    entity.TaskStatus `json:"status"`
	Deadline  time.Time         `json:"deadline"`
	CreatedAt time.Time         `json:"created_at"`
	UpdatedAt time.Time         `json:"updated_at"`
	DeletedAt gorm.DeletedAt    `json:"deleted_at"`
}

func TaskFormatter(tasks entity.Tasks) TasksResponse {
	taskFormatter := TasksResponse{}
	taskFormatter.ID = tasks.ID
	taskFormatter.Task = tasks.Task
	taskFormatter.Assignee = tasks.Assignee
	taskFormatter.Status = tasks.Status
	taskFormatter.Deadline = tasks.Deadline
	taskFormatter.CreatedAt = tasks.CreatedAt
	taskFormatter.UpdatedAt = tasks.UpdatedAt
	taskFormatter.DeletedAt = tasks.DeletedAt
	return taskFormatter
}
