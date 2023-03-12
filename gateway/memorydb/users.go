package memorydb

import (
	"context"
	"fmt"
	"sync"

	"github.com/ShintaNakama/my-graphql-example/domain/entity"
)

type usersImpl struct {
	users map[string]*user
	lock  sync.RWMutex
}

type user struct {
	ID   string
	Name string
}

func (u *usersImpl) Insert(ctx context.Context, e *entity.User) error {
	u.lock.Lock()
	defer u.lock.Unlock()

	m := &user{
		ID:   e.ID,
		Name: e.Name,
	}

	if _, ok := u.users[e.ID]; ok {
		return fmt.Errorf("users already exists, id:%s", e.ID)
	}

	u.users[e.ID] = m

	return nil
}

func (u *usersImpl) List(ctx context.Context) ([]*entity.User, error) {
	u.lock.RLock()
	defer u.lock.RUnlock()

	users := make([]*entity.User, 0, len(u.users))
	for _, m := range u.users {
		user := &entity.User{
			ID:   m.ID,
			Name: m.Name,
		}

		users = append(users, user)
	}

	return users, nil
}

func (u *usersImpl) GetByID(ctx context.Context, id string) (*entity.User, error) {
	u.lock.RLock()
	defer u.lock.RUnlock()

	user, ok := u.users[id]
	if !ok {
		return nil, fmt.Errorf("users not found, id:%s", id)
	}

	return &entity.User{
		ID:   user.ID,
		Name: user.Name,
	}, nil
}
