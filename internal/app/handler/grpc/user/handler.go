package user

import (
	"context"
	"fmt"
	userdto "github.com/TemaKut/messenger-auth/internal/dto/user"
	authv1 "github.com/TemaKut/messenger-service-proto/gen/go/auth"
)

type Handler struct {
	authv1.UnimplementedUserAPIServer

	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Register(
	ctx context.Context,
	req *authv1.UserAPIRegisterRequest,
) (*authv1.UserAPIRegisterResponse, error) {
	user, err := h.service.Register(ctx, userdto.RegisterParams{
		Name:     req.GetName(),
		LastName: req.GetLastName(),
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
	})
	if err != nil {
		return nil, fmt.Errorf("error register user. %w", encodeError(err))
	}

	return &authv1.UserAPIRegisterResponse{
		User: encodeUser(user),
	}, nil
}

func (h *Handler) Authorize(
	ctx context.Context,
	req *authv1.UserAPIAuthorizeRequest,
) (*authv1.UserAPIAuthorizeResponse, error) {
	authorizeParams, err := decodeAuthorizeRequest(req)
	if err != nil {
		return nil, fmt.Errorf("error decode authorize request. %w", err)
	}

	authorizeResult, err := h.service.Authorize(ctx, authorizeParams)
	if err != nil {
		return nil, fmt.Errorf("error authorize user. %w", encodeError(err))
	}

	return encodeAuthorizeResult(authorizeResult), nil
}

func toPtr[T any](v T) *T {
	return &v
}
