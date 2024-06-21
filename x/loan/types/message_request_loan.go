package types

import (
	"strconv"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgRequestLoan{}

func NewMsgRequestLoan(creator string, amount string, fee string, collateral string, deadline string) *MsgRequestLoan {
	return &MsgRequestLoan{
		Creator:    creator,
		Amount:     amount,
		Fee:        fee,
		Collateral: collateral,
		Deadline:   deadline,
	}
}

func (msg *MsgRequestLoan) ValidateBasic() error {
	// 验证入参
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		// 验证地址
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	amount, _ := sdk.ParseCoinsNormalized(msg.Amount)
	if !amount.IsValid() {
		// 验证数量
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "invalid amount (%s)", msg.Amount)
	}
	if amount.Empty() {
		// 验证数量
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "amount cannot be empty")
	}

	fee, _ := sdk.ParseCoinsNormalized(msg.Fee)
	if !fee.IsValid() {
		// 验证fee
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "invalid fee (%s)", msg.Fee)
	}

	deadline, err := strconv.ParseInt(msg.Deadline, 10, 64)
	if err != nil {
		// 验证deadline
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "invalid deadline (%s)", msg.Deadline)
	}

	if deadline <= 0 {
		// 验证deadline
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "deadline cannot be negative")
	}

	collateral, _ := sdk.ParseCoinsNormalized(msg.Collateral)
	if !collateral.IsValid() {
		// 验证collateral
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "invalid collateral (%s)", msg.Collateral)
	}
	if collateral.Empty() {
		// 验证collateral
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "collateral cannot be empty")
	}
	return nil
}
