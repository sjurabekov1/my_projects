package service

import (
	"context"
	pb "hyssa/hyssa_go_billing_service/genproto/billing_service_v2"

	"hyssa/hyssa_go_billing_service/internal/core/ports"
	"hyssa/hyssa_go_billing_service/internal/core/repository"
	"hyssa/hyssa_go_billing_service/internal/core/repository/psql/sqlc"
	"hyssa/hyssa_go_billing_service/internal/pkg/logger"
	"hyssa/hyssa_go_billing_service/internal/pkg/processor"
	"hyssa/hyssa_go_billing_service/internal/pkg/serialize"

	"google.golang.org/protobuf/types/known/emptypb"
	"gopkg.in/guregu/null.v4"
)

var (
	_ ports.PocketService = (*Pocket)(nil)
)

type Pocket struct {
	pb.UnimplementedPocketServiceServer
	repo repository.Store
}

func NewPocketService(repo repository.Store) *Pocket {
	return &Pocket{
		repo: repo,
	}
}

func (p *Pocket) CreateRequiredPockets(ctx context.Context, req *pb.CreatePocketRequest) (
	resp *pb.PocketList, err error) {
	pockets, err := p.repo.GetAllPocketTypes(ctx, sqlc.GetAllPocketTypesParams{Required: true})
	if err != nil {
		return
	}

	createPocketParams := make([]sqlc.CreatePocketParams, len(pockets))
	for i, pocket := range pockets {
		createPocketParams[i] = sqlc.CreatePocketParams{
			UserID:    null.StringFrom(req.UserId),
			Type:      null.IntFrom(int64(pocket.ID)),
			Title:     pocket.Name,
			CompanyID: null.StringFrom(req.CompanyId),
		}
	}

	insertCount, err := p.repo.CreatePocket(ctx, createPocketParams)
	if err != nil {
		return
	}

	logger.Log.Infof("Inserted %d rows into pockets for user id: %s", insertCount, req.UserId)

	userPockets, err := p.repo.GetPocketsByUserID(ctx, null.StringFrom(req.UserId))
	if err != nil {
		return
	}

	resp = &pb.PocketList{}
	for i := range userPockets {
		resp.Pockets = append(resp.Pockets, &pb.Pocket{
			Id:        userPockets[i].OID,
			UserId:    userPockets[i].UserID.String,
			Type:      userPockets[i].PocketTypeName.String,
			Name:      userPockets[i].Name.String,
			Title:     userPockets[i].Title.String,
			Icon:      userPockets[i].Icon.String,
			CreatedAt: userPockets[i].CreatedAt.Time.String(),
			Company:   userPockets[i].CompanyName.String,
		})
	}

	return
}

func (p *Pocket) CreateCustomPocket(ctx context.Context, req *pb.CreatePocketRequest) (
	resp *pb.Pocket, err error) {
	return
}

func (p *Pocket) UpdatePocket(ctx context.Context, req *pb.UpdatePocketRequest) (
	resp *pb.Pocket, err error) {
	return
}

func (p *Pocket) GetPocket(ctx context.Context, req *pb.GetPocketRequest) (
	resp *pb.Pocket, err error) {
	return
}

func (p *Pocket) GetTransitPocket(ctx context.Context, req *pb.GetPocketRequest) (
	resp *pb.Pocket, err error) {
	return
}

func (p *Pocket) GetAllPocketsOfUser(ctx context.Context, req *pb.GetAllPocketsOfUserRequest) (
	resp *pb.PocketList, err error) {
	return
}

func (p *Pocket) TransferPocketOrders(ctx context.Context, req *pb.TransferPocketOrdersRequest) (
	resp *emptypb.Empty, err error) {
	return
}

func (p *Pocket) DeletePocket(ctx context.Context, req *pb.DeletePocketRequest) (
	resp *pb.Pocket, err error) {
	return
}

func (p *Pocket) DeletePocketWithoutTransfer(ctx context.Context, req *pb.DeletePocketRequest) (
	resp *pb.Pocket, err error) {
	return
}

func (p *Pocket) GetPocketSalesforceId(ctx context.Context, req *pb.GetPocketSalesforceIdRequest) (
	resp *pb.GetPocketSalesforceIdResponse, err error) {
	return
}

func (p *Pocket) GetPocketsData(ctx context.Context, req *pb.GetPocketsDataRequest) (
	resp *pb.GetPocketsDataResponse, err error) {
	return
}

func (p *Pocket) GetPocketTypes(ctx context.Context, req *pb.GetPocketTypesRequest) (
	resp *pb.GetPocketTypesResponse, err error) {
	dbResp, err := processor.ExecuteWithResp(ctx, req, p.repo.GetAllPocketTypes)
	if err != nil {
		return
	}

	resp = &pb.GetPocketTypesResponse{}
	err = serialize.MarshalUnMarshal(ctx, dbResp, &resp.PocketsType)
	if err != nil {
		return
	}

	return
}
