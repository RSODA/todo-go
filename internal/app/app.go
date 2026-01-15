package app

import (
	"context"

	"github.com/RSODA/todo-go/internal/config"
)

type App struct {
	serviceProvider *serviceProvider
}

func NewApp(ctx context.Context) (*App, error) {
	a := &App{}

	err := a.initDeps(ctx)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(ctx context.Context) error{
		a.initConfig,
		a.initServiceProvider,
	}

	for _, init := range inits {
		err := init(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initConfig(ctx context.Context) error {
	err := config.Load(".env")
	if err != nil {
		return err
	}

	return nil
}

func (a *App) initServiceProvider(ctx context.Context) error {
	a.serviceProvider = NewService()
	return nil
}

func (a *App) RunHTTP() error {
	return a.serviceProvider.Router().Run(
		a.serviceProvider.HTTPConfig().Address(),
	)
}
