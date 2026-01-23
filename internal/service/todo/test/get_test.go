package test

import (
	"context"
	"errors"
	"testing"

	"github.com/RSODA/todo-go/internal/models"
	"github.com/RSODA/todo-go/internal/repo"
	"github.com/RSODA/todo-go/internal/repo/mocks"
	"github.com/RSODA/todo-go/internal/service/todo"
	"github.com/gojuno/minimock/v3"
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/require"
)

func TestGetService(t *testing.T) {
	type todoRepoMockFunc func(mc *minimock.Controller) repo.Repo

	type args struct {
		ctx context.Context
		req int64
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)
	)

	tests := []struct {
		name         string
		args         args
		want         *models.TODO
		err          error
		todoRepoMock todoRepoMockFunc
	}{
		{
			name: "success",
			args: args{
				ctx: ctx,
				req: 1,
			},
			want: &models.TODO{
				ID:    1,
				Title: "test",
				Tasks: []models.Task{
					{
						ID:          0,
						TodoID:      1,
						Description: "test",
						IsComplete:  false,
					},
				},
			},
			err: nil,
			todoRepoMock: func(mc *minimock.Controller) repo.Repo {
				mock := mocks.NewRepoMock(mc)

				mock.GetMock.Expect(context.Background(), 1).Return(&models.TODO{
					ID:    1,
					Title: "test",
					Tasks: []models.Task{
						{
							ID:          0,
							TodoID:      1,
							Description: "test",
							IsComplete:  false,
						},
					},
				}, nil)

				return mock
			},
		},
		{
			name: "err id <0",
			args: args{
				ctx: ctx,
				req: -1,
			},
			want: nil,
			err:  errors.New("invalid id"),
			todoRepoMock: func(mc *minimock.Controller) repo.Repo {
				mock := mocks.NewRepoMock(mc)

				return mock
			},
		},
		{
			name: "error not found",
			args: args{
				ctx: ctx,
				req: 1,
			},
			want: nil,
			err:  pgx.ErrNoRows,
			todoRepoMock: func(mc *minimock.Controller) repo.Repo {
				mock := mocks.NewRepoMock(mc)
				mock.GetMock.Expect(context.Background(), 1).Return(nil, pgx.ErrNoRows)

				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			service := todo.NewTODOService(tt.todoRepoMock(mc))

			got, err := service.Get(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, got)
		})
	}
}
