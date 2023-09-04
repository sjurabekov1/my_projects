package ports

import (
	"context"

	pb "hyssa/hyssa_go_billing_service/genproto/billing_service_v2"

	"google.golang.org/protobuf/types/known/emptypb"
)

type PocketService interface {
	CreateRequiredPockets(ctx context.Context, req *pb.CreatePocketRequest) (resp *pb.PocketList, err error)
	CreateCustomPocket(ctx context.Context, req *pb.CreatePocketRequest) (resp *pb.Pocket, err error)
	UpdatePocket(ctx context.Context, req *pb.UpdatePocketRequest) (resp *pb.Pocket, err error)
	GetPocket(ctx context.Context, req *pb.GetPocketRequest) (resp *pb.Pocket, err error)
	GetTransitPocket(ctx context.Context, req *pb.GetPocketRequest) (resp *pb.Pocket, err error)
	GetAllPocketsOfUser(ctx context.Context, req *pb.GetAllPocketsOfUserRequest) (resp *pb.PocketList, err error)
	TransferPocketOrders(ctx context.Context, req *pb.TransferPocketOrdersRequest) (resp *emptypb.Empty, err error)
	DeletePocket(ctx context.Context, req *pb.DeletePocketRequest) (resp *pb.Pocket, err error)
	DeletePocketWithoutTransfer(ctx context.Context, req *pb.DeletePocketRequest) (resp *pb.Pocket, err error)
	GetPocketSalesforceId(ctx context.Context, req *pb.GetPocketSalesforceIdRequest) (resp *pb.GetPocketSalesforceIdResponse, err error)
	GetPocketsData(ctx context.Context, req *pb.GetPocketsDataRequest) (resp *pb.GetPocketsDataResponse, err error)
	GetPocketTypes(ctx context.Context, req *pb.GetPocketTypesRequest) (resp *pb.GetPocketTypesResponse, err error)
}
