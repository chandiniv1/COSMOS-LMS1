package cmd

import (
	"fmt"

	"github.com/chandiniv1/COSMOS-LMS1/x/lms/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"

	"github.com/spf13/cobra"
)

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
		GetStudentsCmd(),
		GetLeaveRequestListCmd(),
		GetLeaveApprovesListCmd(),
	)
	return queryTxCmd
}

func GetStudentCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-student",
		Short: "get-student returns the students by taking ID and Address",
		Long:  `get-student returns the students by taking ID and Address`,

		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				panic(err)
			}

			student := types.GetStudentRequest{
				Id:      args[0],
				Address: args[1],
			}

			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.GetStudent(cmd.Context(), &student)

			if err != nil {
				panic(err)
			}

			return clientCtx.PrintProto(res)

		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func GetAdminCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-admin",
		Short: "get-admin is used to get the admin by taking admin address",
		Long:  `get-admin is used to get the admin by taking admin address`,

		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				panic(err)
			}
			admin := types.GetAdminRequest{
				Address: args[0],
			}
			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.GetAdmin(cmd.Context(), &admin)

			if err != nil {
				panic(err)
			}
			return clientCtx.PrintProto(res)

		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func GetStudentsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-students",
		Short: "get-students returns all the students",
		Long:  `get-students returns all the students present in the student store`,

		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				panic(err)
			}
			queryClient := types.NewQueryClient(clientCtx)
			fmt.Println("Im inside the getstudents cli")
			res, err := queryClient.GetStudents(cmd.Context(), &types.GetStudentsRequest{})
			if err != nil {
				panic(err)
			}
			// fmt.Println("the resulted received is ", res.Students)
			return clientCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func GetLeaveRequestListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-leave-requests",
		Short: "get-leave-requests returns all the students who applied for leaves",
		Long:  `get-leave-requests returns all the students who applied for leaves in the applied leaves store`,

		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				panic(err)
			}
			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.GetLeaveRequests(cmd.Context(), &types.GetLeaveRequestsRequest{})
			if err != nil {
				panic(err)
			}
			return clientCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func GetLeaveApprovesListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-leave-approves",
		Short: "get-leave-approves returns all the students whose leaves are approved",
		Long:  `get-leave-approves returns all the students whose leaves are approved in the accepted leaves store`,

		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				panic(err)
			}
			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.GetLeaveApproves(cmd.Context(), &types.GetLeaveApprovesRequest{})
			if err != nil {
				panic(err)
			}
			return clientCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func init() {
	rootCmd.AddCommand(GetStudentCmd())
	rootCmd.AddCommand(GetAdminCmd())
	rootCmd.AddCommand(GetStudentsCmd())
	rootCmd.AddCommand(GetLeaveRequestListCmd())
	rootCmd.AddCommand(GetLeaveApprovesListCmd())

	// queryCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
