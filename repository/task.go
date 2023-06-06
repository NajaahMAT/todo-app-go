package repository

import (
	"context"
	"log"
	"runtime"
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
	_, filename, line, _ := runtime.Caller(1)

	collection := tr.database.Collection(tr.collection)

	var tasks []domain.Task

	idHex, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		log.Printf("[error] %s:%d %v\n", filename, line, err)
		return tasks, err
	}

	cursor, err := collection.Find(c, bson.M{"user_id": idHex})
	if err != nil {
		log.Printf("[error] %s:%d %v\n", filename, line, err)
		return nil, err
	}

	err = cursor.All(c, &tasks)
	if tasks == nil {
		log.Printf("[error] %s:%d %v\n", filename, line, err)
		return []domain.Task{}, err
	}

	return tasks, err
}

func (tr *taskRepository) FetchTasksByDateRangeAndUserID(c context.Context, userID string, startDate string, endDate string) ([]domain.Task, error) {
	_, filename, line, _ := runtime.Caller(1)

	collection := tr.database.Collection(tr.collection)

	var tasks []domain.Task

	idHex, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		log.Printf("[error] %s:%d %v\n", filename, line, err)
		return tasks, err
	}

	start_date, _ := time.Parse(domain.DateTimeLayout, startDate)
	end_date, _ := time.Parse(domain.DateTimeLayout, endDate)

	cursor, err := collection.Find(c, bson.M{"user_id": idHex, "created_at": bson.M{
		"$gte": start_date,
		"$lt":  end_date,
	}})

	if err != nil {
		log.Printf("[error] %s:%d %v\n", filename, line, err)
		return nil, err
	}

	err = cursor.All(c, &tasks)
	if tasks == nil {
		log.Printf("[error] %s:%d %v\n", filename, line, err)
		return []domain.Task{}, err
	}

	return tasks, err
}

func (tr *taskRepository) FetchTasksByStatusAndUserID(c context.Context, userID string, status string) ([]domain.Task, error) {
	_, filename, line, _ := runtime.Caller(1)

	collection := tr.database.Collection(tr.collection)

	var tasks []domain.Task

	idHex, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		log.Printf("[error] %s:%d %v\n", filename, line, err)
		return tasks, err
	}

	cursor, err := collection.Find(c, bson.M{"user_id": idHex, "status": status})
	if err != nil {
		log.Printf("[error] %s:%d %v\n", filename, line, err)
		return nil, err
	}

	err = cursor.All(c, &tasks)
	if tasks == nil {
		log.Printf("[error] %s:%d %v\n", filename, line, err)
		return []domain.Task{}, err
	}

	return tasks, err
}

func (tr *taskRepository) FetchTaskByID(c context.Context, id string) (domain.Task, error) {
	_, filename, line, _ := runtime.Caller(1)

	collection := tr.database.Collection(tr.collection)

	var task domain.Task

	idHex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Printf("[error] %s:%d %v\n", filename, line, err)
		return task, err
	}

	cursor := collection.FindOne(c, bson.M{"_id": idHex})

	err = cursor.Decode(&task)
	if err != nil {
		log.Printf("[error] %s:%d %v\n", filename, line, err)
		return domain.Task{}, err
	}

	return task, err
}

func (tr *taskRepository) UpdateTaskByID(c context.Context, task *domain.Task) error {
	_, filename, line, _ := runtime.Caller(1)

	collection := tr.database.Collection(tr.collection)

	update := bson.M{
		"title":       task.Title,
		"status":      task.Status,
		"description": task.Description,
		"file":        task.File,
		"updated_at":  time.Now(),
	}

	result, err := collection.UpdateOne(c, bson.M{"_id": task.ID}, bson.M{"$set": update})
	if err != nil {
		log.Printf("[error] %s:%d %v\n", filename, line, err)
		return err
	}

	if result.ModifiedCount > 0 {
		log.Println("Successfully updated the task. modified count: ", result.ModifiedCount)
	}

	return nil

}

func (tr *taskRepository) UpdateTaskStatusByID(c context.Context, status string, id string) error {
	_, filename, line, _ := runtime.Caller(1)

	collection := tr.database.Collection(tr.collection)

	update := bson.M{
		"status":     status,
		"updated_at": time.Now(),
	}

	idHex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Printf("[error] %s:%d %v\n", filename, line, err)
		return err
	}

	result, err := collection.UpdateOne(c, bson.M{"_id": idHex}, bson.M{"$set": update})
	if err != nil {
		log.Printf("[error] %s:%d %v\n", filename, line, err)
		return err
	}

	if result.ModifiedCount > 0 {
		log.Println("Successfully updated the status. modified count: ", result.ModifiedCount)
	}

	return nil

}

func (tr *taskRepository) DeleteTaskByID(c context.Context, id string) error {
	_, filename, line, _ := runtime.Caller(1)

	collection := tr.database.Collection(tr.collection)

	idHex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Printf("[error] %s:%d %v\n", filename, line, err)
		return err
	}

	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})
	if err != nil {
		log.Printf("[error] %s:%d %v\n", filename, line, err)
		return err
	}

	return nil
}
