package types

import minttypes "github.com/osmosis-labs/osmosis/v13/x/mint/types"

// MintParams represents the x/mint parameters
type MintParams struct {
	minttypes.Params
	Height int64
}

// NewMintParams allows to build a new MintParams instance
func NewMintParams(params minttypes.Params, height int64) *MintParams {
	return &MintParams{
		Params: params,
		Height: height,
	}
}
