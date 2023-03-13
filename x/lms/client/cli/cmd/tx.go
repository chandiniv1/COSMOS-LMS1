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
		Use:     "add-student [address] [name] [id]",
		Short:   "AddStudentCmd requests to add the student with the given details.",
		Long:    `It takes the details of admin,address,name,id in order to add a  student`,
		Example: `./lmsd tx lms add-student cosmos1et74ecw6wymvftkvfha289rxyxcxk42rfh8d89 saru 00318 --from validator-key --chain-id testnet`,
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
		Use:     "register-admin [address] [name]",
		Short:   "RegisterAdminCmd requests to register a admin with the given details",
		Long:    `It takes the address and name in order to register the admin`,
		Example: `./lmsd tx lms register-admin cosmos122ljzq3e9fucpsz3328g47zdslz0wt45nsnlp2 vitwit --from validator-key --chain-id testnet`,

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
		Use:     "apply-leave [address] [reason] [leaveID] [From] [To]",
		Short:   "ApplyLeaveCmd requests to apply a leave for the students with the given details,",
		Long:    `It gives address,reason,leaveID,from_date,to_date in order to apply leave`,
		Example: `./lmsd tx lms apply-leave 000x724 cold  00318 2023-03-08 2023-03-10 --from validator-key --chain-id testnet`,

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
			msgClient := types.NewApplyLeaveRequest(admin.String(), address, reason, leaveID, from, to, types.LeaveStatus_STATUS_PENDING)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msgClient)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func AcceptLeaveCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "accept-leave [leaveID] [status]",
		Short:   "AcceptLeaveCmd requests to accept a leave for the students with the given details,",
		Long:    `It contains the params such as admin and leaveID `,
		Example: `/lmsd tx lms accept-leave 00318 --from validator-key --chain-id testnet`,

		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				panic(err)
			}

			admin := clientCtx.GetFromAddress()
			leaveID := args[0]
			status := args[1]

			msgClient := types.NewAcceptLeaveRequest(admin.String(), leaveID, status)
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
