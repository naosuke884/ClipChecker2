package provider

import (
	"github.com/naosuke884/ClipChecker2/later/domain"
	"github.com/naosuke884/ClipChecker2/later/share"
)

type TwitchClient struct {
}

func NewTwitchClient() share.Auth {
	return &TwitchClient{}
}

func (t *TwitchClient) Auth(user domain.User) (domain.User, error) {
	return domain.User{}, nil
}
