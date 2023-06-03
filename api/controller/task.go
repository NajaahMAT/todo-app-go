package controller

import (
	"fmt"
	"net/http"
	"todo-app-go/bootstrap"
	"todo-app-go/domain"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskController struct {
	TaskUsecase domain.TaskUsecase
	Env         *bootstrap.Env
}

func (t *TaskController) CreateTask(c *gin.Context) {
	var request domain.CreateTaskReq

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	fmt.Println("request: ", request)

	userId, err := primitive.ObjectIDFromHex(request.UserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Invalid userid"})
	}

	task := domain.Task{
		ID:          primitive.NewObjectID(),
		Title:       request.Title,
		Status:      request.Status,
		Description: request.Description,
		File:        request.File,
		CreatedAt:   request.CreatedAt,
		UserID:      userId,
	}

	fmt.Println("create task request: ", task)

	err = t.TaskUsecase.Create(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, nil)
}

func (t *TaskController) GetAllTasksByUserID(c *gin.Context) {
	var request domain.CreateTaskReq

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	fmt.Println("request: ", request)

	userId, err := primitive.ObjectIDFromHex(request.UserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Invalid userid"})
	}

	task := domain.Task{
		ID:          primitive.NewObjectID(),
		Title:       request.Title,
		Status:      request.Status,
		Description: request.Description,
		File:        request.File,
		CreatedAt:   request.CreatedAt,
		UserID:      userId,
	}

	fmt.Println("create task request: ", task)

	err = t.TaskUsecase.Create(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, nil)
}
