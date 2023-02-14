package handler

import (
	"context"

	"github.com/AlekseyPromet/trinity-example/models"
	sq "github.com/Masterminds/squirrel"
	"github.com/go-kit/kit/endpoint"
)

func (s *ServiceRPC) MakeUpdateOrderEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (resp interface{}, err error) {

		mRequest := models.OrderRes{}
		var ok bool

		if mRequest, ok = request.(models.OrderRes); !ok {
			return nil, errNotParseForm
		}

		update := sq.Update("orders").
			Set("updated_at", mRequest.Datetime).
			Where(sq.Eq{"id": mRequest.ID})

		if mRequest.Status != "" {
			update.Set("status", mRequest.Status)
		}

		query, args, err := update.ToSql()
		if err != nil {
			return nil, err
		}

		_, err = s.DatabaseOrm.ExecContext(ctx, query, args...)
		if err != nil {
			return nil, err
		}

		return mRequest, nil
	}
}
