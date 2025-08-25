package application

import (
	"codecademy-yellowbelt2/core/domain/entity"
	app_interfaces "codecademy-yellowbelt2/infrastructure/interface/application"
	"codecademy-yellowbelt2/infrastructure/interface/repository"
)

type TodoUseCase struct {
	todoRepo repository.ITodoRepository
}

func NewTodoUseCase(todoRepo repository.ITodoRepository) app_interfaces.ITodoUseCase {
	return &TodoUseCase{
		todoRepo: todoRepo,
	}
}

func (uc *TodoUseCase) CreateTodo(title, description string) (*entity.Todo, error) {
	todo := entity.NewTodo(title, description)
	err := uc.todoRepo.Create(todo)
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func (uc *TodoUseCase) GetTodoByID(id string) (*entity.Todo, error) {
	return uc.todoRepo.GetByID(id)
}

func (uc *TodoUseCase) GetAllTodos() ([]*entity.Todo, error) {
	return uc.todoRepo.GetAll()
}

func (uc *TodoUseCase) UpdateTodo(id, title, description string) (*entity.Todo, error) {
	todo, err := uc.todoRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	todo.Update(title, description)
	err = uc.todoRepo.Update(todo)
	if err != nil {
		return nil, err
	}

	return todo, nil
}

func (uc *TodoUseCase) CompleteTodo(id string) (*entity.Todo, error) {
	todo, err := uc.todoRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	todo.MarkAsCompleted()
	err = uc.todoRepo.Update(todo)
	if err != nil {
		return nil, err
	}

	return todo, nil
}

func (uc *TodoUseCase) DeleteTodo(id string) error {
	return uc.todoRepo.Delete(id)
}
