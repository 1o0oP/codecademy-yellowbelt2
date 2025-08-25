package repository

import (
	"codecademy-yellowbelt2/core/domain/entity"
	"codecademy-yellowbelt2/infrastructure/interface/repository"
	"encoding/json"
	"errors"
	"os"
	"sync"
)

type FileTodoRepository struct {
	filename string
	mutex    sync.RWMutex
}

var _ repository.ITodoRepository = (*FileTodoRepository)(nil)

func NewFileTodoRepository(filename string) repository.ITodoRepository {
	return &FileTodoRepository{
		filename: filename,
	}
}

func (r *FileTodoRepository) load() (map[string]*entity.Todo, error) {
	todos := make(map[string]*entity.Todo)

	data, err := os.ReadFile(r.filename)
	if err != nil {
		if os.IsNotExist(err) {
			return todos, nil // Arquivo não existe ainda, retorna mapa vazio
		}
		return nil, err
	}

	if len(data) == 0 {
		return todos, nil // Arquivo vazio
	}

	if err := json.Unmarshal(data, &todos); err != nil {
		return nil, err
	}

	return todos, nil
}

func (r *FileTodoRepository) save(todos map[string]*entity.Todo) error {
	data, err := json.MarshalIndent(todos, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(r.filename, data, 0644)
}

func (r *FileTodoRepository) Create(todo *entity.Todo) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	todos, err := r.load()
	if err != nil {
		return err
	}

	todos[todo.ID] = todo
	return r.save(todos)
}

func (r *FileTodoRepository) GetByID(id string) (*entity.Todo, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	todos, err := r.load()
	if err != nil {
		return nil, err
	}

	todo, exists := todos[id]
	if !exists {
		return nil, errors.New("todo not found")
	}

	return todo, nil
}

func (r *FileTodoRepository) GetAll() ([]*entity.Todo, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	todos, err := r.load()
	if err != nil {
		return nil, err
	}

	todoList := make([]*entity.Todo, 0, len(todos))
	for _, todo := range todos {
		todoList = append(todoList, todo)
	}

	return todoList, nil
}

func (r *FileTodoRepository) Update(todo *entity.Todo) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	todos, err := r.load()
	if err != nil {
		return err
	}

	if _, exists := todos[todo.ID]; !exists {
		return errors.New("todo not found")
	}

	todos[todo.ID] = todo
	return r.save(todos)
}

func (r *FileTodoRepository) Delete(id string) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	todos, err := r.load()
	if err != nil {
		return err
	}

	if _, exists := todos[id]; !exists {
		return errors.New("todo not found")
	}

	delete(todos, id)
	return r.save(todos)
}
