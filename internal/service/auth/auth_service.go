package authsvc

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/javascriptizer1/grpc-cli-chat.backend/internal/domain/user"
	serviceDto "github.com/javascriptizer1/grpc-cli-chat.backend/internal/service/auth/dto"
	"github.com/javascriptizer1/grpc-cli-chat.backend/pkg/helper/jwt"
)

func (s *AuthService) Register(ctx context.Context, input serviceDto.RegisterInputDto) (id uuid.UUID, err error) {
	v, _ := s.userRepo.OneByEmail(ctx, input.Email)

	if v != nil {
		return id, errors.New("user with this email already exists")
	}

	if input.Password != input.PasswordConfirm {
		return id, errors.New("passwords don`t match")
	}

	u, err := user.New(input.Name, input.Email, input.Password, input.Role)

	if err != nil {
		return id, err
	}

	u.HashPassword()

	if err = s.userRepo.Create(ctx, u); err != nil {
		return id, err
	}

	return u.Id, nil
}

func (s *AuthService) Login(ctx context.Context, login string, password string) (string, error) {
	v, err := s.userRepo.OneByEmail(ctx, login)

	if v == nil {
		return "", errors.New("invalid login or password")
	}

	if err != nil {
		return "", err
	}

	u, err := user.NewWithID(v.Id, v.Name, v.Email, v.Password, v.Role, v.CreatedAt, v.UpdatedAt)

	if err != nil {
		return "", err
	}

	if !u.CheckPassword(password) {
		return "", errors.New("invalid login or password")
	}

	t, err := jwt.GenerateToken(*u, s.config.RefreshTokenSecretKey, s.config.RefreshTokenDuration)

	return t, err
}

func (s *AuthService) AccessToken(ctx context.Context, refreshToken string) (string, error) {
	claims, err := jwt.VerifyToken(refreshToken, s.config.RefreshTokenSecretKey)

	if err != nil {
		return "", errors.New("invalid refresh token")
	}

	stringId, err := claims.GetSubject()

	if err != nil {
		return "", err
	}

	id, err := uuid.Parse(stringId)

	if err != nil {
		return "", err
	}

	u, err := s.userRepo.OneById(ctx, uuid.UUID(id))

	if err != nil {
		return "", err
	}

	accessToken, err := jwt.GenerateToken(*u, s.config.AccessTokenSecretKey, s.config.AccessTokenDuration)

	if err != nil {
		return "", err
	}

	return accessToken, nil
}

func (s *AuthService) RefreshToken(ctx context.Context, oldRefreshToken string) (string, error) {
	claims, err := jwt.VerifyToken(oldRefreshToken, s.config.RefreshTokenSecretKey)

	if err != nil {
		return "", errors.New("invalid refresh token")
	}

	stringId, err := claims.GetSubject()

	if err != nil {
		return "", err
	}

	id, err := uuid.Parse(stringId)

	if err != nil {
		return "", err
	}

	u, err := s.userRepo.OneById(ctx, uuid.UUID(id))

	if err != nil {
		return "", err
	}

	refreshToken, err := jwt.GenerateToken(*u, s.config.RefreshTokenSecretKey, s.config.RefreshTokenDuration)

	if err != nil {
		return "", err
	}

	return refreshToken, nil
}

func (s *AuthService) Check(ctx context.Context, endpoint string, role user.Role) bool {
	return true
}
