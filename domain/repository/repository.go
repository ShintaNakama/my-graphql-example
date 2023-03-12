package repository

import (
	"context"

	"github.com/ShintaNakama/my-graphql-example/domain/entity"
)

type Repository interface {
	Users() Users
	Todos() Todos
}

type Todos interface {
	Insert(ctx context.Context, e *entity.NewTodo) error
	List(ctx context.Context) ([]*entity.Todo, error)
}

type Users interface {
	Insert(ctx context.Context, e *entity.User) error
	List(ctx context.Context) ([]*entity.User, error)
	GetByID(ctx context.Context, id string) (*entity.User, error)
}
