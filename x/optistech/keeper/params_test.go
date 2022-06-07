package keeper_test

import (
	"testing"

	testkeeper "github.com/hengmengsroin/optistech/testutil/keeper"
	"github.com/hengmengsroin/optistech/x/optistech/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.OptistechKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
