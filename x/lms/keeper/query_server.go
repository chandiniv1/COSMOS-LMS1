package keeper

import (
	"context"
	"fmt"

	"github.com/chandiniv1/COSMOS-LMS1/x/lms/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

type queryServer struct {
	Keeper
	types.UnimplementedQueryServer
}

var _ types.QueryServer = Keeper{}

func (k Keeper) GetStudents(goCtx context.Context, req *types.GetStudentsRequest) (*types.GetStudentsResponse, error) {
	if req == nil {
		return nil, sdkerrors.ErrInvalidRequest.Wrap("empty request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)
	students := k.GetStdnts(ctx, req)
	res := types.GetStudentsResponse{
		Students: students,
	}
	return &res, nil
}

func (k Keeper) GetStudent(goCtx context.Context, req *types.GetStudentRequest) (*types.GetStudentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	fmt.Println("gtstudent")
	k.GetStdnt(ctx, req.Id)
	return &types.GetStudentResponse{}, nil
}

func (k Keeper) GetAdmin(goCtx context.Context, req *types.GetAdminRequest) (*types.GetAdminResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	k.GetAdmn(ctx, req.Address)
	return &types.GetAdminResponse{}, nil
}

func (k Keeper) GetLeaveRequests(goCtx context.Context, req *types.GetLeaveRequestsRequest) (*types.GetLeaveApprovesResponse, error) {
	if req == nil {
		return nil, sdkerrors.ErrInvalidRequest.Wrap("empty request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)
	k.GetLeaveRqsts(ctx, req)
	return &types.GetLeaveApprovesResponse{}, nil
}

func (k Keeper) GetLeaveApproves(goCtx context.Context, req *types.GetLeaveApprovesRequest) (*types.GetLeaveApprovesResponse, error) {
	if req == nil {
		return nil, sdkerrors.ErrInvalidRequest.Wrap("empty request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)
	k.GetAcceptedLeaves(ctx, req)
	return &types.GetLeaveApprovesResponse{}, nil
}
