package memorydb

import (
	"context"
	"fmt"
	"sync"

	"github.com/ShintaNakama/my-graphql-example/domain/entity"
)

type todosImpl struct {
	todos     map[string]*todo
	todoUsers map[string]string
	lock      sync.RWMutex
}

type todo struct {
	ID   string
	Text string
	Done bool
}

func (t *todosImpl) Insert(ctx context.Context, e *entity.NewTodo) error {
	t.lock.Lock()
	defer t.lock.Unlock()

	m := &todo{
		ID:   e.ID,
		Text: e.Text,
		Done: false,
	}

	if _, ok := t.todos[e.ID]; ok {
		return fmt.Errorf("todo already exists, id:%s", e.ID)
	}

	if _, ok := t.todoUsers[e.ID]; ok {
		return fmt.Errorf("todo_users already exists, todo_id:%s, user_id:%s", e.ID, e.UserID)
	}

	t.todos[e.ID] = m
	t.todoUsers[e.ID] = e.UserID

	return nil
}

func (t *todosImpl) List(ctx context.Context) ([]*entity.Todo, error) {
	t.lock.RLock()
	defer t.lock.RUnlock()

	todos := make([]*entity.Todo, 0, len(t.todos))
	for _, m := range t.todos {
		todo := &entity.Todo{
			ID:     m.ID,
			Text:   m.Text,
			Done:   m.Done,
			UserID: t.todoUsers[m.ID],
		}

		todos = append(todos, todo)
	}

	return todos, nil
}
