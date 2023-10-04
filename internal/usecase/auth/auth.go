package auth

import (
	"context"

	"github.com/TitusW/game-service/internal/entity"
	"github.com/TitusW/game-service/pkg/token"
	"golang.org/x/crypto/bcrypt"
)

func (uc Usecase) Login(ctx context.Context, input entity.User) (entity.TokenResponse, error) {
	user, err := uc.userResource.Get(ctx, input.Email)
	if err != nil {
		return entity.TokenResponse{}, err
	}

	if user.Email == "" {
		return entity.TokenResponse{}, err
	}

	passwordIsValid, err := verifyPassword(input.Password, user.Password)
	if passwordIsValid != true {
		return entity.TokenResponse{}, err
	}

	token, refreshToken, err := token.GenerateAllTokens(user.Email, user.Ksuid)
	if err != nil {
		return entity.TokenResponse{}, err
	}

	tokenResponse := entity.TokenResponse{
		Token:        token,
		RefreshToken: refreshToken,
	}

	tokens, err := uc.tokenResource.ScanUserTokens(ctx, user.Ksuid, "")
	if err != nil {
		return entity.TokenResponse{}, err
	}

	for _, token := range tokens {
		uc.tokenResource.DeleteUserToken(ctx, user.Ksuid, token)
	}

	uc.tokenResource.SetUserToken(ctx, token, refreshToken)

	return tokenResponse, nil
}

func (uc Usecase) Logout(ctx context.Context, input entity.User) error {
	tokens, err := uc.tokenResource.ScanUserTokens(ctx, input.Email, "")
	if err != nil {
		return err
	}

	for _, token := range tokens {
		uc.tokenResource.DeleteUserToken(ctx, input.Email, token)

		if err != nil {
			return err
		}
	}

	return nil
}

func verifyPassword(userPassword string, providedPassword string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(providedPassword))
	if err != nil {
		return false, err
	}

	return true, nil
}
