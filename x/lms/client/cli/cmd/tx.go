package cmd

import (
	"time"
	"github.com/chandiniv1/COSMOS-LMS1/x/lms/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
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
		ApplyLeaveCmd(),
	)
	return studentTxCmd
}

func AddStudentCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-student",
		Short: "|address|name|id",
		Long:  `AddStudentCmd requests to add the student with the given details.It gives the details of admin,address,name,id in order to add a  student`,

		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				panic(err)
			}

			admin := clientCtx.GetFromAddress()
			address := args[0]
			name := args[1]
			id := args[2]

			msgClient := types.NewAddStudentRequest(admin.String(), address, name, id)
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msgClient)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func RegisterAdminCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "register-admin",
		Short: "|address|name",
		Long:  `RegisterAdminCmd requests to register a admin with the given details,It gives address and name in order to register the admin`,

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
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func ApplyLeaveCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "apply-leave",
		Short: "",
		Long:  `ApplyLeaveCmd requests to apply a leave for the students with the given details,It gives address,reason,leaveID,from_date,to_date in order to apply leave`,

		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				panic(err)
			}
			var format string = "2006-Jan-06"
			fromDate, _ := time.Parse(format, args[3])
			toDate, _ := time.Parse(format, args[4])

			admin := clientCtx.GetFromAddress()
			address := args[0]
			reason := args[1]
			leaveID := args[2]
			from := &fromDate
			to := &toDate
			//to:=args[4]
			msgClient := types.NewApplyLeaveRequest(admin.String(), address, reason, leaveID, from, to)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msgClient)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func AcceptLeaveCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "accept-leave",
		Short: "|admin|leaveID|",
		Long:  `AcceptLeaveCmd requests to accept a leave for the students with the given details,It requests with the params such as admin and leaveID `,

		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				panic(err)
			}

			admin := clientCtx.GetFromAddress()

			leaveID := args[0]
			//status:=args[2]

			msgClient := types.NewAcceptLeaveRequest(admin.String(), leaveID, types.LeaveStatus_STATUS_ACCEPTED)
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msgClient)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func init() {
	rootCmd.AddCommand(AddStudentCmd())
	rootCmd.AddCommand(RegisterAdminCmd())
	rootCmd.AddCommand(ApplyLeaveCmd())
	rootCmd.AddCommand(AcceptLeaveCmd())

}
