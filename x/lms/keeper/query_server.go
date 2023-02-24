package keeper

import (
	//"bytes"
	"context"

	"github.com/chandiniv1/COSMOS-LMS1/x/lms/types"
	//sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	//"github.com/cosmos/cosmos-sdk/store/prefix"
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

	//ctx := sdk.UnwrapSDKContext(goCtx)
	//store := ctx.KVStore(k.storeKey)

	// studentStore:=prefix.NewStore(store,types.StudentKey)

	// var students []*types.Student


	return &types.GetStudentsResponse{}, nil
}

func (k queryServer) GetStudent(goCtx context.Context,req *types.GetStudentRequest) (*types.GetStudentResponse, error) {
	// ctx := sdk.UnwrapSDKContext(goCtx)

	// if err := k.GetStudnt(ctx, req.Address); err != nil {
	// 	return nil,err
	// }
	return &types.GetStudentResponse{}, nil
}

func (k queryServer) GetLeaveRequests(context.Context, *types.GetLeaveRequestsRequest) (*types.GetLeaveApprovesResponse, error) {
	return &types.GetLeaveApprovesResponse{}, nil
}

func (k queryServer) GetLeaveApproves(context.Context, *types.GetLeaveApprovesRequest) (*types.GetLeaveApprovesResponse, error) {
	return &types.GetLeaveApprovesResponse{}, nil
}
