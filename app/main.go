package main

import (
	"context"
	"log"

	"MyProject/service"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kelseyhightower/envconfig"
)

func main() {

	ctx := context.Background()

	var cfg Config
	if err := envconfig.Process("", &cfg); err != nil {
		log.Fatal(err)
	}
	cfg.Print()

	conn, err := pgxpool.New(ctx, cfg.PgConnectUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	repos := service.Repositories{
		Users: users_repo.New(ctx, conn),
	}
	srvc := service.New(repos)	

	srv := &http.Server{
		Addr: cfg.Listen,
		Handler: server.New(cfg.Listen, srvc),
	}

}
