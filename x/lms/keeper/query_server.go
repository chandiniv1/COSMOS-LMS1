package keeper

import (
	"context"

	"github.com/chandiniv1/COSMOS-LMS1/x/lms/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

type queryServer struct {
	Keeper
	types.UnimplementedQueryServer
}

var _ types.QueryServer = queryServer{}

func (k queryServer) GetStudents(goCtx context.Context, req *types.GetStudentsRequest) (*types.GetStudentsResponse, error) {
	if req == nil {
		return nil, sdkerrors.ErrInvalidRequest.Wrap("empty request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)
	k.GetStdnts(ctx, req)

	return &types.GetStudentsResponse{}, nil
}

func (k queryServer) GetStudent(goCtx context.Context, req *types.GetStudentRequest) (*types.GetStudentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	k.GetStdnt(ctx, req.Address)
	return &types.GetStudentResponse{}, nil
}

func (k queryServer) GetAdmin(goCtx context.Context, req *types.GetAdminRequest) (*types.GetAdminResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	k.GetStdnt(ctx, req.Address)
	return &types.GetAdminResponse{}, nil
}

func (k queryServer) GetLeaveRequests(goCtx context.Context, req *types.GetLeaveRequestsRequest) (*types.GetLeaveApprovesResponse, error) {
	if req == nil {
		return nil, sdkerrors.ErrInvalidRequest.Wrap("empty request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)
	k.GetLeaveRqsts(ctx, req)
	return &types.GetLeaveApprovesResponse{}, nil
}

func (k queryServer) GetLeaveApproves(goCtx context.Context, req *types.GetLeaveApprovesRequest) (*types.GetLeaveApprovesResponse, error) {
	if req == nil {
		return nil, sdkerrors.ErrInvalidRequest.Wrap("empty request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)
	k.GetAcceptedLeaves(ctx, req)
	return &types.GetLeaveApprovesResponse{}, nil
}
