package keeper

import (
	"github.com/chandiniv1/COSMOS-LMS1/x/lms/types"
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

func (k Keeper) AddStd(ctx sdk.Context, addstudent *types.AddStudentRequest) error {
	
	store := ctx.KVStore(k.storeKey)
	// key := types.StudentKey

	k.cdc.MustMarshal(addstudent)

	marshalAddStudent, err := k.cdc.Marshal(addstudent)
	if err != nil {
		return err
	}
	//store.Set(types.StudentStoreKey((addstudent.Admin), marshalAddStudent))
	//store.Set(types.StudentStoreKey(addstudent.Address),marshalAddStudent)
	store.Set(types.StudentStoreKey(addstudent.Id), marshalAddStudent)
	return nil
}

func (k Keeper) RegisterAdmin(ctx sdk.Context, registeradminreq *types.RegisterAdminRequest) error {
	
	store := ctx.KVStore(k.storeKey)
	// key := types.StudentKey

	//k.cdc.MustMarshal(registeradminreq)

	marshalAdmin, err := k.cdc.Marshal(registeradminreq)
	if err != nil {
		return err
	}
	//store.Set(types.StudentStoreKey((addstudent.Admin), marshalAddStudent))
	//store.Set(types.StudentStoreKey(addstudent.Address),marshalAddStudent)
	store.Set(types.AdminStoreKey(registeradminreq.Address), marshalAdmin)
	return nil
}

//func (r *types.AddStudentRequest) Validate() error {
	// if err := validateName(r.Name); err != nil {
	// 	return err
	// }
	// if err := validateSigVerifyCostED25519(p.SigVerifyCostED25519); err != nil {
	// 	return err
	// }
	// if err := validateSigVerifyCostSecp256k1(p.SigVerifyCostSecp256k1); err != nil {
	// 	return err
	// }
	// if err := validateMaxMemoCharacters(p.MaxMemoCharacters); err != nil {
	// 	return err
	// }
	// if err := validateTxSizeCostPerByte(p.TxSizeCostPerByte); err != nil {
	// 	return err
	// }

//	return nil
//}

// func validateName(i interface{}) error {
// 	v, ok := i.(uint64)
// 	if !ok {
// 		return fmt.Errorf("invalid parameter type: %T", i)
// 	}

// 	if v == 0 {
// 		return fmt.Errorf("invalid max memo characters: %d", v)
// 	}

// 	return nil
// }
//func GetStudent(ctx sdk.Context,)

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
