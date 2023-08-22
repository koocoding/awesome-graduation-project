// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	"github.com/gogf/gf-demo-user/v2/internal/model"
	"github.com/gogf/gf-demo-user/v2/internal/model/entity"
)

type (
	IUser interface {
		// Create creates user account.
		Create(ctx context.Context, in model.UserCreateInput) (err error)
		// IsSignedIn checks and returns whether current user is already signed-in.
		IsSignedIn(ctx context.Context) bool
		// SignIn creates session for given user account.
		SignIn(ctx context.Context, in model.UserSignInInput) (err error)
		// SignOut removes the session for current signed-in user.
		SignOut(ctx context.Context) error
		// IsPassportAvailable checks and returns given passport is available for signing up.
		IsPassportAvailable(ctx context.Context, passport string) (bool, error)
		// IsNicknameAvailable checks and returns given nickname is available for signing up.
		IsNicknameAvailable(ctx context.Context, nickname string) (bool, error)
		// GetProfile retrieves and returns current user info in session.
		GetProfile(ctx context.Context) *entity.User
		VerifyCodeSend(ctx context.Context, PhoneNumber string) (err error)
		VerifyCodeCheck(ctx context.Context, checkverifycode string, phonenumber string) (bool, error)
	}
)

var (
	localUser IUser
)

func User() IUser {
	if localUser == nil {
		panic("implement not found for interface IUser, forgot register?")
	}
	return localUser
}

func RegisterUser(i IUser) {
	localUser = i
}
