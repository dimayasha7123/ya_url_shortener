package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"ya_url_shortener/internal/adapters/mem_repo"
	"ya_url_shortener/internal/adapters/postgres_repo"
	"ya_url_shortener/internal/app"
	"ya_url_shortener/internal/inputs/grpc_server"
	"ya_url_shortener/pkg/api"
)

const (
	repoCapacity = 10000
	inMemoryRepo = "in_memory"
	postgresRepo = "postgres"
)

var (
	node        = int64(1)
	host        = "0.0.0.0"
	port        = "8001"
	repo        = inMemoryRepo
	postgresDsn = "host=localhost port=5432 user=postgres password=postgres dbname=shortener sslmode=disable"
)

func main() {
	flag.Int64Var(&node, "node", node, "uniq node number, need for url generator")
	flag.StringVar(&host, "host", host, "grpc server host")
	flag.StringVar(&port, "port", port, "grpc server port")
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
	server := grpc_server.New(shortener)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen TCP on address %v: %v", addr, err)
	}

	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(grpc_server.LogInterceptor),
	}
	grpcServer := grpc.NewServer(opts...)
	api.RegisterURLShortenerServer(grpcServer, server)

	go func() {
		err = grpcServer.Serve(lis)
		if err != nil && err != grpc.ErrServerStopped {
			log.Fatalf("Error while server working: %v", err)
		}
	}()
	log.Println("Server is running!")

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-done
	fmt.Println()
	log.Println("Server has been stopped")

	grpcServer.GracefulStop()
	log.Println("Server exited properly")
}
