package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/learningPlatform/internal/configs"
	"github.com/learningPlatform/internal/repository"
	"github.com/learningPlatform/internal/service"

	"github.com/learningPlatform/internal/transport/rest"
	"github.com/learningPlatform/internal/transport/rest/handlers"
)

func main() {
	ctx := context.Background()

	cfg, err := configs.New()
	if err != nil {
		log.Fatal("error when initializing the config", err)
	}

	conn, err := pgxpool.Connect(ctx, cfg.DB_URL)
	if err != nil {
		log.Printf("Unable to connect to database: %v\n", err)
		return
	}

	repository := repository.NewRepo(conn)
	services := service.NewService(repository.Course, repository.User)
	restHandler := handlers.NewHandler(services.CourseStorage, services.UserStorage)

	restSrv := rest.NewServer(restHandler.InitRouters(), cfg)

	go func() {
		log.Printf("[INFO] run http server on port: %d\n", cfg.Port)
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

	conn.Close()
}
