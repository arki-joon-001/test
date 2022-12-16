package keeper_test

import (
	"testing"

	testkeeper "github.com/arki-joon-001/test/testutil/keeper"
	"github.com/arki-joon-001/test/x/test/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.TestKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
