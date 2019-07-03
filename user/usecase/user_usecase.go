package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/febrycode/healthy_food/models"
	"github.com/febrycode/healthy_food/user"
	"github.com/febrycode/healthy_food/util"
)

type Usecase struct {
	userRepository user.Repository
	contextTimeout time.Duration
}

func NewUserUsecase(u user.Repository, timeout time.Duration) user.Usecase {
	return &Usecase{
		userRepository: u,
		contextTimeout: timeout,
	}
}

func (uc *Usecase) GetUserByEmail(ctx context.Context, email string) (models.User, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	res, err := uc.userRepository.GetByEmail(ctx, email)
	if err != nil {
		return models.User{}, err
	}

	return res, nil
}

func (uc *Usecase) CreateUser(ctx context.Context, userData *models.User) error {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	userData.CreatedAt = util.GetTimeNow()

	fmt.Println(userData)

	err := uc.userRepository.CreateUser(ctx, userData)
	if err != nil {
		return err
	}

	return nil
}
