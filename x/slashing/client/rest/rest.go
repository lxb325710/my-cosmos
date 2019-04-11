package rest

import (
	"github.com/gorilla/mux"

	"my-cosmos/cosmos-sdk/client/context"
	"my-cosmos/cosmos-sdk/codec"
	"my-cosmos/cosmos-sdk/crypto/keys"
)

// RegisterRoutes registers staking-related REST handlers to a router
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router, cdc *codec.Codec, kb keys.Keybase) {
	registerQueryRoutes(cliCtx, r, cdc)
	registerTxRoutes(cliCtx, r, cdc, kb)
}
