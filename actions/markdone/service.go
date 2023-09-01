package markdone 

type TodoService interface {
	MarkDone(id int) error 
}

type MarkDoneService = func(id int) error

func NewMarkDoneService(s TodoService) MarkDoneService {
	return func(id int) error {
		return s.MarkDone(id)
	}
}