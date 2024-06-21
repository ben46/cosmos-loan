package keeper

import (
	"context"
	"loan/x/loan/types"

	errorsmod "cosmossdk.io/errors"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CancelLoan(goCtx context.Context, msg *types.MsgCancelLoan) (*types.MsgCancelLoanResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	loan, found := k.GetLoan(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "loan %d not found", msg.Id)
	}

	if loan.Borrower != msg.Creator {
		return nil, errorsmod.Wrapf(sdkerrors.ErrUnauthorized, "loan %d can only be canceled by borrower", msg.Id)
	}

	if loan.State != "requested" {
		return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "loan %d cannot be canceled", msg.Id)
	}

	borrower, _ := sdk.AccAddressFromBech32(msg.Creator)
	collateral, _ := sdk.ParseCoinsNormalized(loan.Collateral)
	err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, borrower, collateral)
	if err != nil {
		return nil, err
	}
	loan.State = "canceled"
	k.SetLoan(ctx, loan)
	return &types.MsgCancelLoanResponse{}, nil
}
