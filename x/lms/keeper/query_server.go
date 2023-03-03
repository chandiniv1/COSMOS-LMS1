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

var _ types.QueryServer = Keeper{}

func (k Keeper) GetStudents(goCtx context.Context, req *types.GetStudentsRequest) (*types.GetStudentsResponse, error) {
	if req == nil {
		return nil, sdkerrors.ErrInvalidRequest.Wrap("empty request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)
	students := k.GetStdnts(ctx, req)
	// resStd := []*types.AddStudentRequest{}
	// for _, student := range students {
	// 	resStd = append(resStd, &student)
	// }
	res := types.GetStudentsResponse{
		Students: students,
	}
	return &res, nil
}

func (k Keeper) GetStudent(goCtx context.Context, req *types.GetStudentRequest) (*types.GetStudentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	//fmt.Println("gtstudent")
	student, _ := k.GetStdnt(ctx, req.Address)
	res := types.GetStudentResponse{
		Student: &student,
	}
	return &res, nil
}

func (k Keeper) GetAdmin(goCtx context.Context, req *types.GetAdminRequest) (*types.GetAdminResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	admin, _ := k.GetAdmn(ctx, req.Address)
	res := types.GetAdminResponse{
		Admin: &admin,
	}
	return &res, nil
}

func (k Keeper) GetLeaveRequests(goCtx context.Context, req *types.GetLeaveRequestsRequest) (*types.GetLeaveRequestsResponse, error) {
	if req == nil {
		return nil, sdkerrors.ErrInvalidRequest.Wrap("empty request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)
	leavereqs := k.GetLeaveRqsts(ctx, req)
	
	res := types.GetLeaveRequestsResponse{
		Leaverequest: leavereqs,
	}
	return &res, nil
}

func (k Keeper) GetLeaveApproves(goCtx context.Context, req *types.GetLeaveApprovesRequest) (*types.GetLeaveApprovesResponse, error) {
	if req == nil {
		return nil, sdkerrors.ErrInvalidRequest.Wrap("empty request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)
	leaveapproves := k.GetAcceptedLeaves(ctx, req)
	res := types.GetLeaveApprovesResponse{
		Leaveapprove: leaveapproves,
	}
	return &res, nil
}
