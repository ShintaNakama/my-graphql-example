package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ShintaNakama/my-graphql-example/gateway/memorydb"
	"github.com/ShintaNakama/my-graphql-example/handler"
	"golang.org/x/sync/errgroup"
)

func main() {
	port := flag.Int("port", 8080, "port on which the server will listen")
	flag.Parse()

	runServer(context.Background(), *port)
}

func runServer(ctx context.Context, port int) error {
	log.Printf("connect to http://localhost:%d/ for GraphQL playground", port)
	ctx, stop := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	defer stop()

	repo := memorydb.NewRepository()

	handler.RegisterGQLHandlers(repo)

	srv := http.Server{
		Addr: fmt.Sprintf(":%d", port),
	}

	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			return fmt.Errorf("failed to run server: %w", err)
		}
		return nil
	})
	eg.Go(func() error {
		<-ctx.Done()
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		log.Printf("server shutdown %s", srv.Addr)
		if err := srv.Shutdown(ctx); err != nil {
			return fmt.Errorf("failed to shutdown server: %w", err)
		}
		return nil
	})
	log.Printf("server listening on %s", srv.Addr)
	return eg.Wait()
}
