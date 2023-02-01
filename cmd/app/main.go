package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/jackc/pgx"
	"github.com/learningPlatform/internal/repository"
	"github.com/learningPlatform/internal/service"

	"github.com/learningPlatform/internal/transport/rest"
	"github.com/learningPlatform/internal/transport/rest/handlers"
)

func main() {
	ctx := context.Background()

	// TODO use configuration struct with env
	port := 9000

	conn, err := pgx.Connect(pgx.ConnConfig{
		Host:     "localhost",
		Port:     5431,
		Database: "course_db",
		User:     "admin",
		Password: "admin123",
	})
	if err != nil {
		log.Printf("Unable to connect to database: %v\n", err)
		return
	}

	repository := repository.NewRepo(conn)
	service := service.NewService(repository)

	restHandler := handlers.NewHandler(service)
	restSrv := rest.NewServer(port, restHandler.InitRouters())

	//grpcHandler := j
	//grpcSrv :=

	go func() {
		log.Printf("[INFO] run http server port %d", port)
		if err := restSrv.ListenAndServe(); err != nil {
			log.Printf("[ERROR] listen and serve error: %s", err)
			return
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT)

	<-ch

	err = restSrv.Stop(ctx)
	if err != nil {
		log.Printf("[ERROR] stop server error: %s", err)
	}

	err = conn.Close()
	if err != nil {
		log.Printf("[ERROR] stop database error: %s", err)
	}
}
