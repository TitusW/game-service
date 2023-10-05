package user

import (
	"context"

	"github.com/TitusW/game-service/internal/entity"
	"golang.org/x/crypto/bcrypt"
)

func (uc Usecase) Register(ctx context.Context, input entity.User) (entity.User, error) {
	password, err := hashPassword(input.Password)
	if err != nil {
		return entity.User{}, err
	}

	input.Password = password
	user, err := uc.resource.Create(ctx, input)
	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (uc Usecase) Update(ctx context.Context, input entity.User) (entity.User, error) {
	user, err := uc.resource.Update(ctx, input)
	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
