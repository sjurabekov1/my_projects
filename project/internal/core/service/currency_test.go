package service

import (
	"context"
	pb "hyssa/hyssa_go_billing_service/genproto/billing_service_v2"
	"hyssa/hyssa_go_billing_service/internal/core/repository"
	"reflect"
	"testing"
)

func TestCurrency_CreateCurrency(t *testing.T) {
	type fields struct {
		UnimplementedCurrencyServiceServer pb.UnimplementedCurrencyServiceServer
		repo                               repository.Store
	}
	type args struct {
		ctx context.Context
		req *pb.CreateCurrencyRequest
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantResp *pb.CreateCurrencyResponse
		wantErr  bool
	}{
		{
			name: "TestCurrency_CreateCurrency",
			fields: fields{
				repo: repos,
			},
			args: args{
				ctx: context.Background(),
				req: &pb.CreateCurrencyRequest{
					Name:   "INR",
					Icon:   "INR",
					Active: true,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Currency{
				UnimplementedCurrencyServiceServer: tt.fields.UnimplementedCurrencyServiceServer,
				repo:                               tt.fields.repo,
			}
			gotResp, err := c.CreateCurrency(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Currency.CreateCurrency() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("Currency.CreateCurrency() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
