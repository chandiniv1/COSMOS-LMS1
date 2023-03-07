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
		GetStatusCmd(),
		GetLeaveReqCmd(),
	)
	return queryTxCmd
}

func GetStudentCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-student",
		Short: "|Address|",
		Long:  `get-student returns the students by taking ID and Address`,

		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				panic(err)
			}

			student := types.GetStudentRequest{
				Address: args[0],
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
		Short: "|Address|",
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

func GetStatusCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-status",
		Short: "|admin|leaveID",
		Long:  `get-status returns the leave status of a student by taking ID`,

		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				panic(err)
			}

			status := types.GetStatusRequest{
				Admin:   args[0],
				LeaveID: args[1],
			}

			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.GetStatus(cmd.Context(), &status)

			if err != nil {
				panic(err)
			}

			return clientCtx.PrintProto(res)

		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func GetLeaveReqCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-leave-request",
		Short: "|leaveID|",
		Long:  `get-leave-request is used to get the leave request of a student by taking leaveID`,

		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				panic(err)
			}
			leaveReq := types.GetLeaveRequest{
				LeaveID: args[0],
			}
			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.GetLeaveReq(cmd.Context(), &leaveReq)

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
			res, err := queryClient.GetLeaveStatus(cmd.Context(), &types.GetLeaveApprovesRequest{})
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
	rootCmd.AddCommand(GetStatusCmd())
	rootCmd.AddCommand(GetLeaveReqCmd())
}
