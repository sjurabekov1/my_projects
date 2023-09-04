package service

import (
	"context"
	pb "hyssa/hyssa_go_billing_service/genproto/billing_service_v2"
	"hyssa/hyssa_go_billing_service/internal/core/repository"
	"reflect"
	"testing"
)

func TestCurrencyConverter_CreateCurrencyConvertor(t *testing.T) {
	type fields struct {
		UnimplementedCurrencyConverterServiceServer pb.UnimplementedCurrencyConverterServiceServer
		repo                                        repository.Store
	}
	type args struct {
		ctx context.Context
		req *pb.CreateCurrencyConvertorRequest
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantResp *pb.CreateCurrencyConvertorResponse
		wantErr  bool
	}{
		{
			name: "TestCurrencyConverter_CreateCurrencyConvertor",
			fields: fields{
				repo: repos,
				UnimplementedCurrencyConverterServiceServer: pb.UnimplementedCurrencyConverterServiceServer{},
			},
			args: args{
				ctx: context.Background(),
				req: &pb.CreateCurrencyConvertorRequest{
					PrimaryCurrency:   2,
					SecondaryCurrency: 4,
					Amount:            100,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CurrencyConverter{
				UnimplementedCurrencyConverterServiceServer: tt.fields.UnimplementedCurrencyConverterServiceServer,
				repo: tt.fields.repo,
			}
			gotResp, err := c.CreateCurrencyConvertor(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("CurrencyConverter.CreateCurrencyConvertor() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("CurrencyConverter.CreateCurrencyConvertor() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
