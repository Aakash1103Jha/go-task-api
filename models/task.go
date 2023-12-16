// Description: This package contains the task model
// and the CRUD operations for the task model.

package models

import (
	"context"
	"time"

	"example.com/api/db"
	"go.mongodb.org/mongo-driver/bson"
)

type Task struct {
	Id          string    `json:"_id" bson:"_id" default:"primitive.NewObjectID().Hex()"`
	Title       string    `json:"title" required:"true" max:"100" min:"5"`
	Description string    `json:"description" required:"true" max:"500" min:"10"`
	Completed   bool      `json:"completed" required:"true" default:"false"`
	CreatedAt   time.Time `json:"created_at" default:"time.Now"`
	UpdatedAt   time.Time `json:"updated_at" default:"time.Now"`
	UserId      string    `json:"user_id"`
}

// get all tasks
func (t *Task) GetAllTasks() ([]Task, error) {
	var tasks []Task
	collection := db.InitMongo().Database("GoTaskApi").Collection("tasks")
	cursor, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}
	for cursor.Next(context.Background()) {
		var task Task
		err := cursor.Decode(&task)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

// create new task
func (t *Task) CreateTask() (bool, error) {
	collection := db.InitMongo().Database("GoTaskApi").Collection("tasks")
	_, err := collection.InsertOne(context.Background(), t)
	if err != nil {
		return false, err
	}
	return true, nil
}

// update task by id
func (t *Task) UpdateTask() {

}

// delete task by id
func (t *Task) DeleteTask() {

}
