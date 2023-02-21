package types

import (
	// "github.com/cosmos/cosmos-sdk/codec"
	// "github.com/cosmos/cosmos-sdk/codec/types"
	// "github.com/cosmos/cosmos-sdk/types/msgservice"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&AddStudentRequest{}, "aa", nil)
	cdc.RegisterConcrete(&RegisterAdminRequest{}, "abcd", nil)
	cdc.RegisterConcrete(&ApplyLeaveRequest{}, "aaa", nil)
	cdc.RegisterConcrete(&AcceptLeaveRequest{}, "aaaa", nil)
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
	// RegisterLegacyAminoCodec(authzcodec.Amino)
}
