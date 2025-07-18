package data

import (
	"context"
	"errors"
	"task_manager/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskService struct {
	collection *mongo.Collection
}

func NewTaskService() *TaskService {
	return &TaskService{
		collection: taskCollection,
	}
}

func (s *TaskService) CreateTask(task models.Task) (*models.Task, error) {
	// Set timestamps
	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()

	// Insert into MongoDB
	result, err := s.collection.InsertOne(context.Background(), task)
	if err != nil {
		return nil, err
	}

	// Get the inserted ID and update the task
	task.ID = result.InsertedID.(primitive.ObjectID)
	return &task, nil
}

func (s *TaskService) GetAllTasks() ([]models.Task, error) {
	var tasks []models.Task

	cursor, err := s.collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	if err = cursor.All(context.Background(), &tasks); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (s *TaskService) GetTaskByID(id string) (*models.Task, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid task ID")
	}

	var task models.Task
	err = s.collection.FindOne(context.Background(), bson.M{"_id": objectID}).Decode(&task)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("task not found")
		}
		return nil, err
	}

	return &task, nil
}

func (s *TaskService) UpdateTask(id string, updatedTask models.Task) (*models.Task, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid task ID")
	}

	// Set updated timestamp
	updatedTask.UpdatedAt = time.Now()

	update := bson.M{
		"$set": updatedTask,
	}

	_, err = s.collection.UpdateOne(
		context.Background(),
		bson.M{"_id": objectID},
		update,
	)
	if err != nil {
		return nil, err
	}

	// Return the updated task
	return s.GetTaskByID(id)
}

func (s *TaskService) DeleteTask(id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid task ID")
	}

	_, err = s.collection.DeleteOne(context.Background(), bson.M{"_id": objectID})
	return err
}