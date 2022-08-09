package repository

import (
	"github.com/agusbasari29/simple-xjx-tasking.git/entity"
	"gorm.io/gorm"
)

type TasksRepository interface {
	CreateTask(task entity.Tasks) (entity.Tasks, error)
	GetTasks() ([]entity.Tasks, error)
	GetTaskById(task entity.Tasks) (entity.Tasks, error)
	GetTasksByAssignee(task entity.Tasks) ([]entity.Tasks, error)
	GetTasksByStatus(task entity.Tasks) ([]entity.Tasks, error)
	UpdateTask(task entity.Tasks) (entity.Tasks, error)
	DeleteTask(task entity.Tasks) (entity.Tasks, error)
}

type tasksRepository struct {
	db *gorm.DB
}

func NewTasksRepository(db *gorm.DB) *tasksRepository {
	return &tasksRepository{db}
}

func (r *tasksRepository) CreateTask(task entity.Tasks) (entity.Tasks, error) {
	err := r.db.Raw("INSERT INTO tasks (id, task, assignee, deadline, status) VALUES (@ID, @Tasks, @Assignee, @Deadline, @Status)", task).Save(&task).Error
	if err != nil {
		return task, err
	}
	return task, nil
}

func (r *tasksRepository) GetTasks() ([]entity.Tasks, error) {
	tasks := []entity.Tasks{}
	err := r.db.Raw("SELECT * FROM tasks").Save(&tasks).Error
	if err != nil {
		return tasks, err
	}
	return tasks, nil
}

func (r *tasksRepository) GetTaskById(task entity.Tasks) (entity.Tasks, error) {
	err := r.db.Raw("SELECT * FROM tasks WHERE id=@ID", task).Save(&task).Error
	if err != nil {
		return task, err
	}
	return task, nil
}

func (r *tasksRepository) GetTasksByAssignee(task entity.Tasks) ([]entity.Tasks, error) {
	tasks := []entity.Tasks{}
	err := r.db.Raw("SELECT * FROM tasks WHERE assignee=@Assignee", task).Save(&tasks).Error
	if err != nil {
		return tasks, err
	}
	return tasks, nil
}

func (r *tasksRepository) GetTasksByStatus(task entity.Tasks) ([]entity.Tasks, error) {
	tasks := []entity.Tasks{}
	err := r.db.Raw("SELECT * FROM tasks WHERE status=@Status", task).Save(&tasks).Error
	if err != nil {
		return tasks, err
	}
	return tasks, nil
}

func (r *tasksRepository) UpdateTask(task entity.Tasks) (entity.Tasks, error) {
	err := r.db.Raw("UPDATE tasks SET task = @Task, assignee = @Assignee, deadline = @Deadline, status = @Status WHERE id = @ID", task).Save(&task).Error
	if err != nil {
		return task, err
	}
	return task, nil
}

func (r *tasksRepository) DeleteTask(task entity.Tasks) (entity.Tasks, error) {
	err := r.db.Raw("DELETE FROM tasks WHERE id=@ID", task).Save(&task).Error
	if err != nil {
		return task, err
	}
	return task, nil
}
