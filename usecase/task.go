package usecase

import (
	"context"
	"time"

	"todo-app-go/domain"
)

type taskUsecase struct {
	taskRepository domain.TaskRepository
	contextTimeout time.Duration
}

func NewTaskUsecase(taskRepository domain.TaskRepository, timeout time.Duration) domain.TaskUsecase {
	return &taskUsecase{
		taskRepository: taskRepository,
		contextTimeout: timeout,
	}
}

func (t *taskUsecase) Create(c context.Context, task *domain.Task) error {
	ctx, cancel := context.WithTimeout(c, t.contextTimeout)
	defer cancel()
	return t.taskRepository.Create(ctx, task)
}

func (t *taskUsecase) FetchAllTasksByUserID(c context.Context, userID string) ([]domain.Task, error) {
	ctx, cancel := context.WithTimeout(c, t.contextTimeout)
	defer cancel()
	return t.taskRepository.FetchAllTasksByUserID(ctx, userID)
}

func (t *taskUsecase) FetchTasksByDateRangeAndUserID(c context.Context, userID string, startDate string, endDate string) ([]domain.Task, error) {
	ctx, cancel := context.WithTimeout(c, t.contextTimeout)
	defer cancel()
	return t.taskRepository.FetchTasksByDateRangeAndUserID(ctx, userID, startDate, endDate)
}

func (t *taskUsecase) FetchTasksByStatusAndUserID(c context.Context, userID string, status string) ([]domain.Task, error) {
	ctx, cancel := context.WithTimeout(c, t.contextTimeout)
	defer cancel()
	return t.taskRepository.FetchTasksByStatusAndUserID(ctx, userID, status)
}

func (t *taskUsecase) FetchTaskByID(c context.Context, id string) (domain.Task, error) {
	ctx, cancel := context.WithTimeout(c, t.contextTimeout)
	defer cancel()
	return t.taskRepository.FetchTaskByID(ctx, id)
}

func (t *taskUsecase) UpdateTaskByID(c context.Context, task *domain.Task) error {
	ctx, cancel := context.WithTimeout(c, t.contextTimeout)
	defer cancel()
	return t.taskRepository.UpdateTaskByID(ctx, task)
}

func (t *taskUsecase) UpdateTaskStatusByID(c context.Context, status string, id string) error {
	ctx, cancel := context.WithTimeout(c, t.contextTimeout)
	defer cancel()
	return t.taskRepository.UpdateTaskStatusByID(ctx, status, id)
}

func (t *taskUsecase) DeleteTaskByID(c context.Context, id string) error {
	ctx, cancel := context.WithTimeout(c, t.contextTimeout)
	defer cancel()
	return t.taskRepository.DeleteTaskByID(ctx, id)
}
