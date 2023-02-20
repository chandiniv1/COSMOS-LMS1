package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/chandiniv1/COSMOS-LMS1/x/lms/types"
)

var _ types.MsgServer = msgServer{}

type msgServer struct {
	Keeper
	types.UnimplementedMsgServer
}

func (k msgServer) AddStudent(goCtx context.Context, req *types.AddStudentRequest) (*types.AddStudentResponse, error) {
	ctx:=sdk.UnwrapSDKContext(goCtx)
	
	if err := k.AddStd(ctx,req); err != nil {
		return nil, err
	}

	// ctx.EventManager().EmitTypedEvent(&types.EventAddStudent{
	// 	admin:Admin,
	// 	address:Address,
	// 	name:Name,
	// 	id:Id,
	// })
	return &types.AddStudentResponse{},nil
	
}

func (k msgServer) ApplyLeave(c context.Context, a *types.ApplyLeaveRequest) (*types.ApplyLeaveResponse, error) {
	stdnt := types.ApplyLeaveRequest{}
	k.cdc.MustMarshal(&stdnt)

	return &types.ApplyLeaveResponse{}, nil
}

func (k msgServer) AcceptLeave(c context.Context,a *types.AcceptLeaveRequest) (*types.AcceptLeaveResponse, error) {
	return &types.AcceptLeaveResponse{}, nil
}

func (k msgServer) RegisterAdmin(c context.Context,a *types.RegisterAdminRequest) (*types.RegisterAdminResponse, error) {
	return &types.RegisterAdminResponse{}, nil
}


// admin, err := sdk.AccAddressFromBech32(msg.Admin)
	// if err != nil {
	// 	return nil, err
	// }
	// address, err := sdk.AccAddressFromBech32(msg.Address)
	// if err != nil {
	// 	return nil, err
	// }
	// name, err := sdk.AccAddressFromBech32(msg.Name)
	// if err != nil {
	// 	return nil, err
	// }
	// id, err := sdk.AccAddressFromBech32(msg.Id)
	// if err != nil {
	// 	return nil, err
	// }