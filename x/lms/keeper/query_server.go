package keeper

import (
	"context"

	"github.com/chandiniv1/cosmos-LMS/x/lms/types"
)

type queryServer struct {
	Keeper
	types.UnimplementedQueryServer
}

var _ types.QueryServer = queryServer{}

func (k queryServer) GetStudents(context.Context, *types.GetStudentsRequest) (*types.GetStudentsResponse, error) {
	return &types.GetStudentsResponse{}, nil
}

func (k queryServer) GetLeaveRequests(context.Context, *types.GetLeaveRequestsRequest) (*types.GetLeaveApprovesResponse, error) {
	return &types.GetLeaveApprovesResponse{}, nil
}

func (k queryServer) GetLeaveApproves(context.Context, *types.GetLeaveApprovesRequest) (*types.GetLeaveApprovesResponse, error) {
	return &types.GetLeaveApprovesResponse{}, nil
}
