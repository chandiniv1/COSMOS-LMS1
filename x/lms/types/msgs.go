package types

import (

	//"context"
	sdk "github.com/cosmos/cosmos-sdk/types"

	//"github.com/chandiniv1/COSMOS-LMS1/x/lms/types"
	"time"
)

var (
	_ sdk.Msg = &AddStudentRequest{}
	_ sdk.Msg = &AcceptLeaveRequest{}
	_ sdk.Msg = &ApplyLeaveRequest{}
	_ sdk.Msg = &RegisterAdminRequest{}
)

func NewAddStudentRequest(admin string, address string, name string, id string) *AddStudentRequest {
	return &AddStudentRequest{
		Admin:   admin,
		Address: address,
		Name:    name,
		Id:      id,
	}
}

func (msg AddStudentRequest) GetSignBytes() []byte {
	return []byte{}
}

// GetSigners Implements Msg.
func (msg AddStudentRequest) GetSigners() []sdk.AccAddress {
	fromAddress, _ := sdk.AccAddressFromBech32(msg.Admin)
	return []sdk.AccAddress{fromAddress}
}

func (msg AddStudentRequest) ValidateBasic() error {
	// if _, err := sdk.AccAddressFromBech32("./"); err != nil {
	// 	return sdkerrors.ErrInvalidAddress.Wrapf("invalid from address: %s", err)
	// }
	if msg.Admin == "" {
		return ErrAdminAddressNil
	} else if msg.Address == "" {
		return ErrStudentAddressNil
	} else if msg.Id == "" {
		return ErrStudentIdNil
	} else if msg.Name == "" {
		return ErrStudentNameNil
	} else {
		return nil
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
	fromAddress, _ := sdk.AccAddressFromBech32(msg.Admin)
	return []sdk.AccAddress{fromAddress}
}

func (msg AcceptLeaveRequest) ValidateBasic() error {
	// if _, err := sdk.AccAddressFromBech32(msg.Admin); err != nil {
	// 	return sdkerrors.ErrInvalidAddress.Wrapf("invalid from address: %s", err)
	// }
	if msg.Admin == "" {
		return ErrAdminAddressNil
	} else if msg.LeaveId == "" {
		return ErrStudentIdNil
	} else {
		return nil
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
	// if _, err := sdk.AccAddressFromBech32(""); err != nil {
	// 	return sdkerrors.ErrInvalidAddress.Wrapf("invalid from address: %s", err)
	// }
	if (msg.Address) == "" {
		return ErrStudentAddressNil
	} else if (msg.From) == nil {
		return ErrStudentDatesNil
	} else if (msg.To) == nil {
		return ErrStudentDatesNil
	} else if (msg.Reason) == "" {
		return ErrEmptyReason
	} else {
		return nil
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
		return ErrAdminAddressNil
	} else if msg.Name == "" {
		return ErrAdminNameNil
	} else {
		return nil
	}

}
