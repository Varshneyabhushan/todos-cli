package unmarkdone

type TodoService interface {
	UnmarkDone(id int) error
}

type MarkUndoneService = func(id int) error

func NewunmarkDoneService(s TodoService) MarkUndoneService {
	return func(id int) error {
		return s.UnmarkDone(id)
	}
}
