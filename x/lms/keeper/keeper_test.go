package keeper_test

import (
	"fmt"
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
)

type TestSuite struct {
	suite.Suite
	ctx         sdk.Context
	stdntKeeper keeper.Keeper
}

func (s *TestSuite) SetupTest() {

	fmt.Println("I am in setup")

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

}

func TestTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
