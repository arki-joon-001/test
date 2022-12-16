package test_test

import (
	"testing"

	keepertest "github.com/arki-joon-001/test/testutil/keeper"
	"github.com/arki-joon-001/test/testutil/nullify"
	"github.com/arki-joon-001/test/x/test"
	"github.com/arki-joon-001/test/x/test/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.TestKeeper(t)
	test.InitGenesis(ctx, *k, genesisState)
	got := test.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
