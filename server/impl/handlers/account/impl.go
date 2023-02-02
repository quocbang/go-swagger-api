package account

import (
	"time"

	"quocbang/go-swagger-api/server/database-handlers/dm"
	"quocbang/go-swagger-api/server/impl/service"
	"quocbang/go-swagger-api/server/swagger/restapi/operations/account"

	"github.com/go-openapi/runtime/middleware"
)

type Authorization struct {
	dm            dm.DataManager
	tokenLifeTime time.Duration
}

// NewAuthorization returns Authorization service.
func NewAuthorization(dm dm.DataManager, tokenLifeTime time.Duration) service.AccountAuthorization {
	return Authorization{
		dm:            dm,
		tokenLifeTime: tokenLifeTime,
	}
}

func (a Authorization) Login(p account.LoginParams) middleware.Responder {
	// 	var options []dm.SignInOption
	// 	if a.tokenLifeTime > 0 {
	// 		options = append(options, dm.WithTokenExpiredAfter(a.tokenLifeTime))
	// 	}

	// 	signInResult, err := a.dm.SignIn(p.HTTPRequest.Context(), dm.SignInRequest{
	// 		Account:  "quocbang",
	// 		Password: "quocbang",
	// 	}, options...)

	// 	if err != nil {
	// 		return
	// 	}
	return nil
}
