package memorydb

import "github.com/ShintaNakama/my-graphql-example/domain/repository"

type memoryDB struct {
	users *usersImpl
	todos *todosImpl
}

// NewRepository returns an on-memory Repository.
func NewRepository() repository.Repository {
	return &memoryDB{
		users: &usersImpl{
			users: map[string]*user{},
		},
		todos: &todosImpl{
			todos:     map[string]*todo{},
			todoUsers: map[string]string{},
		},
	}
}

func (m *memoryDB) Todos() repository.Todos {
	return m.todos
}

func (m *memoryDB) Users() repository.Users {
	return m.users
}
