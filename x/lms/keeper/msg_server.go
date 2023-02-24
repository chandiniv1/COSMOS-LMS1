package keeper

import (
	"context"

	"github.com/chandiniv1/COSMOS-LMS1/x/lms/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ types.MsgServer = msgServer{}

type msgServer struct {
	Keeper
	types.UnimplementedMsgServer
}

func (k msgServer) AddStudent(goCtx context.Context, req *types.AddStudentRequest) (*types.AddStudentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := k.AddStdnt(ctx, req); err != nil {
		return nil, err
	}

	// ctx.EventManager().EmitTypedEvent(&types.EventAddStudent{
	// 	admin:Admin,
	// 	address:Address,
	// 	name:Name,
	// 	id:Id,
	// })
	return &types.AddStudentResponse{}, nil

}

func (k msgServer) ApplyLeave(goCtx context.Context, req *types.ApplyLeaveRequest) (*types.ApplyLeaveResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := k.AplyLeave(ctx, req); err != nil {
		return nil, err
	}
	return &types.ApplyLeaveResponse{}, nil
}

func (k msgServer) AcceptLeave(goCtx context.Context, req *types.AcceptLeaveRequest) (*types.AcceptLeaveResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := k.AcptLeave(ctx, req); err != nil {
		return nil, err
	}
	return &types.AcceptLeaveResponse{}, nil
}

func (k msgServer) RegisterAdmin(goCtx context.Context, req *types.RegisterAdminRequest) (*types.RegisterAdminResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := k.RgstrAdmin(ctx, req); err != nil {
		return nil, err
	}
	return &types.RegisterAdminResponse{}, nil
}
