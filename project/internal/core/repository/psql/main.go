package psql

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"

	"hyssa/hyssa_go_billing_service/internal/core/repository/psql/sqlc"
	"hyssa/hyssa_go_billing_service/internal/pkg/logger"
)

type SQLStore struct {
	*sqlc.Queries
	DB *pgxpool.Pool
}

func NewStore(ctx context.Context, psqlUri string) *SQLStore {
	logger.Log.Info("connecting to psql...")
	dbConn, err := pgxpool.Connect(ctx, psqlUri)
	if err != nil {
		logger.Log.Fatal("failed to connecto to psql", err)
	}

	logger.Log.Info("psql connected")
	return &SQLStore{
		Queries: sqlc.New(dbConn),
		DB:      dbConn,
	}
}
