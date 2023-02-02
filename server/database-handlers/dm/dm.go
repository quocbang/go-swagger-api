package dm

import (
	"context"

	"gorm.io/gorm"
)

// import "context"

type DataManager interface {
	DB() *gorm.DB
	// GetTokenInfo returns the information of the token.
	// The following input arguments are required:
	//  - Token
	//
	// The returned USER_ERROR would be as below:
	//  - Code_USER_UNKNOWN_TOKEN
	// GetTokenInfo(context.Context, GetTokenInfoRequest) (GetTokenInfoReply, error)

	// SignIn needs the following required input:
	//  - Name value(employee id)
	//  - Password
	// If ADUser is true, the Name value is AD account, otherwise, is MES account.
	//
	// The returned USER_ERROR would be as below:
	//  - Code_ACCOUNT_NOT_FOUND_OR_BAD_PASSWORD
	//  - Code_INSUFFICIENT_REQUEST
	//  - Code_DEPARTMENT_NOT_FOUND
	//
	// The token is expired in 8 hours in default.
	// SignIn also returns whether the user needs to change the password or not.

	// SignIn(context.Context, SignInRequest, ...SignInOption) (SignInReply, error)

	// // SignOut signs out.
	// //
	// // The returned USER_ERROR would be as below:
	// //  - Code_USER_UNKNOWN_TOKEN
	// SignOut(context.Context, SignOutRequest) error

	GetLimitaryHour(context.Context, GetLimitaryRequest) (GetLimitaryReply, error)
}
