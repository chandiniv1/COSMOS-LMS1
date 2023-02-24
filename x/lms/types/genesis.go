package types

// import (
// 	// cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
// )

//NewGenesisState creates new GenesisState object
// func NewGenesisState(responses []*AcceptLeaveRequest) *GenesisState {
// 	return &GenesisState{
// 		admin:
// 	}
// }

// // ValidateGenesis check the given genesis state has no integrity issues
func ValidateGenesis(data *GenesisState) error {
	return nil
}

// // DefaultGenesisState - Return a default genesis state
func DefaultGenesisState() *GenesisState {
	return &GenesisState{}
}

// var _ cdctypes.UnpackInterfacesMessage = GenesisState{}

// // UnpackInterfaces implements UnpackInterfacesMessage.UnpackInterfaces
// func (data GenesisState) UnpackInterfaces(unpacker cdctypes.AnyUnpacker) error {
// 	for _, a := range data.Student {
// 		err := a.UnpackInterfaces(unpacker)
// 		if err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }

// // UnpackInterfaces implements UnpackInterfacesMessage.UnpackInterfaces
// func (msg GrantAuthorization) UnpackInterfaces(unpacker cdctypes.AnyUnpacker) error {
// 	var a Authorization
// 	return unpacker.UnpackAny(msg.Authorization, &a)
// }

// func ValidateGenesis(data GenesisState) error {
// 	if err := data.; err != nil {
// 		return err
// 	}

	
// }