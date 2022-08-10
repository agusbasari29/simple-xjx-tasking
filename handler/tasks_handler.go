package handler

import (
	"net/http"
	"strconv"

	"github.com/agusbasari29/simple-xjx-tasking.git/helper"
	"github.com/agusbasari29/simple-xjx-tasking.git/request"
	"github.com/agusbasari29/simple-xjx-tasking.git/response"
	"github.com/agusbasari29/simple-xjx-tasking.git/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type TasksHandler interface {
}

type tasksHandler struct {
	taskService service.TasksService
}

func NewTasksHandler(taskService service.TasksService) *tasksHandler {
	return &tasksHandler{taskService}
}

func (h *tasksHandler) CreateTask(ctx *gin.Context) {}

func (h *tasksHandler) GetTasks(ctx *gin.Context) {
	gets, err := h.taskService.GetTasks()
	if err != nil {
		errFormat := helper.ErrorFormatter(err)
		errMessage := helper.M{"error": errFormat}
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", errMessage, nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	var data []response.TasksResponse
	for _, get := range gets {
		data = append(data, response.TaskFormatter(get))
	}
	response := helper.ResponseFormatter(http.StatusOK, "success", "Tasks data is successfuly retrieved", data)
	ctx.JSON(http.StatusOK, response)
}

func (h *tasksHandler) GetTaskById(ctx *gin.Context) {}

func (h *tasksHandler) GetTasksByAssignee(ctx *gin.Context) {}

func (h *tasksHandler) GetTasksByStatus(ctx *gin.Context) {}

func (h *tasksHandler) UpdateTask(ctx *gin.Context) {}

func (h *tasksHandler) DeleteTask(ctx *gin.Context) {
	var req request.TaskIdRequest
	id := ctx.Param("id")
	if id == "" {
		err := ctx.ShouldBind(req)
		if err != nil {
			errFormat := helper.ErrorFormatter(err)
			errMessage := helper.M{"error": errFormat}
			response := helper.ResponseFormatter(http.StatusBadRequest, "data_type", errMessage, nil)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}
		err = validate.Struct(req)
		if err != nil {
			errFormat := helper.ErrorFormatter(err)
			errMessage := helper.M{"error": errFormat}
			response := helper.ResponseFormatter(http.StatusBadRequest, "validation", errMessage, nil)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", "Parameter can not be empty!", nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	} else {
		x, _ := strconv.Atoi(id)
		req.ID = uint(x)
	}
	del, err := h.taskService.DeleteTask(req)
	if err != nil {
		task := response.TaskFormatter(del)
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", "Failed to delete task.", task)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	} else {
		task := response.TaskFormatter(del)
		response := helper.ResponseFormatter(http.StatusOK, "success", "Task is  successfully deleted.", task)
		ctx.JSON(http.StatusOK, response)
	}
}
