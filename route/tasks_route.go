package route

import (
	"github.com/agusbasari29/simple-xjx-tasking.git/database"
	"github.com/agusbasari29/simple-xjx-tasking.git/handler"
	"github.com/agusbasari29/simple-xjx-tasking.git/helper"
	"github.com/agusbasari29/simple-xjx-tasking.git/repository"
	"github.com/agusbasari29/simple-xjx-tasking.git/service"
	"github.com/gin-gonic/gin"
)

type TaskRoutes struct{}

func (r TaskRoutes) Route() []helper.Route {
	db := database.SetupDatabaseConnection()
	repo := repository.NewTasksRepository(db)
	serv := service.NewTasksServices(repo)
	hand := handler.NewTasksHandler(serv)

	return []helper.Route{
		{
			Path:    "/tasks",
			Method:  "GET",
			Handler: []gin.HandlerFunc{hand.GetTasks},
		},
		{
			Path:    "/tasks",
			Method:  "POST",
			Handler: []gin.HandlerFunc{hand.CreateTask},
		},
		{
			Path:    "/tasks",
			Method:  "POST",
			Handler: []gin.HandlerFunc{hand.GetTaskById},
		},
		{
			Path:    "/tasks/id/:id",
			Method:  "GET",
			Handler: []gin.HandlerFunc{hand.GetTaskById},
		},
		{
			Path:    "/tasks/assignee",
			Method:  "POST",
			Handler: []gin.HandlerFunc{hand.GetTasksByAssignee},
		},
		{
			Path:    "/tasks/assignee/:assignee",
			Method:  "GET",
			Handler: []gin.HandlerFunc{hand.GetTasksByAssignee},
		},
		{
			Path:    "/tasks/status",
			Method:  "POST",
			Handler: []gin.HandlerFunc{hand.GetTasksByStatus},
		},
		{
			Path:    "/tasks/status/:status",
			Method:  "GET",
			Handler: []gin.HandlerFunc{hand.GetTasksByStatus},
		},
		{
			Path:    "/tasks",
			Method:  "PUT",
			Handler: []gin.HandlerFunc{hand.UpdateTask},
		},
		{
			Path:    "/tasks",
			Method:  "DELETE",
			Handler: []gin.HandlerFunc{hand.DeleteTask},
		},
		{
			Path:    "/tasks/delete/:id",
			Method:  "GET",
			Handler: []gin.HandlerFunc{hand.DeleteTask},
		},
	}
}
