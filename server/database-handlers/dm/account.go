package dm

import (
	"time"
)

// import (
// 	"time"

// 	"gitlab.kenda.com.tw/kenda/mcom/impl/orm/models"
// 	"gitlab.kenda.com.tw/kenda/mcom/utils/roles"
// )

// GetTokenInfoRequest definition.
// type GetTokenInfoRequest struct {
// 	Token string
// }

// GetTokenInfoReply definition.
// type GetTokenInfoReply struct {
// 	User        string
// 	Valid       bool
// 	ExpiryTime  time.Time
// 	CreatedTime time.Time
// 	Roles       []roles.Role
// }

type SignInRequest struct {
	Account  string
	Password string

	ADUser bool
}

// Department definition.
type Department struct {
	OID string
	ID  string
}

// SignInReply definition.
type SignInReply struct {
	Token       string
	TokenExpiry time.Time
	Departments []Department
	// Roles              []roles.Role
	MustChangePassword bool
}

// SignInOptions definition.
type SignInOptions struct {
	ExpiredAfter time.Duration
}

// SignInOption definition.
type SignInOption func(*SignInOptions)

func WithTokenExpiredAfter(d time.Duration) SignInOption {
	return func(o *SignInOptions) {
		o.ExpiredAfter = d
	}
}

func newSignInOptions() SignInOptions {
	return SignInOptions{
		ExpiredAfter: 8 * time.Hour,
	}
}

// ParseSignInOptions returns options for sign-in.
func ParseSignInOptions(opts []SignInOption) SignInOptions {
	o := newSignInOptions()
	for _, opt := range opts {
		opt(&o)
	}
	return o
}
