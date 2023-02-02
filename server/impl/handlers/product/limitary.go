package product

import (
	"quocbang/go-swagger-api/server/database-handlers/dm"
	"quocbang/go-swagger-api/server/impl/service"
)

type NewLimitary struct {
	dm dm.DataManager
}

// NewAuthorization returns Authorization service.
func NewLimitaryHour(dm dm.DataManager) service.Production {
	return NewLimitary{
		dm: dm,
	}
}

func (d dm.DataManager) Getlimitary() {
	a := d.GetLimitaryHour()
}
