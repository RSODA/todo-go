package test

import (
	"context"
	"errors"
	"go/types"
	"testing"

	"github.com/RSODA/todo-go/internal/repo"
	"github.com/RSODA/todo-go/internal/repo/mocks"
	"github.com/RSODA/todo-go/internal/service/todo"
	"github.com/gojuno/minimock/v3"
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/require"
)

func TestDeleteService(t *testing.T) {
	type todoRepoMock func(mc *minimock.Controller) repo.Repo

	type args struct {
		ctx context.Context
		id  int64
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
		todoRepoMock todoRepoMock
	}{
		{
			name: "success",
			args: args{
				ctx: ctx,
				id:  1,
			},
			want: types.Nil{},
			err:  nil,
			todoRepoMock: func(mc *minimock.Controller) repo.Repo {
				mock := mocks.NewRepoMock(mc)

				mock.DeleteMock.Expect(ctx, 1).Return(nil)

				return mock
			},
		},
		{
			name: "error id <0",
			args: args{
				ctx: ctx,
				id:  -1,
			},
			want: types.Nil{},
			err:  errors.New("invalid id"),
			todoRepoMock: func(mc *minimock.Controller) repo.Repo {
				mock := mocks.NewRepoMock(mc)

				return mock
			},
		},
		{
			name: "Error not found",
			args: args{
				ctx: ctx,
				id:  0,
			},
			want: types.Nil{},
			err:  pgx.ErrNoRows,
			todoRepoMock: func(mc *minimock.Controller) repo.Repo {
				mock := mocks.NewRepoMock(mc)
				mock.DeleteMock.Expect(ctx, 0).Return(pgx.ErrNoRows)

				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			service := todo.NewTODOService(tt.todoRepoMock(mc))

			err := service.Delete(tt.args.ctx, tt.args.id)
			require.Equal(t, tt.err, err)
		})
	}
}
