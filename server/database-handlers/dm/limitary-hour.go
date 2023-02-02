package dm

import "time"

type GetLimitaryRequest struct {
	ProductType string
}

type GetLimitaryReply struct {
	Min time.Time
	Max time.Time
}
