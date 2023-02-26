// Copyright © 2023 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"time"

	"github.com/chandiniv1/COSMOS-LMS1/x/lms/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	studentTxCmd := &cobra.Command{
		Use:   types.ModuleName,
		Short: "|lms|",
		Long:  `lms module commands`,
		RunE:  client.ValidateCmd,
	}

	studentTxCmd.AddCommand(
		RegisterAdminCmd(),
		AddStudentCmd(),
		AcceptLeaveCmd(),
		//NewCmdApplyLeave(),
	)

	return studentTxCmd
}

func AddStudentCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-user",
		Short: "A brief description of your command",
		Long:  `A longer description that spans multiple lines and likely contains example and usage of using your command.`,

		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				panic(err)
			}

			admin := args[0]
			address := args[1]
			name := args[2]
			id := args[3]

			msgClient := types.NewAddStudentRequest(admin, address, name, id)
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msgClient)
		},
	}
	return cmd
}

func RegisterAdminCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-user",
		Short: "A brief description of your command",
		Long:  `A longer description that spans multiple lines and likely contains example and usage of using your command.`,

		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				panic(err)
			}

			address := args[0]
			name := args[1]

			msgClient := types.NewRegisterAdminRequest(address, name)
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msgClient)
		},
	}
	return cmd
}

func AcceptLeaveCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-user",
		Short: "A brief description of your command",
		Long:  `A longer description that spans multiple lines and likely contains example and usage of using your command.`,

		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				panic(err)
			}
			var format string = "2006-Jan-06"
			fromDate, _ := time.Parse(format, args[3])
			toDate, _ := time.Parse(format, args[4])

			address := args[0]
			reason := args[1]
			leaveID := args[2]
			from := &fromDate
			to := &toDate
			//to:=args[4]

			msgClient := types.NewApplyLeaveRequest(address, reason, leaveID, from, to)
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msgClient)
		},
	}
	return cmd
}

func init() {
	rootCmd.AddCommand(AddStudentCmd())
	rootCmd.AddCommand(RegisterAdminCmd())

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// txCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// txCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
