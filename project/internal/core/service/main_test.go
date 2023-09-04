package service

import (
	"context"
	"hyssa/hyssa_go_billing_service/internal/config"
	"hyssa/hyssa_go_billing_service/internal/core/repository"
	"hyssa/hyssa_go_billing_service/mocks"
	"os"
	"testing"
)

var (
	cfg = &config.Config{
		PSQL: struct {
			URI string "env:\"PSQL_URI\""
		}{
			URI: "postgres://postgres:postgres@localhost:5432/s3?sslmode=disable",
		},
	}

	repos repository.Store
)

func TestMain(m *testing.M) {
	mocks.MockAppLogger()

	repos = repository.New(context.Background(), cfg)
	os.Exit(m.Run())
}
