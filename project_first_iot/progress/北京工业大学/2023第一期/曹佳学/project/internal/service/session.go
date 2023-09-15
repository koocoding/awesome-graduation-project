// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	"github.com/gogf/gf-demo-user/v2/internal/model/entity"
)

type (
	ISession interface {
		// SetUser sets user into the session.
		SetUser(ctx context.Context, user *entity.User) (string, error)
		// GetUser retrieves and returns the user from session.
		// It returns nil if the user did not sign in.
		GetUser(ctx context.Context) *entity.User
		// RemoveUser removes user rom session.
		RemoveUser(ctx context.Context) error
	}
)

var (
	localSession ISession
)

func Session() ISession {
	if localSession == nil {
		panic("implement not found for interface ISession, forgot register?")
	}
	return localSession
}

func RegisterSession(i ISession) {
	localSession = i
}
