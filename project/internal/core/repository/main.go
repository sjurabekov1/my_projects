package repository

import (
	"context"

	"hyssa/hyssa_go_billing_service/internal/config"
	"hyssa/hyssa_go_billing_service/internal/core/repository/psql"
	"hyssa/hyssa_go_billing_service/internal/core/repository/psql/sqlc"
)

type Store interface {
	sqlc.Querier
}

func New(ctx context.Context, cfg *config.Config) Store {
	return psql.NewStore(ctx, cfg.PSQL.URI)
}
