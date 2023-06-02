package repository

import (
	"context"
	"time"
	"todo-app-go/domain"
	"todo-app-go/mongo"

	"go.mongodb.org/mongo-driver/bson"
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
	collection := tr.database.Collection(tr.collection)

	var tasks []domain.Task

	idHex, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return tasks, err
	}

	cursor, err := collection.Find(c, bson.M{"userID": idHex})
	if err != nil {
		return nil, err
	}

	err = cursor.All(c, &tasks)
	if tasks == nil {
		return []domain.Task{}, err
	}

	return tasks, err
}

func (tr *taskRepository) FetchTasksByDateRangeAndUserID(c context.Context, userID string, startDate time.Time, endDate time.Time) ([]domain.Task, error) {
	var response []domain.Task

	return response, nil
}

func (tr *taskRepository) FetchTasksByStatusAndUserID(c context.Context, userID string, status string) ([]domain.Task, error) {
	var response []domain.Task

	return response, nil
}

func (tr *taskRepository) FetchTaskByID(c context.Context, id string) (domain.Task, error) {
	var response domain.Task

	return response, nil
}

func (tr *taskRepository) UpdateTaskByID(c context.Context, task *domain.Task) error {
	return nil

}

func (tr *taskRepository) UpdateTaskStatusByID(c context.Context, status string, id string) error {
	return nil
}

func (tr *taskRepository) DeleteTaskByID(c context.Context, id string) error {
	return nil
}
