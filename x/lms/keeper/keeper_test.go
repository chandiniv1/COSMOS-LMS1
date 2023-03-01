package keeper_test

import (
	"testing"

	"github.com/chandiniv1/COSMOS-LMS1/x/lms/keeper"
	"github.com/chandiniv1/COSMOS-LMS1/x/lms/types"
	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/cosmos/cosmos-sdk/store"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	"github.com/cosmos/cosmos-sdk/testutil"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"
	dbm "github.com/tendermint/tm-db"

	//"google/protobuf/timestamp.proto"
	"time"
)

type TestSuite struct {
	suite.Suite
	ctx         sdk.Context
	stdntKeeper keeper.Keeper
}

func (s *TestSuite) SetupTest() {

	db := dbm.NewMemDB()
	cms := store.NewCommitMultiStore(db)

	encCfg := simapp.MakeTestEncodingConfig()
	lmsKey := sdk.NewKVStoreKey(types.StoreKey)
	ctx := testutil.DefaultContext(lmsKey, sdk.NewTransientStoreKey("transient_test"))
	keeper := keeper.NewKeeper(lmsKey, encCfg.Codec)

	cms.MountStoreWithDB(lmsKey, storetypes.StoreTypeIAVL, db)
	s.Require().NoError(cms.LoadLatestVersion())

	s.stdntKeeper = keeper
	s.ctx = ctx

}

func (s *TestSuite) TestAddStudent() {
	// ss := nil
	tests := []struct {
		Address string
		Admin   string
		Name    string
		Id      string
		Res     error
	}{
		{"00012", "mango", "", "8723", types.ErrStudentNameNil},
		{"00003", "mango", "apple", "", types.ErrStudentIdNil},
		{"", "mango", "apple", "0003", types.ErrStudentAddressNil},
	}
	for _, test := range tests {
		err := s.stdntKeeper.AddStdnt(s.ctx, &types.AddStudentRequest{
			Address: test.Address,
			Admin:   test.Admin,
			Name:    test.Name,
			Id:      test.Id,
		})
		s.Require().Equal(err, test.Res)
	}
	err := s.stdntKeeper.AddStdnt(s.ctx, &types.AddStudentRequest{
		Address: "0001",
		Admin:   "vitwit",
		Name:    "apple",
		Id:      "1001",
	})
	s.Require().NoError(err)

	// details := s.stdntKeeper.AddStdnt(s.ctx, &types.AddStudentRequest{
	// 	Admin:   "vitwit",
	// 	Address: "0002",
	// 	Name:    "",
	// 	Id:      "1002",
	// })
	// expected := "name cant be null"
	// s.Require().EqualError(details, expected)
}

func (s *TestSuite) TestRegisterAdmin() {
	tests := []struct {
		Address string
		Name    string
		Res     error
	}{
		{"", "vitwit", types.ErrAdminAddressNil},
		{"00x02", "", types.ErrAdminNameNil},
	}

	for _, test := range tests {
		err := s.stdntKeeper.RgstrAdmin(s.ctx, &types.RegisterAdminRequest{
			Address: test.Address,
			Name:    test.Name,
		})
		s.Require().Equal(err, test.Res)
	}
	err := s.stdntKeeper.RgstrAdmin(s.ctx, &types.RegisterAdminRequest{
		Address: "00x01",
		Name:    "vitwit",
	})
	s.Require().NoError(err)
}

func (s *TestSuite) TestApplyLeave() {
	// type test struct {
	// 	args1    types.ApplyLeaveRequest
	// 	expected error
	// }
	// date := "2006-Jan-02"
	// fromdate, _ := time.Parse(date, "2023-Feb-23")
	// todate, _ := time.Parse(date, "2023-Feb-24")
	// var testcases = []test{
	// 	{
	// 		args1: types.ApplyLeaveRequest{
	// 			Address: "00x01",
	// 			Reason:  "cold",
	// 			LeaveId: "0001",
	// 			From:    &fromdate,
	// 			To:      &todate,
	// 		},
	// 		expected: "",
	// 	},
	// }
	// for _, test := range testcases {
	// 	if err := s.stdntKeeper.AplyLeave(s.ctx, &test.args1); err != nil {
	// 		s.Require().Equal(test.expected, "")
	// 	}
	// }
	date := "2006-Jan-02"
	fromdate2, _ := time.Parse(date, "2023-Feb-23")
	todate2, _ := time.Parse(date, "2023-Feb-24")

	err := s.stdntKeeper.AplyLeave(s.ctx, &types.ApplyLeaveRequest{
		Address: "00x01",
		Reason:  "cold",
		LeaveId: "1001",
		From:    &fromdate2,
		To:      &todate2,
	})
	s.Require().NoError(err)
}

func (s *TestSuite) TestAcceptLeave() {
	tests := []struct {
		Admin   string
		LeaveID string
		Status  types.LeaveStatus
		res     error
	}{
		{"vitwit", "", types.LeaveStatus_STATUS_ACCEPTED, types.ErrStudentIdNil},
		{"", "0001", types.LeaveStatus_STATUS_ACCEPTED, types.ErrAdminNameNil},
		{"sita", "0001", types.LeaveStatus_STATUS_ACCEPTED, types.ErrAdminDoesNotExist},
	}
	for _, test := range tests {
		err := s.stdntKeeper.AcptLeave(s.ctx, &types.AcceptLeaveRequest{
			Admin:   test.Admin,
			LeaveId: test.LeaveID,
			Status:  test.Status,
		})
		s.Require().Equal(err, test.res)
	}

}

func TestTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
