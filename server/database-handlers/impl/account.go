package impl

import (
	"context"
	"quocbang/go-swagger-api/server/database-handlers/dm"
	api "quocbang/go-swagger-api/server/database-handlers/dm"
)

func (dmdb *DataManager) SignIn(ctx context.Context, req dm.SignInRequest, opts ...dm.SignInOption) (dm.SignInReply, error) {
	// opt := dm.ParseSignInOptions(opts)

	// session := dmdb.newSession(ctx)

	// expiryTime := time.Now().Add(opt.ExpiredAfter).In(time.Local)
	// a, err := session.createToken(req.Account, expiryTime)
	return api.SignInReply{}, nil
}
