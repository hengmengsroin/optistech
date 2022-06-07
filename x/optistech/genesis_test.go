package optistech_test

import (
	"testing"

	keepertest "github.com/hengmengsroin/optistech/testutil/keeper"
	"github.com/hengmengsroin/optistech/testutil/nullify"
	"github.com/hengmengsroin/optistech/x/optistech"
	"github.com/hengmengsroin/optistech/x/optistech/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.OptistechKeeper(t)
	optistech.InitGenesis(ctx, *k, genesisState)
	got := optistech.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
