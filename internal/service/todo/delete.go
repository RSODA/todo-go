package todo

import (
	"context"
	"errors"
)

func (s *todoService) Delete(ctx context.Context, id int64) error {
	if id < 0 {
		return errors.New("invalid id")
	}

	err := s.r.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
