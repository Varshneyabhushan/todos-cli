package add

type TodoService interface {
	Add(item string) error
}

type AddingService = func(item string) error

func NewAddingService(todoService TodoService) AddingService {
	return func(item string) error {
		return todoService.Add(item)
	}
}