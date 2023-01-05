package usecase

import (
	"context"
	"fmt"

	"github.com/Doer-org/google-cloud-challenge-2022/domain/entity"
	"github.com/Doer-org/google-cloud-challenge-2022/domain/repository"
)

type IUserUsecase interface {
	Create(ctx context.Context, name string, authenticated bool, mail string, icon string) (*entity.User, error)
	GetByMail(ctx context.Context, mail string) (*entity.User, error)
	GetById(ctx context.Context, id string) (*entity.User, error)
}

type UserUsecase struct {
	repo repository.IUserRepository
}

func NewUserUsecase(r repository.IUserRepository) IUserUsecase {
	return &UserUsecase{
		repo: r,
	}
}

func (uc *UserUsecase) Create(ctx context.Context, name string, authenticated bool, mail string, icon string) (*entity.User, error) {
	if name == "" {
		return nil, fmt.Errorf("UserUsecase: name is empty")
	}
	if icon == "" {
		return nil, fmt.Errorf("UserUsecase: icon is empty")
	}
	// TODO: mailが存在するかの確認

	user := &entity.User{
		Name:          name,
		Authenticated: authenticated,
		Mail:          mail,
		Icon:          icon,
	}
	return uc.repo.Create(ctx, user)
}

func (uc *UserUsecase) GetByMail(ctx context.Context, mail string) (*entity.User, error) {
	if mail == "" {
		return nil, fmt.Errorf("UserUsecase: mail is empty")
	}
	return uc.repo.GetByMail(ctx, mail)
}

func (uc *UserUsecase) GetById(ctx context.Context, userIdString string) (*entity.User, error) {
	userId := entity.UserId(userIdString)
	if userId == "" {
		return nil, fmt.Errorf("UserUsecase: userId parse failed")
	}
	return uc.repo.GetById(ctx, userId)
}
