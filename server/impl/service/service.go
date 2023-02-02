package service

import (
	"quocbang/go-swagger-api/server/database-handlers/dm"
	"quocbang/go-swagger-api/server/swagger/restapi/operations/account"
	"quocbang/go-swagger-api/server/swagger/restapi/operations/production"

	"github.com/go-openapi/runtime/middleware"
)

// Service Definition.
type Service struct {
	accountAuthorization AccountAuthorization
	production           Production
	// add more service
}

// NewService return new RESTful API services.
func NewService(
	accountAuthorization AccountAuthorization,
	production Production,

) *Service {
	return &Service{
		accountAuthorization: accountAuthorization,
		production:           production,
	}
}

// AccountAuthorization return account authorization services.
func (s *Service) AccountAuthorization() AccountAuthorization {
	return s.accountAuthorization
}

type AccountAuthorization interface {
	Login(params account.LoginParams) middleware.Responder
}

func (s *Service) Production() Production {
	return s.production
}

type Production interface {
	Getlimitary(params production.GetlimitaryParams) (dm.GetLimitaryReply, error)
}
