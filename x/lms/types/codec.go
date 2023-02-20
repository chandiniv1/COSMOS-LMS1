package types

import (
	// "github.com/cosmos/cosmos-sdk/codec"
	// "github.com/cosmos/cosmos-sdk/codec/types"
	// "github.com/cosmos/cosmos-sdk/types/msgservice"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec"github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	authzcodec "github.com/cosmos/cosmos-sdk/x/authz/codec"
)

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&AddStudentRequest{}, "/leavemanagementsystem", nil)
	cdc.RegisterConcrete(&RegisterAdminRequest{}, "/leavemanagementsystem", nil)
	cdc.RegisterConcrete(&ApplyLeaveRequest{}, "/leavemanagementsystem", nil)
	cdc.RegisterConcrete(&AcceptLeaveRequest{}, "/leavemanagementsystem", nil)
}

func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&AddStudentRequest{},
		&RegisterAdminRequest{},
		&ApplyLeaveRequest{},
		&AcceptLeaveRequest{},
	)
	//MsgServer.RegisterMsgServiceDesc(registry, &Msg_ServiceDesc)
	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)

func init() {
	RegisterLegacyAminoCodec(amino)
	cryptocodec.RegisterCrypto(amino)
	sdk.RegisterLegacyAminoCodec(amino)
	RegisterLegacyAminoCodec(authzcodec.Amino)
}
