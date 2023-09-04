package services

import (
	"errors"
)

type TodoItem struct {
	Id int
	Text string
	IsDone bool 
}

type TodoService struct {
	CurrentId int
	storage JSONStorage
	Todos []TodoItem
}

func NewTodoService(storage JSONStorage) (TodoService, error) {
	newService := TodoService{
		CurrentId: 0,
		storage: storage,
		Todos: make([]TodoItem, 0),
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
	s.Todos = append(s.Todos, newTodo)
	return s.save()
}

func (s *TodoService) List() ([]TodoItem, error) {
	return s.Todos, nil
}

func (s *TodoService) has(index int) bool {
	return 0 <= index && index < len(s.Todos)
}

func (s *TodoService) Delete(index int) error {
	if !s.has(index) {
		return errors.New("todo not found")
	}

	s.Todos = append(s.Todos[:index], s.Todos[index + 1:]...)
	return s.save()
}

func (s *TodoService) MarkDone(index int) error {
	if !s.has(index) {
		return errors.New("todo not found")
	}

	s.Todos[index].IsDone = true
	return s.save()
}

func (s *TodoService) UnmarkDone(index int) error {
	if !s.has(index) {
		return errors.New("todo not found")
	}

	s.Todos[index].IsDone = false
	return s.save()
}