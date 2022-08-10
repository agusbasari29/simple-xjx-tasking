package request

import "github.com/agusbasari29/simple-xjx-tasking.git/entity"

type TaskRequest struct {
	Task     string `json:"task"`
	Assignee string `json:"assignee"`
	Deadline string `json:"deadline"`
}

type TaskAssigneeRequest struct {
	Assignee string `json:"assignee"`
}

type TaskStatusRequest struct {
	Status entity.TaskStatus `json:"status"`
}

type TaskIdRequest struct {
	ID uint `json:"id"`
}

type TaskUpdateRequest struct {
	ID       uint              `json:"id"`
	Task     string            `json:"task"`
	Assignee string            `json:"assignee"`
	Deadline string            `json:"deadline"`
	Status   entity.TaskStatus `json:"status"`
}
