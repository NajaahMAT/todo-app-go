package usecase

import (
	"context"
	"time"
	"todo-app-go/domain"
	"todo-app-go/internal/util"
)

type loginUsecase struct {
	userRepository domain.UserRepository
	contextTimeout time.Duration
}

func NewLoginUsecase(userRepository domain.UserRepository, timeout time.Duration) domain.LoginUsecase {
	return &loginUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func (l *loginUsecase) GetUserByEmail(c context.Context, email string) (domain.User, error) {
	ctx, cancel := context.WithTimeout(c, l.contextTimeout)
	defer cancel()
	return l.userRepository.GetByEmail(ctx, email)
}

func (l *loginUsecase) CreateAccessToken(user *domain.User, secret string, expiry int) (accessToken string, err error) {
	return util.CreateAccessToken(user, secret, expiry)
}

func (l *loginUsecase) CreateRefreshToken(user *domain.User, secret string, expiry int) (refreshToken string, err error) {
	return util.CreateRefreshToken(user, secret, expiry)
}
