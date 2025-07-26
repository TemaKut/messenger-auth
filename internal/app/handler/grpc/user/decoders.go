package user

import (
	"fmt"
	userdto "github.com/TemaKut/messenger-auth/internal/dto/user"
	authv1 "github.com/TemaKut/messenger-service-proto/gen/go/auth"
)

func decodeAuthorizeRequest(req *authv1.UserAPIAuthorizeRequest) (userdto.UserAuthorizeParams, error) {
	var params userdto.UserAuthorizeParams

	switch {
	case req.GetEmail() != nil:
		params.Credentials.Email = toPtr(decodeUserAuthorizeEmailCredentials(req.GetEmail()))
	default:
		return userdto.UserAuthorizeParams{}, fmt.Errorf("error credentials not passed")
	}

	return params, nil
}

func decodeUserAuthorizeEmailCredentials(email *authv1.UserAPIAuthorizeEmailCredentials) userdto.UserAuthorizeEmailCredential {
	return userdto.UserAuthorizeEmailCredential{
		Email:    email.Email,
		Password: email.Password,
	}
}
