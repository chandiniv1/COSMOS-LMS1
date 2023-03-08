package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrStudentNameNil       = sdkerrors.Register(ModuleName, 1, "Student Name cant be nil")
	ErrStudentIdNil         = sdkerrors.Register(ModuleName, 2, "student id cant be nil")
	ErrStudentAddressNil    = sdkerrors.Register(ModuleName, 3, "Student address cant be nil")
	ErrStudentAlreadyExists = sdkerrors.Register(ModuleName, 4, "Student is already exist")
	ErrAdminNameNil         = sdkerrors.Register(ModuleName, 5, "admin name cant be nil")
	ErrAdminAddressNil      = sdkerrors.Register(ModuleName, 6, "Admin address cant be nil")
	ErrStudentDoesNotExist  = sdkerrors.Register(ModuleName, 7, "Student Does not exist")
	ErrAdminDoesNotExist    = sdkerrors.Register(ModuleName, 8, "Admin Does not Exist")
	ErrAdminAlreadyExists   = sdkerrors.Register(ModuleName, 9, "Admin is already exist")
	ErrStudentDatesNil      = sdkerrors.Register(ModuleName, 10, "Student Dates cant be nil")
	ErrEmptyReason          = sdkerrors.Register(ModuleName, 11, "Reason cant be empty")
	ErrLeaveNeverApplied    = sdkerrors.Register(ModuleName, 12, "Student never applied leave")
	ErrInvalidAddress       = sdkerrors.Register(ModuleName, 13, "Invalid address")
)
