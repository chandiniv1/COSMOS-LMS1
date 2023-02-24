package keeper

import (
	"context"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/chandiniv1/COSMOS-LMS1/x/lms/types"
)

type queryServer struct {
	Keeper
	types.UnimplementedQueryServer
}

var _ types.QueryServer = queryServer{}

func (k queryServer) GetStudents(goCtx context.Context,req *types.GetStudentsRequest) (*types.GetStudentsResponse, error) {
	if req==nil{
		return nil,sdkerrors.ErrInvalidRequest.Wrap("empty request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	store := ctx.KVStore(k.storeKey)


	return &types.GetStudentsResponse{}, nil
}

func (k queryServer) GetStudent(context.Context, *types.GetStudentRequest) (*types.GetStudentResponse, error) {
	return &types.GetStudentResponse{}, nil
}

func (k queryServer) GetLeaveRequests(context.Context, *types.GetLeaveRequestsRequest) (*types.GetLeaveApprovesResponse, error) {
	return &types.GetLeaveApprovesResponse{}, nil
}

func (k queryServer) GetLeaveApproves(context.Context, *types.GetLeaveApprovesRequest) (*types.GetLeaveApprovesResponse, error) {
	return &types.GetLeaveApprovesResponse{}, nil
}
