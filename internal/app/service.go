package app

import (
	"context"
	"log"

	"github.com/RSODA/todo-go/internal/config"
	"github.com/RSODA/todo-go/internal/handlers/todo"
	"github.com/RSODA/todo-go/internal/repo"
	"github.com/RSODA/todo-go/internal/repo/postgress"
	"github.com/RSODA/todo-go/internal/router"
	"github.com/RSODA/todo-go/internal/service"
	ts "github.com/RSODA/todo-go/internal/service/todo"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type serviceProvider struct {
	httpConfig config.HTTPConfig
	pgConfig   config.PGConfig

	db *pgxpool.Pool

	repo    repo.Repo
	service service.Service

	todoHandler *todo.Handler
	router      *gin.Engine
}

func NewService() *serviceProvider {
	return &serviceProvider{}
}

func (sp *serviceProvider) HTTPConfig() config.HTTPConfig {
	if sp.httpConfig == nil {
		cfg, err := config.NewHTTPConfig()
		if err != nil {
			log.Fatal(err)
		}

		sp.httpConfig = cfg
	}

	return sp.httpConfig
}

func (sp *serviceProvider) PGConfig(_ context.Context) config.PGConfig {
	if sp.pgConfig == nil {
		cfg, err := config.NewPGConfig()
		if err != nil {
			log.Fatal(err)
		}

		sp.pgConfig = cfg
	}

	return sp.pgConfig
}

func (sp *serviceProvider) DB(ctx context.Context) *pgxpool.Pool {
	if sp.db == nil {
		cl, err := pgxpool.New(ctx, sp.pgConfig.DSN())
		if err != nil {
			log.Fatal("Failed to connect to database", err)
		}

		err = cl.Ping(ctx)
		if err != nil {
			log.Fatalf("Failed to ping database", err)
		}

		sp.db = cl
	}

	return sp.db
}

func (sp *serviceProvider) Repo(_ context.Context) repo.Repo {
	if sp.repo == nil {
		repo := postgress.NewPostgres(sp.db)
		sp.repo = repo
	}

	return sp.repo
}

func (sp *serviceProvider) Service(_ context.Context) service.Service {
	if sp.service == nil {
		s := ts.NewTODOService(sp.repo)

		sp.service = s
	}

	return sp.service
}

func (sp *serviceProvider) Handler(_ context.Context) *todo.Handler {
	if sp.todoHandler == nil {
		h := todo.NewHandler(sp.service)

		sp.todoHandler = h
	}

	return sp.todoHandler
}

func (sp *serviceProvider) Router() *gin.Engine {
	if sp.router == nil {
		r := router.NewRouter(sp.todoHandler)

		sp.router = r
	}

	return sp.router
}
