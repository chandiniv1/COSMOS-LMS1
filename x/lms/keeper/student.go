package keeper

import (
	"errors"

	"github.com/chandiniv1/COSMOS-LMS1/x/lms/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

//var _ types.ApplyLeave = *types.NewApplyLeaveRequest()
//var _ types.AddStudent = *types.AddStudentRequest()

// func (k Keeper) ApplyLeave(ctx sdk.Context,applyleave types.AcceptLeaveRequest) error{
// 	store:=ctx.KVStore(k.storeKey)
// 	key:=types.LeavesKey
// 	// Leaveskey:=applyleave.GetLeaveId()
// 	marshalApplyLeave,err:=k.cdc.Marshal(applyleave)
// 	// value:= k.cdc.MustMarshalJSON(&applyleave)
// 	// //store.Set(key,value)
// 	if err!=nil{
// 		return err
// 	}
// 	store.Set(types.AdminStoreKey((applyleave.Admin,applyleave.LeaveId),string(marshalApplyLeave)))
// }

func (k Keeper) AddStdnt(ctx sdk.Context, addstudent *types.AddStudentRequest) error {

	if addstudent.Name == "" {
		return errors.New("name cant be null")
	} else if addstudent.Address == "" {
		return errors.New("address cant be null")
	} else if addstudent.Id == "" {
		return errors.New("Id cant be null")
	} else {
		store := ctx.KVStore(k.storeKey)
		// key := types.StudentKey

		k.cdc.MustMarshal(addstudent)

		marshalAddStudent, err := k.cdc.Marshal(addstudent)
		if err != nil {
			panic(err)
		}
		store.Set(types.StudentStoreKey(addstudent.Id), marshalAddStudent)
	}
	return nil
}

func (k Keeper) RgstrAdmin(ctx sdk.Context, registeradminreq *types.RegisterAdminRequest) error {
	if registeradminreq.Address == "" {
		return errors.New("Admin address cant be null")
	} else if registeradminreq.Name == "" {
		return errors.New("Admin name cant be null")
	} else {
		store := ctx.KVStore(k.storeKey)
		// key := types.StudentKey

		//k.cdc.MustMarshal(registeradminreq)

		marshalAdmin, err := k.cdc.Marshal(registeradminreq)
		if err != nil {
			return err
		}
		//store.Set(types.StudentStoreKey((addstudent.Admin), marshalAddStudent))
		store.Set(types.AdminStoreKey(registeradminreq.Address), marshalAdmin)
	}
	return nil
}

func (k Keeper) AplyLeave(ctx sdk.Context, applyleavereq *types.ApplyLeaveRequest) error {

	if applyleavereq.Address == "" {
		return errors.New("Address cant be null")
	} else if applyleavereq.Reason == "" {
		return errors.New("Reason cant be null")
	} else if applyleavereq.From == nil {
		return errors.New("From date cant be null")
	} else if applyleavereq.To == nil {
		return errors.New("To date cant be null")
	} else if k.GetStudent(ctx,applyleavereq.Address)==false{
		return errors.New("No admin present with this address")
	}else {
		store := ctx.KVStore(k.storeKey)
		// key := types.StudentKey

		//k.cdc.MustMarshal(applyleavereq)

		marshalApplyLeave, err := k.cdc.Marshal(applyleavereq)
		if err != nil {
			panic(err)
		}
		store.Set(types.AppliedLeavesStoreKey(applyleavereq.Address, applyleavereq.LeaveId), marshalApplyLeave)
	}
	return nil
}

func (k Keeper) AcptLeave(ctx sdk.Context, acceptleavereq *types.AcceptLeaveRequest) error {
	if acceptleavereq.Admin == "" {
		return errors.New("Admin cant be null")
	} else if acceptleavereq.LeaveId == "" {
		return errors.New("LeaveId cant be null")
	} else if acceptleavereq.Status == 0 {
		return errors.New("Status cant be null")
	} else if k.GetAdmin(ctx,acceptleavereq.LeaveId)==false{
		return errors.New("Admin does not exist")
	}else {
		store := ctx.KVStore(k.storeKey)
		// key := types.StudentKey

		//k.cdc.MustMarshal(applyleavereq)
		acceptleavereq.Status=types.LeaveStatus_STATUS_ACCEPTED
		marshalAcceptLeave, err := k.cdc.Marshal(acceptleavereq)
		if err != nil {
			panic(err)
		}
		store.Set(types.AcceptedLeavesStoreKey(acceptleavereq.Admin, acceptleavereq.LeaveId), marshalAcceptLeave)
	}
	return errors.New("stored successfully")
}

func (k Keeper) GetAdmin(ctx sdk.Context, adminAddress string) bool {
	store := ctx.KVStore(k.storeKey)
	adminname := store.Get(types.AdminStoreKey(adminAddress))

	if adminname == nil {
		return false
	}
	return true
}

func (k Keeper) GetStudent(ctx sdk.Context,studentID string) bool{
	store:=ctx.KVStore(k.storeKey)
	studentName:=store.Get(types.StudentStoreKey(studentID))
	if studentName==nil{
		return false
	}
	return true
}


