package mint

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/she-protocol/she-chain/x/mint/keeper"
	"github.com/she-protocol/she-chain/x/mint/types"
)

func HandleUpdateMinterProposal(ctx sdk.Context, k *keeper.Keeper, p *types.UpdateMinterProposal) error {
	err := types.ValidateMinter(*p.Minter)
	if err != nil {
		return err
	}
	k.SetMinter(ctx, *p.Minter)
	return nil
}
