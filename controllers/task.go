package controllers

import (
	"net/http"
	"todo-api/initializers"
	"todo-api/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)


type TaskController struct {
	DB *gorm.DB
}

func NewTaskController(DB *gorm.DB) TaskController {
	return TaskController{DB}
}

func GetTasks(c *gin.Context) {
	var tasks []models.Task
	initializers.DB.Find(&tasks)

	c.JSON(http.StatusOK, gin.H{"data": tasks})
}

func CreateTask(c *gin.Context) {
	var input models.CreateTask
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task := models.Task{Title: input.Title, Description: input.Description, IsCompleted: false}
	initializers.DB.Create(&task)

	c.JSON(http.StatusCreated, gin.H{"data": task})
}

func GetTaskById(c *gin.Context) {
	var task models.Task

	if err := initializers.DB.Find("id = ?", c.Param("id")).First(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"data": task})
}