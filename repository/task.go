package repository

import (
	"context"
	"time"
	"todo-app-go/domain"
	"todo-app-go/mongo"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type taskRepository struct {
	database   mongo.Database
	collection string
}

func NewTaskRepository(db mongo.Database, collection string) domain.TaskRepository {
	return &taskRepository{
		database:   db,
		collection: collection,
	}
}

func (tr *taskRepository) Create(c context.Context, task *domain.Task) error {
	collection := tr.database.Collection(tr.collection)

	_, err := collection.InsertOne(c, task)

	return err
}

func (tr *taskRepository) FetchAllTasksByUserID(c context.Context, userID string) ([]domain.Task, error) {
	var response []domain.Task

	return response, nil
}

func (tr *taskRepository) FetchTasksByDateRangeAndUserID(c context.Context, userID string, startDate time.Time, endDate time.Time) ([]domain.Task, error) {
	var response []domain.Task

	return response, nil
}

func (tr *taskRepository) FetchTasksByStatusAndUserID(c context.Context, userID string, status string) ([]domain.Task, error) {
	var response []domain.Task

	return response, nil
}

func (tr *taskRepository) FetchTaskByID(c context.Context, id primitive.ObjectID) (domain.Task, error) {
	var response domain.Task

	return response, nil
}

func (tr *taskRepository) UpdateTaskByID(c context.Context, task *domain.Task) error {
	return nil

}

func (tr *taskRepository) UpdateTaskStatusByID(c context.Context, status string, id primitive.ObjectID) error {
	return nil
}

func (tr *taskRepository) DeleteTaskByID(c context.Context, id primitive.ObjectID) error {
	return nil
}
