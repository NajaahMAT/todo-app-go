package controller

import (
	"fmt"
	"net/http"
	"time"
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

	// dateTime, _ := time.Parse(domain.DateTimeLayout, request.CreatedAt)

	task := domain.Task{
		ID:          primitive.NewObjectID(),
		Title:       request.Title,
		Status:      request.Status,
		Description: request.Description,
		File:        request.File,
		CreatedAt:   time.Now(),
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

func (t *TaskController) GetTasksByUserID(c *gin.Context) {
	userId := c.Param("user_id")

	fmt.Println("request user_id: ", userId)

	tasks, err := t.TaskUsecase.FetchAllTasksByUserID(c, userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

func (t *TaskController) GetTasksByStatus(c *gin.Context) {
	userId := c.Param("user_id")
	status := c.Param("status")

	fmt.Println("request user_id: ", userId)

	tasks, err := t.TaskUsecase.FetchTasksByStatusAndUserID(c, userId, status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

func (t *TaskController) GetTaskByID(c *gin.Context) {
	id := c.Param("id")

	fmt.Println("request task id: ", id)

	tasks, err := t.TaskUsecase.FetchTaskByID(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

func (t *TaskController) GetTasksByDateRange(c *gin.Context) {
	userId := c.Param("user_id")
	startDate := c.Param("start_date")
	endDate := c.Param("end_date")

	fmt.Println("request user_id: ", userId)

	tasks, err := t.TaskUsecase.FetchTasksByDateRangeAndUserID(c, userId, startDate, endDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

func (t *TaskController) DeleteTaskByID(c *gin.Context) {
	id := c.Param("id")

	fmt.Println("request user_id: ", id)

	err := t.TaskUsecase.DeleteTaskByID(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, nil)
}

func (t *TaskController) UpdateTask(c *gin.Context) {
	var request domain.UpdateTaskReq

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	fmt.Println("request: ", request)

	id, err := primitive.ObjectIDFromHex(request.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Invalid userid"})
	}

	task := domain.Task{
		ID:          id,
		Title:       request.Title,
		Status:      request.Status,
		Description: request.Description,
		File:        request.File,
	}

	fmt.Println("create task request: ", task)

	err = t.TaskUsecase.UpdateTaskByID(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

func (t *TaskController) UpdateStatus(c *gin.Context) {
	var request domain.UpdateStatusReq

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	fmt.Println("request: ", request)

	err = t.TaskUsecase.UpdateTaskStatusByID(c, request.Status, request.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
