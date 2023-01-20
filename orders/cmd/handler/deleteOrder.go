package handler

import (
	"context"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/go-kit/kit/endpoint"
	"github.com/google/uuid"
)

func (s *ServiceRPC) MakeDeleteOrderEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (resp interface{}, err error) {

		srcId := request.(string)

		id, err := uuid.Parse(srcId)
		if err != nil {
			return nil, err
		}

		update := sq.Update("orders").
			Set("is_deleted", "true").
			Set("updated_at", time.Now().Format(time.RFC3339)).
			Where(sq.Eq{"id": id})

		query, args, err := update.ToSql()
		if err != nil {
			return nil, err
		}

		_, err = s.DatabaseOrm.ExecContext(ctx, query, args...)
		if err != nil {
			return nil, err
		}

		return "OK", nil
	}
}
