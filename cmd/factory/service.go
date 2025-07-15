package factory

import (
	usergrpchandler "github.com/TemaKut/messenger-auth/internal/app/handler/grpc/user"
	userservice "github.com/TemaKut/messenger-auth/internal/service/user"
	"github.com/google/wire"
)

var ServiceSet = wire.NewSet(
	userservice.NewService,
	wire.Bind(new(usergrpchandler.Service), new(*userservice.Service)),
)
