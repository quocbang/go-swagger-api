package handlers

import (
	"time"

	"quocbang/go-swagger-api/server/database-handlers/dm"
	accountImpl "quocbang/go-swagger-api/server/impl/handlers/account"
	productImpl "quocbang/go-swagger-api/server/impl/handlers/product"
	"quocbang/go-swagger-api/server/impl/service"
	"quocbang/go-swagger-api/server/swagger/restapi/operations"
	"quocbang/go-swagger-api/server/swagger/restapi/operations/account"
	"quocbang/go-swagger-api/server/swagger/restapi/operations/production"
)

type ServiceConfig struct {
	TokenLifeTime time.Duration
}

func RegisterServices(dm dm.DataManager, config ServiceConfig) (*service.Service, error) {

	return service.NewService(
		accountImpl.NewAuthorization(dm, config.TokenLifeTime),
		productImpl.NewLimitaryHour(dm),
	), nil

}

func RegisterHandlers(dm dm.DataManager, api *operations.GoSwaggerAPIAPI, config ServiceConfig) error {

	s, err := RegisterServices(dm, config)
	if err != nil {
		return err
	}

	api.AccountLoginHandler = account.LoginHandlerFunc(s.AccountAuthorization().Login)

	api.ProductionGetlimitaryHandler = production.GetlimitaryHandlerFunc(s.Production().Getlimitary)
	return nil
}
