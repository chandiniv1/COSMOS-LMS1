package keeper

import (
	"errors"

	"fmt"
	"strconv"

	"github.com/chandiniv1/COSMOS-LMS1/x/lms/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

//<-------------------ADD STUDENT-------------------------------------------->//

func (k Keeper) AddStdnt(ctx sdk.Context, addstudent *types.AddStudentRequest) error {
	if addstudent.Name == "" {
		return types.ErrStudentNameNil
	} else if addstudent.Address == "" {
		return types.ErrStudentAddressNil
	} else if addstudent.Id == "" {
		return types.ErrStudentIdNil
	} else {
		store := ctx.KVStore(k.storeKey)
		marshalAddStudent, err := k.cdc.Marshal(addstudent)
		if err != nil {
			panic(err)
		} else {
			if k.CheckStudent(ctx, addstudent.Address) != false {
				return types.ErrStudentAlreadyExists
			}
			store.Set(types.StudentStoreKey(addstudent.Id), marshalAddStudent)
		}
	}
	return nil
}
//<----------------------Register Admin---------------------------------->//

func (k Keeper) RgstrAdmin(ctx sdk.Context, registeradminreq *types.RegisterAdminRequest) error {
	if registeradminreq.Address == "" {
		return types.ErrAdminAddressNil
	} else if registeradminreq.Name == "" {
		return types.ErrAdminNameNil
	} else {
		store := ctx.KVStore(k.storeKey)

		marshalAdmin, err := k.cdc.Marshal(registeradminreq)
		if err != nil {
			return err
		} else {
			if k.CheckAdmin(ctx, registeradminreq.Address) != false {
				return types.ErrAdminAlreadyExists
			}
			store.Set(types.AdminStoreKey(registeradminreq.Address), marshalAdmin)
		}
	}
	return nil
}

//<-----------------APPLY LEAVE-------------------------------------->//

func (k Keeper) AplyLeave(ctx sdk.Context, applyleavereq *types.ApplyLeaveRequest) error {

	if applyleavereq.Address == "" {
		return types.ErrStudentAddressNil
	} else if applyleavereq.Reason == "" {
		return types.ErrEmptyReason
	} else if applyleavereq.From == nil {
		return types.ErrStudentDatesNil
	} else if applyleavereq.To == nil {
		return types.ErrStudentDatesNil
	} else if k.CheckStudent(ctx, applyleavereq.Address) == false {
		return types.ErrStudentDoesNotExist
	} else {
		store := ctx.KVStore(k.storeKey)
		marshalApplyLeave, err := k.cdc.Marshal(applyleavereq)
		if err != nil {
			panic(err)
		}
		addr := types.LeavesCounterKey(sdk.AccAddress(string(applyleavereq.Address)).String())
		counter := store.Get(addr)
		if counter == nil {
			store.Set(addr, []byte("1"))
		} else {
			c, err := strconv.Atoi(string(counter))
			if err != nil {
				panic(err)
			}
			c = c + 1
			store.Set(addr, []byte(fmt.Sprint(c)))
		}
		counter = store.Get(addr)
		store.Set(types.AppliedLeavesStoreKey(applyleavereq.LeaveId, string(counter)), marshalApplyLeave)
		//store.Set(types.AppliedLeavesStoreKey(applyleavereq.Address, applyleavereq.LeaveId), marshalApplyLeave)
	}
	return nil
}

//<-----------------ACCEPT LEAVE------------------------------------>//

func (k Keeper) AcptLeave(ctx sdk.Context, acceptleavereq *types.AcceptLeaveRequest) error {
	if acceptleavereq.Admin == "" {
		return types.ErrAdminNameNil
	} else if acceptleavereq.LeaveId == "" {
		return types.ErrStudentIdNil
	} else if k.CheckAdmin(ctx, acceptleavereq.Admin) == false {
		return types.ErrAdminDoesNotExist
	} else {
		store := ctx.KVStore(k.storeKey)
		acceptleavereq.Status = types.LeaveStatus_STATUS_ACCEPTED
		marshalAcceptLeave, err := k.cdc.Marshal(acceptleavereq)
		if err != nil {
			panic(err)
		}
		store.Set(types.AcceptedLeavesStoreKey(acceptleavereq.Admin, acceptleavereq.LeaveId), marshalAcceptLeave)
	}
	return errors.New("stored successfully")
}

//<---------------CHECK ADMIN------------------------------------->//

func (k Keeper) CheckAdmin(ctx sdk.Context, address string) bool {
	store := ctx.KVStore(k.storeKey)
	admin := store.Get(types.AdminStoreKey(address))

	if admin == nil {
		return false
	}
	return true
}

//<-----------------CHECK STUDENT------------------------------>//

func (k Keeper) CheckStudent(ctx sdk.Context, studentAddress string) bool {
	store := ctx.KVStore(k.storeKey)
	studentName := store.Get(types.StudentStoreKey(studentAddress))
	//fmt.Println("studentname:",studentName)
	if studentName == nil {
		return false
	}
	return true
}

//<-----------------GET STUDENT------------------------------->//

func (k Keeper) GetStdnt(ctx sdk.Context, address string) (req types.Student, err error) {
	store := ctx.KVStore(k.storeKey)
	student := store.Get(types.StudentStoreKey(address))
	if student == nil {
		//errors.New("student not found")
		panic("student not found")
	}
	k.cdc.MustUnmarshal(student, &req)
	return req, err
}

//<---------------GET ADMIN------------------------------------>//

func (k Keeper) GetAdmn(ctx sdk.Context, address string) []byte {
	if _, err := sdk.AccAddressFromBech32(address); err != nil {
		panic(err)
	}
	store := ctx.KVStore(k.storeKey)
	return store.Get(types.AdminStoreKey(address))

}

//<-------------GET STUDENTS------------------------------------->//

func (k Keeper) GetStdnts(ctx sdk.Context, getStudents *types.GetStudentsRequest) {
	store := ctx.KVStore(k.storeKey)
	var student types.Student
	iter := store.Iterator(types.StudentKey, nil)
	for ; iter.Valid(); iter.Next() {
		k.cdc.Unmarshal(iter.Value(), &student)
		fmt.Println(student)
	}
}

//<------------------GET LEAVE REQUESTS------------------------------->//

func (k Keeper) GetLeaveRqsts(ctx sdk.Context, getLeaves *types.GetLeaveRequestsRequest) {
	store := ctx.KVStore(k.storeKey)
	var leaves types.ApplyLeaveRequest
	iter := store.Iterator(types.AppliedLeavesKey, nil)
	for ; iter.Valid(); iter.Next() {
		k.cdc.Unmarshal(iter.Value(), &leaves)
		fmt.Println(leaves)
	}
}

//<--------------GET APPROVED lEAVES------------------------------->//

func (k Keeper) GetAcceptedLeaves(ctx sdk.Context, getLeaves *types.GetLeaveApprovesRequest) {
	store := ctx.KVStore(k.storeKey)
	var approves types.AcceptLeaveRequest
	iter := store.Iterator(types.AcceptedLeavesKey, nil)
	for ; iter.Valid(); iter.Next() {
		k.cdc.Unmarshal(iter.Value(), &approves)
		fmt.Println(approves)
	}
}

// func (k Keeper) AcceptLeaves(ctx sdk.Context, req *types.AcceptLeaveRequest) error {
// 	if _, err := sdk.AccAddressFromBech32(req.Admin); err != nil {
// 		panic(fmt.Errorf("invalid address: %w", err))
// 	}
// 	store := ctx.KVStore(k.storeKey)
// 	adminpresent := store.Get(types.AdminStoreKey(req.Admin))
// 	if adminpresent == nil {
// 		req.Status = types.LeaveStatus_STATUS_UNDEFINED
// 		_, err := k.cdc.Marshal(req)
// 		if err != nil {
// 			panic(err)
// 		} else {
// 			val := store.Get(types.StudentStoreKey(req.LeaveId))
// 			if val == nil {
// 				return nil//types.ErrStudentDidNotLogin 
// 			}
// 			leaveval := store.Get(types.AppliedLeavesStoreKey(req.,))
// 			if leaveval == nil {
// 				fmt.Println("student did not request the leave")
// 			}
// 		}
// 	} else {
// 		req.Status = types.LeaveStatus_STATUS_ACCEPTED
// 		marshalaccepteddata, err := k.cdc.Marshal(req)
// 		if err != nil {
// 			log.Println(err)
// 		}
// 		marshaldata := store.Get(types.LeaveStoreId(req.StudentId))
// 		if marshaldata == nil {
// 			fmt.Println("student did not request leave")
// 		} else {
// 			store.Set(types.AllLeavesStoreId(req.StudentId), marshalaccepteddata)
// 		}
// 	}
// 	return nil
// }