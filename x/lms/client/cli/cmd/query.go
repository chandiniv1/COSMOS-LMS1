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
	"github.com/chandiniv1/COSMOS-LMS1/x/lms/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"

	"github.com/spf13/cobra"
)

// queryCmd represents the query command
// var queryCmd = &cobra.Command{
// 	Use:   "query",
// 	Short: "A brief description of your command",
// 	Long: `A longer description that spans multiple lines and likely contains examples
// and usage of using your command. For example:

// Cobra is a CLI library for Go that empowers applications.
// This application is a tool to generate the needed files
// to quickly create a Cobra application.`,
// 	Run: func(cmd *cobra.Command, args []string) {
// 		fmt.Println("query called")
// 	},
// }

func GetQueryCmd() *cobra.Command {
	queryTxCmd := &cobra.Command{
		Use:   types.ModuleName,
		Short: "|lms|",
		Long:  `lms module commands`,
		RunE:  client.ValidateCmd,
	}

	queryTxCmd.AddCommand(
		GetStudentCmd(),
		GetAdminCmd(),
	)

	return queryTxCmd
}

func GetStudentCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-student",
		Short: "A brief description of your command",
		Long:  `A longer description that spans multiple lines and likely contains example and usage of using your command.`,

		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				panic(err)
			}

			id := args[0]
			address := args[1]

			msgClient := types.NewGetStudentRequest(id, address)
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msgClient)
		},
	}
	return cmd
}

func GetAdminCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-student",
		Short: "A brief description of your command",
		Long:  `A longer description that spans multiple lines and likely contains example and usage of using your command.`,

		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				panic(err)
			}

			address := args[0]

			msgClient := types.NewGetAdminRequest(address)
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msgClient)
		},
	}
	return cmd
}

func init() {
	rootCmd.AddCommand(GetStudentCmd())
	rootCmd.AddCommand(GetAdminCmd())
	// queryCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
