package inmemory

import (
	"github.com/MrTj458/fiber-api-todo/model"
	"github.com/google/uuid"
)

// This is used to make sure TodoService implements the model.TodoService
// interface
var _ model.TodoService = (*TodoService)(nil)

type TodoService struct {
	// s is the in-memory map used to simulate a database
	s map[uuid.UUID]*model.Todo
}

func NewTodoService() *TodoService {
	return &TodoService{
		s: make(map[uuid.UUID]*model.Todo),
	}
}

func (ts *TodoService) FindTodoByID(id uuid.UUID) (*model.Todo, error) {
	t, exists := ts.s[id]
	if !exists {
		return nil, model.ErrNotFound
	}

	return t, nil
}

func (ts *TodoService) AllTodos() ([]*model.Todo, error) {
	list := []*model.Todo{}

	for _, t := range ts.s {
		list = append(list, t)
	}

	return list, nil
}

func (ts *TodoService) CreateTodo(todo model.TodoCreate) (*model.Todo, error) {
	newTodo := &model.Todo{
		Desc:      todo.Desc,
		Completed: false,
		ID:        uuid.New(),
	}

	ts.s[newTodo.ID] = newTodo
	return newTodo, nil
}

func (ts *TodoService) UpdateTodo(id uuid.UUID, upd model.TodoUpdate) (*model.Todo, error) {
	t, exists := ts.s[id]
	if !exists {
		return nil, model.ErrNotFound
	}

	t.Desc = upd.Desc
	t.Completed = upd.Completed

	return t, nil
}

func (ts *TodoService) DeleteTodo(id uuid.UUID) error {
	delete(ts.s, id)
	return nil
}
