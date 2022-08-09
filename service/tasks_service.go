package service

import (
	"log"
	"time"

	"github.com/agusbasari29/simple-xjx-tasking.git/entity"
	"github.com/agusbasari29/simple-xjx-tasking.git/repository"
	"github.com/agusbasari29/simple-xjx-tasking.git/request"
	"github.com/mashingan/smapping"
)

type TasksService interface {
	CreateTask(req request.TaskRequest) (entity.Tasks, error)
	GetTasks() ([]entity.Tasks, error)
	GetTaskById(req request.TaskIdRequest) (entity.Tasks, error)
	GetTasksByAssignee(req request.TaskRequest) ([]entity.Tasks, error)
	GetTasksByStatus(req request.TaskRequest) ([]entity.Tasks, error)
	UpdateTask(req request.TaskUpdateRequest) (entity.Tasks, error)
	DeleteTask(req request.TaskIdRequest) (entity.Tasks, error)
}

type tasksServices struct {
	tasksRepository repository.TasksRepository
}

func NewTasksServices(tasksRepository repository.TasksRepository) *tasksServices {
	return &tasksServices{tasksRepository}
}

func (s *tasksServices) CreateTask(req request.TaskRequest) (entity.Tasks, error) {
	task := entity.Tasks{}
	err := smapping.FillStruct(&task, smapping.MapFields(&req))
	if err != nil {
		log.Fatalf("Failed to fill task: %v", err)
	}
	task.CreatedAt = time.Now()
	new, err := s.tasksRepository.CreateTask(task)
	if err != nil {
		return task, err
	}
	return new, nil
}

func (s *tasksServices) GetTasks() ([]entity.Tasks, error) {
	var tasks []entity.Tasks
	result, err := s.tasksRepository.GetTasks()
	if err != nil {
		return tasks, err
	}
	return result, nil
}

func (s *tasksServices) GetTaskById(req request.TaskIdRequest) (entity.Tasks, error) {
	task := entity.Tasks{}
	err := smapping.FillStruct(&task, smapping.MapFields(&req))
	if err != nil {
		log.Fatalf("Failed to fill task: %v", err)
	}
	get, err := s.tasksRepository.GetTaskById(task)
	if err != nil {
		return task, err
	}
	return get, nil
}

func (s *tasksServices) GetTasksByAssignee(req request.TaskRequest) ([]entity.Tasks, error) {
	task := entity.Tasks{}
	err := smapping.FillStruct(&task, smapping.MapFields(&req))
	if err != nil {
		log.Fatalf("Failed to fill task: %v", err)
	}
	gets, err := s.tasksRepository.GetTasksByAssignee(task)
	if err != nil {
		return gets, err
	}
	return gets, nil
}

func (s *tasksServices) GetTasksByStatus(req request.TaskRequest) ([]entity.Tasks, error) {
	task := entity.Tasks{}
	err := smapping.FillStruct(&task, smapping.MapFields(&req))
	if err != nil {
		log.Fatalf("Failed to fill task: %v", err)
	}
	gets, err := s.tasksRepository.GetTasksByStatus(task)
	if err != nil {
		return gets, err
	}
	return gets, nil
}

func (s *tasksServices) UpdateTask(req request.TaskUpdateRequest) (entity.Tasks, error) {
	task := entity.Tasks{}
	err := smapping.FillStruct(&task, smapping.MapFields(&req))
	if err != nil {
		log.Fatalf("Failed to fill task: %v", err)
	}
	update, err := s.tasksRepository.UpdateTask(task)
	if err != nil {
		return task, err
	}
	return update, nil
}

func (s *tasksServices) DeleteTask(req request.TaskIdRequest) (entity.Tasks, error) {
	task := entity.Tasks{}
	err := smapping.FillStruct(&task, smapping.MapFields(&req))
	if err != nil {
		log.Fatalf("Failed to fill task: %v", err)
	}
	del, err := s.tasksRepository.DeleteTask(task)
	if err != nil {
		return task, err
	}
	return del, nil
}
