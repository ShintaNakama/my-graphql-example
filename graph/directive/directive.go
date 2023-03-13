package directive

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/ShintaNakama/my-graphql-example/ctxutil"
	"github.com/ShintaNakama/my-graphql-example/graph/model"
)

type directiveImpl struct {
}

func NewDirective() *directiveImpl {
	return &directiveImpl{}
}

func (d directiveImpl) Mask(ctx context.Context, obj interface{}, next graphql.Resolver, roles []model.Role) (res interface{}, err error) {
	// @maskが付与されたフィールドの値を取得する
	v, err := next(ctx)
	if v == nil || err != nil {
		return v, err
	}

	// ctxからroleを取得
	role := model.Role(ctxutil.RoleKey(ctx))
	// roleが設定されたリクエストでかつschemaで定義した値である場合maskしない
	if role.IsValid() {
		return v, nil
	}

	// 上記以外はmask
	return "xxx", nil
}
