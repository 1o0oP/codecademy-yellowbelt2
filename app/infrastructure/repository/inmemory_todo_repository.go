package repository

import (
	"codecademy-yellowbelt2/core/domain/entity"
	"codecademy-yellowbelt2/infrastructure/interface/repository"
	"errors"
	"sync"
)

type InMemoryTodoRepository struct {
	todos map[string]*entity.Todo
	mutex sync.RWMutex
}

var _ repository.ITodoRepository = (*InMemoryTodoRepository)(nil)

func NewInMemoryTodoRepository() repository.ITodoRepository {
	return &InMemoryTodoRepository{
		todos: make(map[string]*entity.Todo),
	}
}

// ...existing code...
func (r *InMemoryTodoRepository) Create(todo *entity.Todo) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	r.todos[todo.ID] = todo
	return nil
}

func (r *InMemoryTodoRepository) GetByID(id string) (*entity.Todo, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	todo, exists := r.todos[id]
	if !exists {
		return nil, errors.New("todo not found")
	}
	return todo, nil
}

func (r *InMemoryTodoRepository) GetAll() ([]*entity.Todo, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	todos := make([]*entity.Todo, 0, len(r.todos))
	for _, todo := range r.todos {
		todos = append(todos, todo)
	}
	return todos, nil
}

func (r *InMemoryTodoRepository) Update(todo *entity.Todo) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, exists := r.todos[todo.ID]; !exists {
		return errors.New("todo not found")
	}

	r.todos[todo.ID] = todo
	return nil
}

func (r *InMemoryTodoRepository) Delete(id string) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, exists := r.todos[id]; !exists {
		return errors.New("todo not found")
	}

	delete(r.todos, id)
	return nil
}
