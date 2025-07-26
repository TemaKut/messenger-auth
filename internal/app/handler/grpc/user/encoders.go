package user

import (
	"errors"
	userdto "github.com/TemaKut/messenger-auth/internal/dto/user"
	authv1 "github.com/TemaKut/messenger-service-proto/gen/go/auth"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func encodeUser(user userdto.User) *authv1.User {
	return &authv1.User{
		Id:       user.Id,
		Name:     user.Name,
		LastName: user.LastName,
	}
}

func encodeAuthorizeResult(result userdto.UserAuthorizeResult) *authv1.UserAPIAuthorizeResponse {
	return &authv1.UserAPIAuthorizeResponse{
		User:       encodeUser(result.User),
		AuthParams: encodeAuthParams(result.AuthParams),
	}
}

func encodeAuthParams(params userdto.AuthParams) *authv1.AuthParams {
	return &authv1.AuthParams{
		AccessToken:  encodeAuthToken(params.AccessToken),
		RefreshToken: encodeAuthToken(params.RefreshToken),
	}
}

func encodeAuthToken(token userdto.AuthToken) *authv1.AuthToken {
	return &authv1.AuthToken{
		Token:     token.Token,
		ExpiredAt: timestamppb.New(token.ExpiredAt),
	}
}

func encodeError(err error) error {
	if err == nil {
		return nil
	}

	switch {
	case errors.Is(err, userdto.ErrUserEmailAlreadyExists):
		st, innerErr := status.Convert(err).
			WithDetails(&errdetails.ErrorInfo{Reason: ErrorReasonUserEmailAlreadyExists})
		if innerErr != nil {
			return status.Errorf(codes.Unknown, "error-unknown. %s. %s", innerErr, err)
		}

		return st.Err()
	case errors.Is(err, userdto.ErrInvalidCredentials):
		st, innerErr := status.Convert(err).
			WithDetails(&errdetails.ErrorInfo{Reason: ErrorInvalidCredentials})
		if innerErr != nil {
			return status.Errorf(codes.Unknown, "error-unknown. %s. %s", innerErr, err)
		}

		return st.Err()
	default:
		return status.Errorf(codes.Unknown, "error-unknown. %s", err)
	}
}
