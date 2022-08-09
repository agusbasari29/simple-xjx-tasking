package request

type TaskRequest struct {
	Task     string `json:"task"`
	Assignee string `json:"assignee"`
	Deadline string `json:"deadline"`
	Status   string `json:"status"`
}

type TaskIdRequest struct {
	ID uint `json:"id"`
}

type TaskUpdateRequest struct {
	ID       uint   `json:"id"`
	Task     string `json:"task"`
	Assignee string `json:"assignee"`
	Deadline string `json:"deadline"`
	Status   string `json:"status"`
}
