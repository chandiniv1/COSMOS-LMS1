package types

import (
	"errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	//sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	//"time"
)

var (
	_ sdk.Msg = &GetStudentRequest{}
	_ sdk.Msg = &GetAdminRequest{}
)

func NewGetStudentRequest(id string, address string) *GetStudentRequest {
	return &GetStudentRequest{
		Id:      id,
		Address: address,
	}
}

func (msg GetStudentRequest) GetSignBytes() []byte {
	return []byte{}
}

// GetSigners Implements Msg.
func (msg GetStudentRequest) GetSigners() []sdk.AccAddress {
	fromAddress, _ := sdk.AccAddressFromBech32(msg.Address)
	return []sdk.AccAddress{fromAddress}
}

func (msg GetStudentRequest) ValidateBasic() error {
	if msg.Id == "" {
		return errors.New("Id cant be null")
	} else if msg.Address == "" {
		return errors.New("Address cant be null")
	} else {
		return errors.New("Basic validations done successfully")
	}
}

func NewGetAdminRequest(address string) *GetAdminRequest {
	return &GetAdminRequest{
		Address: address,
	}
}

func (msg GetAdminRequest) GetSignBytes() []byte {
	return []byte{}
}

// GetSigners Implements Msg.
func (msg GetAdminRequest) GetSigners() []sdk.AccAddress {
	fromAddress, _ := sdk.AccAddressFromBech32(msg.Address)
	return []sdk.AccAddress{fromAddress}
}

func (msg GetAdminRequest) ValidateBasic() error {

	if msg.Address == "" {
		return errors.New("Address cant be null")
	} else {
		return errors.New("Basic validations done successfully")
	}
}
