package test

import (
	"context"
	"go/types"
	"testing"

	"github.com/RSODA/todo-go/internal/repo"
	"github.com/RSODA/todo-go/internal/repo/mocks"
	"github.com/RSODA/todo-go/internal/service/todo"
	"github.com/gojuno/minimock/v3"
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
				id:  int64(1),
			},
			want: types.Nil{},
			err:  nil,
			todoRepoMock: func(mc *minimock.Controller) repo.Repo {
				mock := mocks.NewRepoMock(mc)

				mock.DeleteMock.Expect(ctx, int64(1)).Return(nil)

				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			rMock := tt.todoRepoMock(mc)
			service := todo.NewTODOService(rMock)

			err := service.Delete(tt.args.ctx, tt.args.id)
			require.Equal(t, tt.err, err)
		})
	}
}
