package keeper

import (
	"github.com/chandiniv1/cosmos-LMS/x/lms/types"
	//"github.com/chandiniv1/cosmos-LMS/x/lms/keeper/keeper.go"
	codec "github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"

	//"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

//var _ types.ApplyLeave = *types.NewApplyLeaveRequest()
//var _ types.AddStudent = *types.AddStudentRequest()

func NewKeeper(
	storeKey storetypes.StoreKey,

	cdc codec.BinaryCodec,
) Keeper {
	return Keeper{
		cdc:      cdc,
		storeKey: storeKey,
	}
}

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

func (k Keeper) AddStudent(ctx sdk.Context, addstudent types.AddStudentRequest) error {
	store := ctx.KVStore(k.storeKey)
	// key := types.StudentKey

	k.cdc.MustMarshal(&addstudent)

	marshalAddStudent, err := k.cdc.Marshal(addstudent)
	if err != nil {
		return err
	}
	store.Set(types.StudentStoreKey((addstudent.Admin), marshalAddStudent))
	//store.Set(types.StudentStoreKey((addstudent.Admin,addstudent.Id),string(marshalApplyLeave)))
}

//func check(key []byte,storetypes.KVStore,cdc codec.Marshaler)bool{
//	user1:=store.Get(key)
//	if user1==nil{
//		return false
//	}
//	var user map[string]interface{}
//	err:=cdc.UnmarshalJSON(user1,&user)
//	if err!=nil{
//		panic("err")
//	}
//	return len(user)>0
//}

// func (k Keeper) ApplyLeave(ctx sdk.Context,*types.AcceptLeaveRequest)error{
// 	store := ctx.KVStore(k.storeKey)
//     key := types.AcceptLeaveRequest.Bytes()
//     value := k.cdc.MustMarshalBinaryBare()
//     store.Set(key, value)
//     return nil
// }
