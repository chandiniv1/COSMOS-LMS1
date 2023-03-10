package main

import (
	"os"

	"github.com/cosmos/cosmos-sdk/server"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"

	//"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/chandiniv1/COSMOS-LMS1/simapp"
	"github.com/chandiniv1/COSMOS-LMS1/simapp/lmsd/cmd"

	//"github.com/chandiniv1/COSMOS-LMS1/simapp/simd/cmd"
	//"github.com/cosmos/cosmos-sdk/simapp/simd/cmd"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd()

	if err := svrcmd.Execute(rootCmd, "", simapp.DefaultNodeHome); err != nil {
		switch e := err.(type) {
		case server.ErrorCode:
			os.Exit(e.Code)

		default:
			os.Exit(1)
		}
	}
}
