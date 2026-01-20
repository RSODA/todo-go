package app

import (
	"context"
	"log"

	"github.com/RSODA/todo-go/internal/config"
	"github.com/RSODA/todo-go/internal/handlers/todo"
	"github.com/RSODA/todo-go/internal/repo"
	"github.com/RSODA/todo-go/internal/repo/postgres"
	"github.com/RSODA/todo-go/internal/router"
	"github.com/RSODA/todo-go/internal/service"
	ts "github.com/RSODA/todo-go/internal/service/todo"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ServiceProvider struct {
	httpConfig config.HTTPConfig
	pgConfig   config.PGConfig

	db *pgxpool.Pool

	repo    repo.Repo
	service service.Service

	todoHandler *todo.Handler
	router      *gin.Engine
}

func NewService() *ServiceProvider {
	return &ServiceProvider{}
}

func (sp *ServiceProvider) HTTPConfig() config.HTTPConfig {
	if sp.httpConfig == nil {
		cfg, err := config.NewHTTPConfig()
		if err != nil {
			log.Fatal(err)
		}

		sp.httpConfig = cfg
	}

	return sp.httpConfig
}

func (sp *ServiceProvider) PGConfig(_ context.Context) config.PGConfig {
	if sp.pgConfig == nil {
		cfg, err := config.NewPGConfig()
		if err != nil {
			log.Fatal(err)
		}

		sp.pgConfig = cfg
	}

	return sp.pgConfig
}

func (sp *ServiceProvider) DB(ctx context.Context) *pgxpool.Pool {
	if sp.db == nil {
		cl, err := pgxpool.New(ctx, sp.PGConfig(ctx).DSN())
		if err != nil {
			log.Fatal("Failed to connect to database", err)
		}

		err = cl.Ping(ctx)
		if err != nil {
			log.Fatal("Failed to ping database", err)
		}

		sp.db = cl
	}

	return sp.db
}

func (sp *ServiceProvider) Repo(ctx context.Context) repo.Repo {
	if sp.repo == nil {
		r := postgres.NewPostgres(sp.DB(ctx))
		sp.repo = r
	}

	return sp.repo
}

func (sp *ServiceProvider) Service(ctx context.Context) service.Service {
	if sp.service == nil {
		s := ts.NewTODOService(sp.Repo(ctx))

		sp.service = s
	}

	return sp.service
}

func (sp *ServiceProvider) Handler(ctx context.Context) *todo.Handler {
	if sp.todoHandler == nil {
		h := todo.NewHandler(sp.Service(ctx))

		sp.todoHandler = h
	}

	return sp.todoHandler
}

func (sp *ServiceProvider) Router(ctx context.Context) *gin.Engine {
	if sp.router == nil {
		r := router.NewRouter(sp.Handler(ctx))

		sp.router = r
	}

	return sp.router
}
