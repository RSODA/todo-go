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
	"github.com/stretchr/testify/require"
)

func TestCreateService(t *testing.T) {
	type todoRepoMock func(mc *minimock.Controller) repo.Repo

	type args struct {
		ctx context.Context
		req *models.TODO
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
		todoRepoMock todoRepoMock
	}{
		{
			name: "success",
			args: args{
				ctx: ctx,
				req: &models.TODO{
					Title: "Void",
					Tasks: []models.Task{
						{
							Description: "Create Void",
						},
						{
							Description: "Delete Void",
						},
					},
				},
			},
			want: &models.TODO{
				ID:    0,
				Title: "Void",
				Tasks: []models.Task{
					{
						ID:          0,
						TodoID:      0,
						Description: "Create Void",
						IsComplete:  false,
					},
					{
						ID:          1,
						TodoID:      0,
						Description: "Delete Void",
						IsComplete:  false,
					},
				},
			},
			err: nil,
			todoRepoMock: func(mc *minimock.Controller) repo.Repo {
				mock := mocks.NewRepoMock(mc)

				mock.CreateMock.Expect(ctx, &models.TODO{
					Title: "Void",
					Tasks: []models.Task{
						{
							Description: "Create Void",
						},
						{
							Description: "Delete Void",
						},
					},
				}).Return(&models.TODO{
					ID:    0,
					Title: "Void",
					Tasks: []models.Task{
						{
							ID:          0,
							TodoID:      0,
							Description: "Create Void",
							IsComplete:  false,
						},
						{
							ID:          1,
							TodoID:      0,
							Description: "Delete Void",
							IsComplete:  false,
						},
					},
				}, nil)

				return mock
			},
		},
		{
			name: "Err case. Empty title",
			args: args{
				ctx: ctx,
				req: &models.TODO{
					Title: "",
					Tasks: []models.Task{
						{
							Description: "Create",
						},
						{
							Description: "Delete",
						},
					},
				},
			},
			want: nil,
			err:  errors.New("title or tasks required"),
			todoRepoMock: func(mc *minimock.Controller) repo.Repo {
				mock := mocks.NewRepoMock(mc)

				return mock
			},
		},
		{
			name: "Err case. Empty tasks",
			args: args{
				ctx: ctx,
				req: &models.TODO{
					Title: "Empty",
				},
			},
			want: nil,
			err:  errors.New("title or tasks required"),
			todoRepoMock: func(mc *minimock.Controller) repo.Repo {
				mock := mocks.NewRepoMock(mc)
				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			rMock := tt.todoRepoMock(mc)
			todoService := todo.NewTODOService(rMock)

			got, err := todoService.Create(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, got)
		})
	}
}
