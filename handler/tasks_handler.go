package handler

import (
	"net/http"

	"github.com/agusbasari29/simple-xjx-tasking.git/helper"
	"github.com/agusbasari29/simple-xjx-tasking.git/service"
	"github.com/gin-gonic/gin"
)

type TasksHandler interface {
}

type tasksHandler struct {
	taskService service.TasksService
}

func NewTasksHandler(taskService service.TasksService) *tasksHandler {
	return &tasksHandler{taskService}
}

func (h *tasksHandler) GetTasks(ctx *gin.Context) {
	gets, err := h.taskService.GetTasks()
	if err != nil {
		errFormat := helper.ErrorFormatter(err)
		errMessage := helper.M{"error": errFormat}
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", errMessage, nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ResponseFormatter(http.StatusOK, "success", "Successfuly retrieved data tasks", gets)
	ctx.JSON(http.StatusOK, response)
}
