package auth

import (
	"context"
	"fmt"

	"github.com/TitusW/game-service/internal/entity"
	tokenhelper "github.com/TitusW/game-service/pkg/token"
	"golang.org/x/crypto/bcrypt"
)

func (uc Usecase) Login(ctx context.Context, input entity.User) (entity.TokenResponse, error) {
	user, err := uc.userResource.GetByEmail(ctx, input.Email)
	if err != nil {
		return entity.TokenResponse{}, err
	}

	if user.Email == "" {
		return entity.TokenResponse{}, err
	}

	passwordIsValid, err := verifyPassword(user.Password, input.Password)
	if passwordIsValid != true {
		return entity.TokenResponse{}, err
	}

	token, refreshToken, err := tokenhelper.GenerateAllTokens(user.Email, user.Ksuid)
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
		uc.tokenResource.DeleteUserTokenByKey(ctx, token)
	}

	uc.tokenResource.SetUserToken(ctx, user.Ksuid, token)

	return tokenResponse, nil
}

func (uc Usecase) Logout(ctx context.Context, tokenString string) error {
	claims, err := tokenhelper.ExtractUnverifiedClaims(tokenString)
	if err != nil {
		return err
	}

	key := fmt.Sprintf("%s:%s", claims["Ksuid"], tokenString)
	uc.tokenResource.DeleteUserTokenByKey(ctx, key)

	return nil
}

func verifyPassword(userPassword string, providedPassword string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(providedPassword))
	if err != nil {
		return false, err
	}

	return true, nil
}
