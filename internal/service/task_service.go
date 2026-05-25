package service

import (
	"context"
	"demo09/internal/model"
	"demo09/internal/repository"
)

type TaskService interface {
	GetAll(ctx context.Context) ([]model.Task, error)
	GetById(ctx context.Context, id string) (model.Task, error)
	Create(ctx context.Context, task model.Task) (model.Task, error)
}

type taskService struct {
	taskRepository repository.TaskRepository
}

func NewTaskService(taskRepository repository.TaskRepository) TaskService {
	return &taskService{
		taskRepository: taskRepository,
	}
}

func (s *taskService) GetAll(ctx context.Context) ([]model.Task, error) {
	return s.taskRepository.GetAll(ctx)
}

func (s *taskService) GetById(ctx context.Context, id string) (model.Task, error) {
	return s.taskRepository.GetById(ctx, id)
}

func (s *taskService) Create(ctx context.Context, task model.Task) (model.Task, error) {
	return s.taskRepository.Create(ctx, task)
}
