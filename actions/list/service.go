package list

import "todos/services"

type TodoService interface {
	List() ([]services.TodoItem, error)
}

type ListService = func () ([]services.TodoItem, error)

func NewListService(todoService TodoService) ListService {
	return func() ([]services.TodoItem, error) {
		return todoService.List()
	}
}