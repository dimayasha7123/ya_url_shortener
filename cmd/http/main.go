package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"ya_url_shortener/internal/adapters/mem_repo"
	"ya_url_shortener/internal/adapters/postgres_repo"
	"ya_url_shortener/internal/app"
	"ya_url_shortener/internal/inputs/http_server"
)

const (
	repoCapacity = 10000
	inMemoryRepo = "in_memory"
	postgresRepo = "postgres"
)

var (
	node        = int64(0)
	host        = "0.0.0.0"
	port        = "8000"
	repo        = inMemoryRepo
	postgresDsn = "host=localhost port=5432 user=postgres password=postgres dbname=shortener sslmode=disable"
)

func main() {
	flag.Int64Var(&node, "node", node, "uniq node number, need for url generator")
	flag.StringVar(&host, "host", host, "http server host")
	flag.StringVar(&port, "port", port, "http server port")
	flag.StringVar(&repo, "repo", repo,
		fmt.Sprintf("repository type, available %q and %q", inMemoryRepo, postgresRepo),
	)
	flag.StringVar(&postgresDsn, "postgres_dsn", postgresDsn,
		fmt.Sprintf("dsn to connect postgres, need if %q = %q", "repo", postgresRepo),
	)
	flag.Parse()

	ctx := context.Background()
	var repository app.Repository
	switch repo {
	case inMemoryRepo:
		repository = mem_repo.New(repoCapacity)
		log.Println("Create in memory repository")
	case postgresRepo:
		var err error
		repository, err = postgres_repo.New(ctx, postgresDsn)
		if err != nil {
			log.Fatalf("Can't create postgres repository: %v", err)
		}
		log.Println("Create postgres repository")
	default:
		log.Fatalf("Repo can't be %b", repo)
	}

	shortener, err := app.New(node, repository)
	if err != nil {
		log.Fatalf("can't create shortener app: %v", err)
	}

	addr := fmt.Sprintf("%s:%s", host, port)
	server := http_server.New(shortener, addr)

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			if errors.Is(err, http.ErrServerClosed) {
				log.Println("Server was closed")
			} else {
				log.Fatalf("Get error while serving: %v", err)
			}
		}
	}()
	log.Printf("Server was started on: http://%s", addr)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	fmt.Println()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	err = server.Shutdown(ctx)
	if err != nil {
		log.Fatalf("Get error while shutting down server: %v", err)
	}

	log.Println("Server was shutdown successfully!")
	os.Exit(0)
}
