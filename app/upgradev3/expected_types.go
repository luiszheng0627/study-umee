package upgradev3

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/staking/types"
)

type StakingKeeper interface {
	GetParams(ctx sdk.Context) types.Params
	SetParams(ctx sdk.Context, params types.Params)
	GetAllValidators(ctx sdk.Context) (validators []types.Validator)
	BeforeValidatorModified(ctx sdk.Context, valAddr sdk.ValAddress) error
	SetValidator(ctx sdk.Context, validator types.Validator)
	GetValidator(ctx sdk.Context, addr sdk.ValAddress) (types.Validator, bool)
}
