package handler

import (
	"context"
	"errors"
	"time"

	"github.com/AlekseyPromet/trinity-example/models"
	sq "github.com/Masterminds/squirrel"
	"github.com/go-kit/kit/endpoint"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

var (
	errUnexpected       = errors.New("unexpected error")
	errNotHttpTransport = errors.New("unexpected transport protocol")
	errNotParseForm     = errors.New("not parse form request")
)

type ServiceRPC struct {
	CtxService          context.Context
	DatabaseOrm         *sqlx.DB
	StatusEndpoint      endpoint.Endpoint
	CreateOrderEndpoint endpoint.Endpoint
	UpdateOrderEndpoint endpoint.Endpoint
	DeleteOrderEndpoint endpoint.Endpoint
}

func (s *ServiceRPC) MakeStatusEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (resp interface{}, err error) {

		if err := s.DatabaseOrm.Ping(); err != nil {
			return "Bad", err
		}

		return "OK", nil
	}
}

func (s *ServiceRPC) MakeCreateOrderEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (resp interface{}, err error) {

		mRequest := models.Order{}
		var ok bool

		if mRequest, ok = request.(models.Order); !ok {
			return nil, errNotParseForm
		}

		id := uuid.New()

		query, args, err := sq.Insert("orders").
			Columns("id", "phone", "customer", "created_at").
			Values(id, mRequest.Phone, mRequest.Customer, mRequest.Datetime.Format(time.RFC3339)).
			ToSql()
		if err != nil {
			return nil, err
		}

		_, err = s.DatabaseOrm.ExecContext(ctx, query, args...)
		if err != nil {
			return nil, err
		}

		return id.String(), nil
	}
}
