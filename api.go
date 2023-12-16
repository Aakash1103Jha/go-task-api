package main

import (
	"fmt"
	"net/http"

	"example.com/api/models"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// load the .env file
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	// create a new gin server
	server := gin.Default()

	// define a new route
	server.GET("/", getAllTasks)
	server.GET("/:id", getOneTaskById)
	server.POST("/", createTaskHandler)
	server.PUT("/:id", updateTaskHandler)
	server.DELETE("/:id", deleteTaskHandler)

	// start the server
	server.Run(":4000")
}

func getOneTaskById(c *gin.Context) {
}
func getAllTasks(c *gin.Context) {
	var tasks models.Task
	data, err := tasks.GetAllTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"tasks": data})
}
func createTaskHandler(c *gin.Context) {
	var task models.Task
	err := c.BindJSON(&task)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	task.Id = string(primitive.NewObjectID().Hex())
	task.CreateTask()
	c.JSON(http.StatusCreated, gin.H{"task": task})
}
func updateTaskHandler(c *gin.Context) {
	id := c.Params.ByName("id")
	fmt.Println("Task ID to update: ", id)
}
func deleteTaskHandler(c *gin.Context) {
	id := c.Params.ByName("id")
	fmt.Println("Task ID to delete: ", id)
}
