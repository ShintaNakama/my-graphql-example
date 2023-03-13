package ctxutil

import (
	"context"
)

type key struct{}

const (
	roleCtxKey = iota
)

func RoleKey(ctx context.Context) string {
	v := ctx.Value(roleCtxKey)

	roleKey, ok := v.(string)
	if !ok {
		return ""
	}

	return roleKey
}

// SetRoleKey set Role Key
func SetRoleKey(parent context.Context, key string) context.Context {
	return context.WithValue(parent, roleCtxKey, key)
}
