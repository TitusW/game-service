package auth

import (
	"context"

	"github.com/TitusW/game-service/internal/entity"
)

type UserModuleItf interface {
	Get(context.Context, string) (entity.User, error)
}

type TokenModuleItf interface {
	SetUserToken(context.Context, string, string) error
	ScanUserTokens(context.Context, string, string) ([]string, error)
	DeleteUserToken(context.Context, string, string) error
}

type Usecase struct {
	userResource  UserModuleItf
	tokenResource TokenModuleItf
}

func New(userResource UserModuleItf, tokenResource TokenModuleItf) Usecase {
	return Usecase{
		userResource:  userResource,
		tokenResource: tokenResource,
	}
}
