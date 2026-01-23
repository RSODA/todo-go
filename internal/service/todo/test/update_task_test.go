package test

import (
	"context"
	"errors"
	"go/types"
	"testing"

	"github.com/RSODA/todo-go/internal/models"
	"github.com/RSODA/todo-go/internal/repo"
	"github.com/RSODA/todo-go/internal/repo/mocks"
	"github.com/RSODA/todo-go/internal/service/todo"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
)

func TestUpdateTaskService(t *testing.T) {
	type todoRepoMockFunc func(mc *minimock.Controller) repo.Repo

	type args struct {
		ctx context.Context
		req *models.UpdateTaskRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)
	)

	tests := []struct {
		name         string
		args         args
		want         types.Nil
		err          error
		todoRepoMock todoRepoMockFunc
	}{
		{
			name: "success",
			args: args{
				ctx: ctx,
				req: &models.UpdateTaskRequest{
					ID:         0,
					IsComplete: false,
				},
			},
			want: types.Nil{},
			err:  nil,
			todoRepoMock: func(mc *minimock.Controller) repo.Repo {
				mock := mocks.NewRepoMock(mc)
				mock.UpdateTaskMock.Expect(ctx, &models.UpdateTaskRequest{
					ID:         0,
					IsComplete: false,
				}).Return(nil)

				return mock
			},
		},
		{
			name: "error nil models",
			args: args{
				ctx: ctx,
				req: nil,
			},
			want: types.Nil{},
			err:  errors.New("UpdateTaskRequest cannot be nil"),
			todoRepoMock: func(mc *minimock.Controller) repo.Repo {
				mock := mocks.NewRepoMock(mc)
				return mock
			},
		},
		{
			name: "error on repo layer",
			args: args{
				ctx: ctx,
				req: &models.UpdateTaskRequest{
					ID:         1,
					IsComplete: false,
				},
			},
			want: types.Nil{},
			err:  errors.New("repo layer error"),
			todoRepoMock: func(mc *minimock.Controller) repo.Repo {
				mock := mocks.NewRepoMock(mc)
				mock.UpdateTaskMock.Expect(ctx, &models.UpdateTaskRequest{
					ID:         1,
					IsComplete: false,
				}).Return(errors.New("repo layer error"))

				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			service := todo.NewTODOService(tt.todoRepoMock(mc))

			err := service.UpdateTask(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
		})
	}
}
