package cmd

import (
	"github.com/chandiniv1/COSMOS-LMS1/x/lms/types"
	"github.com/cosmos/cosmos-sdk/client"

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
		Use:   "add-student",
		Short: "A brief description of your command",
		Long:  `A longer description that spans multiple lines and likely contains example and usage of using your command.`,

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
	return cmd
}

func GetAdminCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-student",
		Short: "A brief description of your command",
		Long:  `A longer description that spans multiple lines and likely contains example and usage of using your command.`,

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
	return cmd
}

func GetStudentsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-student",
		Short: "A brief description of your command",
		Long:  `A longer description that spans multiple lines and likely contains example and usage of using your command.`,

		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				panic(err)
			}
			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.GetStudents(cmd.Context(), &types.GetStudentsRequest{})
			if err != nil {
				panic(err)
			}
			return clientCtx.PrintProto(res)
		},
	}
	return cmd
}

func GetLeaveRequestListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-student",
		Short: "A brief description of your command",
		Long:  `A longer description that spans multiple lines and likely contains example and usage of using your command.`,

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
	return cmd
}

func GetLeaveApprovesListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-student",
		Short: "A brief description of your command",
		Long:  `A longer description that spans multiple lines and likely contains example and usage of using your command.`,

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
