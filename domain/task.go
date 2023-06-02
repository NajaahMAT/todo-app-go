package domain

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionTask = "tasks"
)

type Task struct {
	ID          primitive.ObjectID `bson:"_id" json:"id"`
	Title       string             `bson:"title" json:"title"`
	Status      string             `bson:"status" json:"status"`
	Description string             `bson:"description" json:"description"`
	Files       string             `bson:"files" json:"files"`
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
	UserID      primitive.ObjectID `bson:"user_id" json:"user_id"`
}

type TaskRepository interface {
	Create(c context.Context, task *Task) error
	FetchAllTasksByUserID(c context.Context, userID string) ([]Task, error)
	FetchTasksByDateRangeAndUserID(c context.Context, userID string, startDate time.Time, endDate time.Time) ([]Task, error)
	FetchTasksByStatusAndUserID(c context.Context, userID string, status string) ([]Task, error)
	FetchTaskByID(c context.Context, id string) (Task, error)
	UpdateTaskByID(c context.Context, task *Task) error
	UpdateTaskStatusByID(c context.Context, status string, id string) error
	DeleteTaskByID(c context.Context, id string) error
}

type TaskUsecase interface {
	Create(c context.Context, task *Task) error
	FetchAllTasksByUserID(c context.Context, userID string) ([]Task, error)
	FetchTasksByDateRangeAndUserID(c context.Context, userID string, startDate time.Time, endDate time.Time) ([]Task, error)
	FetchTasksByStatusAndUserID(c context.Context, userID string, status string) ([]Task, error)
	FetchTaskByID(c context.Context, id string) (Task, error)
	UpdateTaskByID(c context.Context, task *Task) error
	UpdateTaskStatusByID(c context.Context, status string, id string) error
	DeleteTaskByID(c context.Context, id string) error
}
