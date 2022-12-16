package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/arki-joon-001/test/testutil/keeper"
	"github.com/arki-joon-001/test/x/test/keeper"
	"github.com/arki-joon-001/test/x/test/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.TestKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
