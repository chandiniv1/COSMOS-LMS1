package types

import (
	"strconv"
	sdk "github.com/cosmos/cosmos-sdk/types"
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
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// GetSigners Implements Msg.
func (msg AddStudentRequest) GetSigners() []sdk.AccAddress {
	fromAddress, _ := sdk.AccAddressFromBech32(msg.Admin)
	return []sdk.AccAddress{fromAddress}
}

func (msg AddStudentRequest) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Address); err != nil {
		return ErrInvalidAddress
	}
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

func NewAcceptLeaveRequest(admin string, leaveID string, status string) *AcceptLeaveRequest {
	s, _ := strconv.Atoi(status)
	var leaveStatus LeaveStatus
	if s == 0 {
		leaveStatus = 0
	} else if s == 1 {
		leaveStatus = 1
	} else {
		leaveStatus = 2
	}
	return &AcceptLeaveRequest{
		Admin:   admin,
		LeaveId: leaveID,
		Status:  leaveStatus,
	}
}

func (msg AcceptLeaveRequest) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// GetSigners Implements Msg.
func (msg AcceptLeaveRequest) GetSigners() []sdk.AccAddress {
	fromAddress, _ := sdk.AccAddressFromBech32(msg.Admin)
	return []sdk.AccAddress{fromAddress}
}

func (msg AcceptLeaveRequest) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Admin); err != nil {
		return ErrInvalidAddress
	} else if msg.Admin == "" {
		return ErrAdminAddressNil
	} else if msg.LeaveId == "" {
		return ErrStudentIdNil
	} else {
		return nil
	}
}

func NewApplyLeaveRequest(admin string, address string, reason string, leaveID string, from *time.Time, to *time.Time, status LeaveStatus) *ApplyLeaveRequest {
	return &ApplyLeaveRequest{
		Admin:   admin,
		Address: address,
		Reason:  reason,
		LeaveId: leaveID,
		From:    from,
		To:      to,
		Status:  status,
	}
}

func (msg ApplyLeaveRequest) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// GetSigners Implements Msg.
func (msg ApplyLeaveRequest) GetSigners() []sdk.AccAddress {
	fromAddress, _ := sdk.AccAddressFromBech32(msg.Admin)
	return []sdk.AccAddress{fromAddress}
}

func (msg ApplyLeaveRequest) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Address); err != nil {
		return ErrInvalidAddress
	} else if (msg.Address) == "" {
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
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// GetSigners Implements Msg.
func (msg RegisterAdminRequest) GetSigners() []sdk.AccAddress {
	fromAddress, _ := sdk.AccAddressFromBech32(msg.Address)
	return []sdk.AccAddress{fromAddress}
}

func (msg RegisterAdminRequest) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Address); err != nil {
		return ErrInvalidAddress
	} else if msg.Address == "" {
		return ErrAdminAddressNil
	} else if msg.Name == "" {
		return ErrAdminNameNil
	} else {
		return nil
	}
}
