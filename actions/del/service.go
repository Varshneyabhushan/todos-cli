package del

type TodoService interface {
	Delete(id int) error
}

type DeleteService = func(id int) error

func NewDeleteService(s TodoService) DeleteService {
	return func(id int) error {
		return s.Delete(id)
	}
}
