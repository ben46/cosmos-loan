package keeper

import (
	"context"
	"strconv"

	"loan/x/loan/types"

	errorsmod "cosmossdk.io/errors"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) LiquidateLoan(goCtx context.Context, msg *types.MsgLiquidateLoan) (*types.MsgLiquidateLoanResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// get loan
	loan, found := k.GetLoan(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrapf(sdkerrors.ErrKeyNotFound, "loan with id %d not found", msg.Id)
	}

	// check loan state
	if loan.State != "approved" {
		return nil, errorsmod.Wrapf(types.ErrWrongLoanState, "loan state is %s", loan.State)
	}

	// only lender can liquidate loan
	if loan.Lender != msg.Creator {
		return nil, errorsmod.Wrapf(sdkerrors.ErrUnauthorized, "only lender can liquidate loan")
	}

	lender, _ := sdk.AccAddressFromBech32(loan.Lender)
	collateral, _ := sdk.ParseCoinsNormalized(loan.Collateral)
	deadline, err := strconv.ParseInt(loan.Deadline, 10, 64)
	if err != nil {
		return nil, errorsmod.Wrapf(types.ErrDeadline, "invalid deadline")
	}

	if ctx.BlockHeight() < deadline {
		return nil, errorsmod.Wrap(types.ErrDeadline, "can not liquidate loan before deadline")
	}

	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, lender, collateral)
	if err != nil {
		return nil, err
	}

	loan.State = "liquidated"
	k.SetLoan(ctx, loan)
	return &types.MsgLiquidateLoanResponse{}, nil
}
