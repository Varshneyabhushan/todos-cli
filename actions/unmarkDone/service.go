package unmarkdone

type TodoService interface {
	MarkUndone(id int) error
}

type MarkUndoneService = func(id int) error

func NewMarkUndoneService(s TodoService) MarkUndoneService {
	return func(id int) error {
		return s.MarkUndone(id)
	}
}
