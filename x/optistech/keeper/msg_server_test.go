package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/hengmengsroin/optistech/testutil/keeper"
	"github.com/hengmengsroin/optistech/x/optistech/keeper"
	"github.com/hengmengsroin/optistech/x/optistech/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.OptistechKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
