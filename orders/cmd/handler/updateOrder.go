package handler

import (
	"context"
	"time"

	"github.com/AlekseyPromet/trinity-example/models"
	sq "github.com/Masterminds/squirrel"
	"github.com/go-kit/kit/endpoint"
)

func (s *ServiceRPC) MakeUpdateOrderEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (resp interface{}, err error) {

		mRequest := models.Order{}
		var ok bool

		if mRequest, ok = request.(models.Order); !ok {
			return nil, errNotParseForm
		}

		update := sq.Update("orders").
			Set("updated_at", mRequest.Datetime.Format(time.RFC3339)).
			Where(sq.Eq{"id": mRequest.ID})

		if mRequest.Phone != "" {
			update.Set("phone", mRequest.Phone)
		}

		if mRequest.Customer != "" {
			update.Set("customer", mRequest.Customer)
		}

		query, args, err := update.ToSql()
		if err != nil {
			return nil, err
		}

		_, err = s.DatabaseOrm.ExecContext(ctx, query, args...)
		if err != nil {
			return nil, err
		}

		return mRequest.ID.String(), nil
	}
}
