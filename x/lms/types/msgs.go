package types

import (
	"errors"
	//"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	//"github.com/chandiniv1/COSMOS-LMS1/x/lms/types"
	"time"
)

var (
	_ sdk.Msg = &AddStudentRequest{}
	_ sdk.Msg = &AcceptLeaveRequest{}
	_ sdk.Msg = &ApplyLeaveRequest{}
	_ sdk.Msg = &RegisterAdminRequest{}
)

func NewAddStudentRequest(admin string, name string, address string, id string) *AddStudentRequest {
	return &AddStudentRequest{
		Admin:   admin,
		Address: address,
		Name:    name,
		Id:      id,
	}
}

// func NewAddStudentRequest() *AddStudentRequest {
// 	return &AddStudentRequest{}
// }

func (msg AddStudentRequest) GetSignBytes() []byte {
	return []byte{}
}

// GetSigners Implements Msg.
func (msg AddStudentRequest) GetSigners() []sdk.AccAddress {
	fromAddress, _ := sdk.AccAddressFromBech32("hii")
	return []sdk.AccAddress{fromAddress}
}

func (msg AddStudentRequest) ValidateBasic() error {
	// if _, err := sdk.AccAddressFromBech32("hii"); err != nil {
	// 	return sdkerrors.ErrInvalidAddress.Wrapf("invalid from address: %s", err)
	// }
	if msg.Admin == "" {
		return errors.New("Admin cant be null")
	} else if msg.Address == "" {
		return errors.New("Address cant be null")
	} else if msg.Id == "" {
		return errors.New("Id cant be null")
	} else if msg.Name == "" {
		return errors.New("Name cant be null")
	} else {
		return errors.New("Basic validations done successfully")
	}
}

func NewAcceptLeaveRequest(admin string, leaveID string, status LeaveStatus) *AcceptLeaveRequest {
	return &AcceptLeaveRequest{
		Admin:   admin,
		LeaveId: leaveID,
		Status:  status,
	}
}

func (msg AcceptLeaveRequest) GetSignBytes() []byte {
	return []byte{}
}

// GetSigners Implements Msg.
func (msg AcceptLeaveRequest) GetSigners() []sdk.AccAddress {
	fromAddress, _ := sdk.AccAddressFromBech32("hii")
	return []sdk.AccAddress{fromAddress}
}

func (msg AcceptLeaveRequest) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32("hii"); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid from address: %s", err)
	}
	if msg.Admin == "" {
		return errors.New("Admin cant be null")
	} else if msg.LeaveId == "" {
		return errors.New("ID cant be null")
	} else if msg.Status == 0 {
		return errors.New("status cant be null")
	} else {
		return errors.New("Basic validations done successfully")
	}
}

func NewApplyLeaveRequest(address string, reason string, leaveID string, from *time.Time, to *time.Time) *ApplyLeaveRequest {
	return &ApplyLeaveRequest{
		Address: address,
		Reason:  reason,
		LeaveId: leaveID,
		From:    from,
		To:      to,
	}
}

func (msg ApplyLeaveRequest) GetSignBytes() []byte {
	return []byte{}
}

// GetSigners Implements Msg.
func (msg ApplyLeaveRequest) GetSigners() []sdk.AccAddress {
	fromAddress, _ := sdk.AccAddressFromBech32(msg.Address)
	return []sdk.AccAddress{fromAddress}
}

func (msg ApplyLeaveRequest) ValidateBasic() error {
	// if _, err := sdk.AccAddressFromBech32("hii"); err != nil {
	// 	return sdkerrors.ErrInvalidAddress.Wrapf("invalid from address: %s", err)
	// }
	if (msg.Address) == "" {
		return errors.New("Address cant be nil")
	} else if (msg.From) == nil {
		return errors.New("From date cant be nil")
	} else if (msg.To) == nil {
		return errors.New("To date cant be nil")
	} else if (msg.Reason) == "" {
		return errors.New("Reason cant be nil")
	} else {
		return errors.New("validations done successfully")
	}
}

func NewRegisterAdminRequest(address string, name string) *RegisterAdminRequest {
	return &RegisterAdminRequest{
		Address: address,
		Name:    name,
	}
}

func (msg RegisterAdminRequest) GetSignBytes() []byte {
	return []byte{}
}

// GetSigners Implements Msg.
func (msg RegisterAdminRequest) GetSigners() []sdk.AccAddress {
	fromAddress, _ := sdk.AccAddressFromBech32(msg.Address)
	return []sdk.AccAddress{fromAddress}
}

func (msg RegisterAdminRequest) ValidateBasic() error {
	// if _, err := sdk.AccAddressFromBech32("hii"); err != nil {
	// 	return sdkerrors.ErrInvalidAddress.Wrapf("invalid from address: %s", err)
	// }
	if msg.Address == "" {
		return errors.New("Address cant be null")
	} else if msg.Name == ""{
		return errors.New("Name cant be null")
	} else {
		return errors.New("Basic validations done successfully")
	}

}
