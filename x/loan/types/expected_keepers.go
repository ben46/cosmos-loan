package types

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// AccountKeeper defines the expected interface for the Account module.
type AccountKeeper interface {
	GetAccount(context.Context, sdk.AccAddress) sdk.AccountI // only used for simulation
	// Methods imported from account should be defined here
}

// BankKeeper defines the expected interface for the Bank module.
type BankKeeper interface {
	// SpendableCoins(context.Context, sdk.AccAddress) sdk.Coins
	// 发送代币
	SendCoins(ctx context.Context, fromAddress sdk.AccAddress, toAddress sdk.AccAddress, amt sdk.Coins) error
	// 发送代币
	SendCoinsFromAccountToModule(ctx context.Context, senderAddress sdk.AccAddress, receipientModule string, amt sdk.Coins) error
	// 发送代币
	SendCoinsFromModuleToAccount(ctx context.Context, senderModule string, receiptAddr sdk.AccAddress, amt sdk.Coins) error
}

// ParamSubspace defines the expected Subspace interface for parameters.
type ParamSubspace interface {
	Get(context.Context, []byte, interface{})
	Set(context.Context, []byte, interface{})
}
