package services

import (
	"errors"
	"sort"
)

type TodoItem struct {
	Id int
	Text string
	IsDone bool 
}

type TodoService struct {
	CurrentId int
	storage JSONStorage
	Todos map[int]TodoItem
}

func NewTodoService(storage JSONStorage) (TodoService, error) {
	newService := TodoService{
		CurrentId: 0,
		storage: storage,
		Todos: make(map[int]TodoItem, 0),
	}

	if err := storage.LoadJSON(&newService); err != nil {
		return newService, errors.New("error while loading from storage : " + err.Error())
	}

	return newService, nil
}

func (s TodoService) save() error {
	return s.storage.SaveJSON(s)
}

func (s *TodoService) Add(text string) error {
	newTodo := TodoItem{
		Id : s.CurrentId,
		Text : text,
		IsDone: false,
	}

	s.CurrentId += 1
	s.Todos[newTodo.Id] = newTodo
	return s.save()
}

func (s *TodoService) List() ([]TodoItem, error) {
	var result []TodoItem
	for _, todoItem := range s.Todos {
		result = append(result, todoItem)
	}

	sort.SliceStable(result, func(i, j int) bool {
		return result[i].Id < result[j].Id
	})
	
	return result, nil
}

func (s *TodoService) has(id int) bool {
	_, found := s.Todos[id]
	return found
}

func (s *TodoService) Delete(id int) error {
	if !s.has(id) {
		return errors.New("todo not found")
	}

	delete(s.Todos, id)
	return s.save()
}

func (s *TodoService) MarkDone(id int) error {
	if !s.has(id) {
		return errors.New("todo not found")
	}

	todo, _ := s.Todos[id]
	todo.IsDone = true
	s.Todos[id] = todo
	return s.save()
}

func (s *TodoService) UnmarkDone(id int) error {
	if !s.has(id) {
		return errors.New("todo not found")
	}

	todo, _ := s.Todos[id]
	todo.IsDone = false
	s.Todos[id] = todo
	return s.save()
}