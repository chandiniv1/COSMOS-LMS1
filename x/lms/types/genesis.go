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

// func InitGenesis()
