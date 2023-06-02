package usecase

import (
	"context"
	"time"
	"todo-app-go/domain"
	"todo-app-go/internal/util"
)

type refreshTokenUsecase struct {
	userRepository domain.UserRepository
	contextTimeout time.Duration
}

func NewRefreshTokenUsecase(userRepository domain.UserRepository, timeout time.Duration) domain.RefreshTokenUsecase {
	return &refreshTokenUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func (rt *refreshTokenUsecase) GetUserByID(c context.Context, id string) (domain.User, error) {
	ctx, cancel := context.WithTimeout(c, rt.contextTimeout)
	defer cancel()
	return rt.userRepository.GetByID(ctx, id)
}

func (rt *refreshTokenUsecase) CreateAccessToken(user *domain.User, secret string, expiry int) (accessToken string, err error) {
	return util.CreateAccessToken(user, secret, expiry)
}

func (rt *refreshTokenUsecase) CreateRefreshToken(user *domain.User, secret string, expiry int) (refreshToken string, err error) {
	return util.CreateRefreshToken(user, secret, expiry)
}

func (rt *refreshTokenUsecase) ExtractIDFromToken(requestToken string, secret string) (string, error) {
	return util.ExtractIDFromToken(requestToken, secret)
}
