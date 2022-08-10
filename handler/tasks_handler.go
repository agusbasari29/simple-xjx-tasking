package handler

import (
	"net/http"
	"strconv"

	"github.com/agusbasari29/simple-xjx-tasking.git/entity"
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

func (h *tasksHandler) CreateTask(ctx *gin.Context) {
	var req request.TaskRequest
	err := ctx.ShouldBind(&req)
	if err != nil {
		errFormat := helper.ErrorFormatter(err)
		errMessage := helper.M{"data_type": errFormat}
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", errMessage, nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	err = validate.Struct(req)
	if err != nil {
		errFormat := helper.ErrorFormatter(err)
		errMessage := helper.M{"validation": errFormat}
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", errMessage, nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	new, err := h.taskService.CreateTask(req)
	if err != nil {
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", "Fialed to create new task", nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	data := response.TaskFormatter(new)
	response := helper.ResponseFormatter(http.StatusOK, "success", "Task created successfully.", data)
	ctx.JSON(http.StatusOK, response)
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
	var data []response.TasksResponse
	for _, get := range gets {
		data = append(data, response.TaskFormatter(get))
	}
	response := helper.ResponseFormatter(http.StatusOK, "success", "Tasks data is successfuly retrieved", data)
	ctx.JSON(http.StatusOK, response)
}

func (h *tasksHandler) GetTaskById(ctx *gin.Context) {
	var req request.TaskIdRequest
	id := ctx.Param("id")
	if id == "" {
		err := ctx.ShouldBind(&req)
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
	} else {
		x, _ := strconv.Atoi(id)
		req.ID = uint(x)
	}
	task, err := h.taskService.GetTaskById(req)
	if err != nil {
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", "Failed to retrieved task by id", nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ResponseFormatter(http.StatusOK, "success", "Task successfully retrieved by id", task)
	ctx.JSON(http.StatusOK, response)
}

func (h *tasksHandler) GetTasksByAssignee(ctx *gin.Context) {
	var req request.TaskAssigneeRequest
	assignee := ctx.Param("assignee")
	if assignee == "" {
		err := ctx.ShouldBind(&req)
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
	} else {
		req.Assignee = assignee
	}
	results, err := h.taskService.GetTasksByAssignee(req)
	if err != nil {
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", "Failed to retrieved tasks by assignee", nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	var tasks []response.TasksResponse
	for _, result := range results {
		tasks = append(tasks, response.TaskFormatter(result))
	}
	response := helper.ResponseFormatter(http.StatusOK, "success", "Tasks successfully retrieved by assignee.", tasks)
	ctx.JSON(http.StatusOK, response)
}

func (h *tasksHandler) GetTasksByStatus(ctx *gin.Context) {
	var req request.TaskStatusRequest
	status := ctx.Param("status")
	if status == "" {
		err := ctx.ShouldBind(&req)
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
	} else {
		switch status {
		case "idle":
			{
				req.Status = entity.Idle
				break
			}
		case "progress":
			{
				req.Status = entity.Progress
				break
			}
		case "completed":
			{
				req.Status = entity.Completed
				break
			}
		}
	}
	results, err := h.taskService.GetTasksByStatus(req)
	if err != nil {
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", "Failed to retrieved tasks by status", nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	var tasks []response.TasksResponse
	for _, result := range results {
		tasks = append(tasks, response.TaskFormatter(result))
	}
	response := helper.ResponseFormatter(http.StatusOK, "success", "Tasks successfully retrieved by status.", tasks)
	ctx.JSON(http.StatusOK, response)
}

func (h *tasksHandler) UpdateTask(ctx *gin.Context) {
	var req request.TaskUpdateRequest
	err := ctx.ShouldBind(&req)
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
	update, err := h.taskService.UpdateTask(req)
	if err != nil {
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", "Failed to update task.", nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ResponseFormatter(http.StatusOK, "success", "Task was successfully updated.", update)
	ctx.JSON(http.StatusOK, response)
}

func (h *tasksHandler) DeleteTask(ctx *gin.Context) {
	var req request.TaskIdRequest
	id := ctx.Param("id")
	if id == "" {
		err := ctx.ShouldBind(&req)
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
