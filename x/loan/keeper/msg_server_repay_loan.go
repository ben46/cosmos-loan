package keeper

import (
	"context"

	"loan/x/loan/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) RepayLoan(goCtx context.Context, msg *types.MsgRepayLoan) (*types.MsgRepayLoanResponse, error) {
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

	lender, _ := sdk.AccAddressFromBech32(loan.Lender)
	borrower, _ := sdk.AccAddressFromBech32(msg.Creator)
	amount, err := sdk.ParseCoinsNormalized(loan.Amount)
	if err != nil {
		return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "parse amount error")
	}
	if amount.Empty() {
		return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "amount is empty")
	}
	if amount.IsZero() {
		return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "amount is zero")
	}
	if amount == nil {
		return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "amount is nil")
	}

	err = k.bankKeeper.SendCoins(ctx, borrower, lender, amount)
	if err != nil {
		return nil, err
	}

	collateral, _ := sdk.ParseCoinsNormalized(loan.Collateral)
	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, borrower, collateral)
	if err != nil {
		return nil, err
	}
	loan.State = "repayed"
	k.SetLoan(ctx, loan)

	return &types.MsgRepayLoanResponse{}, nil
}
