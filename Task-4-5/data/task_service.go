package data

import (
	"errors"
	"sync"
	"task_manager/models"

	"github.com/google/uuid"
)

var (
	ErrTaskNotFound = errors.New("task not found")
)

type TaskService struct {
	tasks map[string]models.Task
	mu    sync.RWMutex
}

func NewTaskService() *TaskService {
	return &TaskService{
		tasks: make(map[string]models.Task),
	}
}

func (s *TaskService) CreateTask(task models.Task) models.Task {
	s.mu.Lock()
	defer s.mu.Unlock()

	task.ID = uuid.New().String()
	s.tasks[task.ID] = task
	return task
}

func (s *TaskService) GetAllTasks() []models.Task {
	s.mu.RLock()
	defer s.mu.RUnlock()

	tasks := make([]models.Task, 0, len(s.tasks))
	for _, task := range s.tasks {
		tasks = append(tasks, task)
	}
	return tasks
}

func (s *TaskService) GetTaskByID(id string) (models.Task, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	task, exists := s.tasks[id]
	if !exists {
		return models.Task{}, ErrTaskNotFound
	}
	return task, nil
}

func (s *TaskService) UpdateTask(id string, updatedTask models.Task) (models.Task, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.tasks[id]; !exists {
		return models.Task{}, ErrTaskNotFound
	}

	updatedTask.ID = id
	s.tasks[id] = updatedTask
	return updatedTask, nil
}

func (s *TaskService) DeleteTask(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.tasks[id]; !exists {
		return ErrTaskNotFound
	}

	delete(s.tasks, id)
	return nil
}