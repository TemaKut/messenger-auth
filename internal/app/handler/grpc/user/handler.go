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

func (h *Handler) UserRegister(
	ctx context.Context,
	req *authv1.UserAPIUserRegisterRequest,
) (*authv1.UserAPIUserRegisterResponse, error) {
	user, err := h.service.Register(ctx, userdto.RegisterParams{
		Name:     req.GetName(),
		LastName: req.GetLastName(),
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
	})
	if err != nil {
		return nil, fmt.Errorf("error register user. %w", err)
	}

	return &authv1.UserAPIUserRegisterResponse{
		User: encodeUser(user),
	}, nil
}
