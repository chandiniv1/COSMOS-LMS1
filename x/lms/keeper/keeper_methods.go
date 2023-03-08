package keeper

import (
	"fmt"
	"strconv"

	"github.com/chandiniv1/COSMOS-LMS1/x/lms/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

//<-------------------------------ADD STUDENT------------------------------------>//

//this method registers student to the students store and return error if the fields are not properly specified
//or if the student already exist

func (k Keeper) AddStdnt(ctx sdk.Context, addStudent *types.AddStudentRequest) error {
	if addStudent.Name == "" {
		return types.ErrStudentNameNil
	} else if addStudent.Address == "" {
		return types.ErrStudentAddressNil
	} else if addStudent.Id == "" {
		return types.ErrStudentIdNil
	} else if addStudent.Admin == "" {
		return types.ErrAdminAddressNil
	} else {
		store := ctx.KVStore(k.storeKey)
		marshalAddStudent, err := k.cdc.Marshal(addStudent)
		if err != nil {
			panic(err)
		} else {
			if k.CheckStudent(ctx, addStudent.Address) != false {
				return types.ErrStudentAlreadyExists
			}
			store.Set(types.StudentStoreKey(addStudent.Address), marshalAddStudent)
		}
	}
	return nil
}

//<-------------------------------Register Admin---------------------------------->//

//This method is used to register the admin and it can be retrieved through AdminStoreKey

func (k Keeper) RgstrAdmin(ctx sdk.Context, registerAdminReq *types.RegisterAdminRequest) error {
	if registerAdminReq.Address == "" {
		return types.ErrAdminAddressNil
	} else if registerAdminReq.Name == "" {
		return types.ErrAdminNameNil
	} else {
		store := ctx.KVStore(k.storeKey)

		marshalAdmin, err := k.cdc.Marshal(registerAdminReq)
		if err != nil {
			return err
		} else {
			if k.CheckAdmin(ctx, registerAdminReq.Address) != false {
				return types.ErrAdminAlreadyExists
			}
			store.Set(types.AdminStoreKey(registerAdminReq.Address), marshalAdmin)
		}
	}
	return nil
}

//<-------------------------------APPLY LEAVE-------------------------------------->//

//This methtod is used to apply the leave by a student ,it checks whether the student is existed in the store or not after that
//it adds that leave request to the AppliedLeavesStore as well as it updates the counter

func (k Keeper) AplyLeave(ctx sdk.Context, applyleavereq *types.ApplyLeaveRequest) error {

	if applyleavereq.Address == "" {
		return types.ErrStudentAddressNil
	} else if applyleavereq.Reason == "" {
		return types.ErrEmptyReason
	} else if applyleavereq.From == nil {
		return types.ErrStudentDatesNil
	} else if applyleavereq.To == nil {
		return types.ErrStudentDatesNil
	} else {
		store := ctx.KVStore(k.storeKey)
		marshalApplyLeave, err := k.cdc.Marshal(applyleavereq)
		if err != nil {
			panic(err)
		}
		addr := types.LeavesCounterKey(applyleavereq.LeaveId)
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
		store.Set(types.LeavesStoreKey(applyleavereq.LeaveId), marshalApplyLeave)
	}
	return nil
}

//<-------------------------------ACCEPT LEAVE------------------------------------>//

//this methods is accepts the leaves if all fields are valid and rejects the leaves if any of the fileds
//are not valid

func (k Keeper) AcptLeave(ctx sdk.Context, acceptleavereq *types.AcceptLeaveRequest) error {
	store := ctx.KVStore(k.storeKey)
	//if the admin is nil then the leave status is undefined
	marshalAcceptLeave, err := k.cdc.Marshal(acceptleavereq)
	if err != nil {
		panic(err)
	}
	store.Set(types.AcceptedLeavesStoreKey(acceptleavereq.Admin, acceptleavereq.LeaveId), marshalAcceptLeave)

	c := store.Get(types.LeavesCounterKey(acceptleavereq.LeaveId))
	count := string(c)
	l := store.Get(types.AppliedLeavesStoreKey(acceptleavereq.LeaveId, count))
	var leave types.ApplyLeaveRequest
	k.cdc.Unmarshal(l, &leave)
	leave.Status = acceptleavereq.Status
	leaveMarhsal, _ := k.cdc.Marshal(&leave)
	store.Set(types.AppliedLeavesStoreKey(acceptleavereq.LeaveId, count), leaveMarhsal)
	return nil
}

//<-------------------------CHECK ADMIN--------------------------->//

//This method takes the address as parameter and check whether the admin with that address is present or not

func (k Keeper) CheckAdmin(ctx sdk.Context, address string) bool {
	store := ctx.KVStore(k.storeKey)
	admin := store.Get(types.AdminStoreKey(address))

	if admin == nil {
		return false
	}
	return true
}

//<----------------------CHECK STUDENT--------------------------->//

//This method takes the address as parameter and check whether the student with that address is present or not

func (k Keeper) CheckStudent(ctx sdk.Context, address string) bool {
	store := ctx.KVStore(k.storeKey)
	student := store.Get(types.StudentStoreKey(address))
	if student == nil {
		return false
	}
	return true
}

//<---------------------------GET STUDENT---------------------------->//

//this methods is used to get a student  by taking address as parameter

func (k Keeper) GetStdnt(ctx sdk.Context, address string) (req types.AddStudentRequest, err error) {
	store := ctx.KVStore(k.storeKey)
	student := store.Get(types.StudentStoreKey(address))
	if student == nil {
		panic("student not found")
	}
	fmt.Println(student)
	k.cdc.MustUnmarshal(student, &req)
	fmt.Println(req)
	return req, err
}

//<---------------GET ADMIN------------------------------------>//

//This method is used to get Admin by taking address as parameter

func (k Keeper) GetAdmn(ctx sdk.Context, address string) (req types.RegisterAdminRequest, err error) {
	if _, err := sdk.AccAddressFromBech32(address); err != nil {
		panic(err)
	}
	store := ctx.KVStore(k.storeKey)
	admin := store.Get(types.AdminStoreKey(address))
	if admin == nil {
		panic("admin not found")
	}
	k.cdc.MustUnmarshal(admin, &req)
	fmt.Println(req)
	return req, err
}

//<-------------GET STUDENTS------------------------------------->//

//This method is used to return all the  students that are present in the students store

func (k Keeper) GetStdnts(ctx sdk.Context, getStudents *types.GetStudentsRequest) []*types.AddStudentRequest {
	store := ctx.KVStore(k.storeKey)
	var students []*types.AddStudentRequest
	iter := sdk.KVStorePrefixIterator(store, types.StudentKey)
	for ; iter.Valid(); iter.Next() {
		var student types.AddStudentRequest
		k.cdc.Unmarshal(iter.Value(), &student)
		students = append(students, &student)
		fmt.Println("the student are", student)
	}
	return students
}

//<---------------------------GET LEAVE REQUESTS------------------------------->//

//This method is used to get all the leave Request from the applied  leaves store

func (k Keeper) GetLeaveRqsts(ctx sdk.Context, getLeaves *types.GetLeaveRequestsRequest) []*types.ApplyLeaveRequest {
	store := ctx.KVStore(k.storeKey)
	var leavereqs []*types.ApplyLeaveRequest
	iter := sdk.KVStorePrefixIterator(store, types.AppliedLeavesKey)
	for ; iter.Valid(); iter.Next() {
		var leaves types.ApplyLeaveRequest
		k.cdc.Unmarshal(iter.Value(), &leaves)
		leavereqs = append(leavereqs, &leaves)
		fmt.Println("students who applied for leaves : ", leaves)
	}
	return leavereqs
}

//<--------------------------GET APPROVED lEAVES------------------------------>//

//This method is used to Get the Status of the leaves from applied leaves store

func (k Keeper) GetLeavesStatus(ctx sdk.Context, getLeaves *types.GetLeaveApprovesRequest) []*types.AcceptLeaveRequest {
	store := ctx.KVStore(k.storeKey)
	var status []*types.AcceptLeaveRequest
	iter := sdk.KVStorePrefixIterator(store, types.AcceptedLeavesKey)
	for ; iter.Valid(); iter.Next() {
		var approve types.AcceptLeaveRequest
		k.cdc.Unmarshal(iter.Value(), &approve)
		status = append(status, &approve)
		fmt.Println("the students whose leaves are approved", approve)
	}
	return status
}

func (k Keeper) GetStatusByID(ctx sdk.Context, address string, leaveID string) (req types.AcceptLeaveRequest, err error) {
	store := ctx.KVStore(k.storeKey)
	student := store.Get(types.AcceptedLeavesStoreKey(address, leaveID))
	if student == nil {
		panic("student  not found")
	}
	k.cdc.MustUnmarshal(student, &req)
	fmt.Println(req)
	return req, err
}

func (k Keeper) GetLeaveReqByID(ctx sdk.Context, leaveID string) (req types.ApplyLeaveRequest, err error) {
	store := ctx.KVStore(k.storeKey)
	leave := store.Get(types.LeavesStoreKey(leaveID))
	if leave == nil {
		panic("student leave request not found")
	}
	k.cdc.MustUnmarshal(leave, &req)
	fmt.Println(req)
	return req, err
}
