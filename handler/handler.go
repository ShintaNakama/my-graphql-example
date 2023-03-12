package handler

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/ShintaNakama/my-graphql-example/domain/repository"
	"github.com/ShintaNakama/my-graphql-example/graph"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// RegisterGQLHandlers is register graphql endpoint on http handler.
func RegisterGQLHandlers(repo repository.Repository) {
	r := graph.NewResolver(repo)

	gqlconfig := graph.Config{Resolvers: r}

	server := handler.NewDefaultServer(graph.NewExecutableSchema(gqlconfig))

	server.SetRecoverFunc(func(ctx context.Context, err interface{}) error {
		fmt.Fprintln(os.Stdout, err)
		return &gqlerror.Error{
			Message: "recover panic",
			Path:    graphql.GetPath(ctx),
			Extensions: map[string]interface{}{
				"httpStatusCode": 500,
			},
		}
	})

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", server)
}
