package provider

import (
	"ClipChecker2/domain"
	"ClipChecker2/share"
)

type TwitchClient struct {
}

func NewTwitchClient() share.Auth {
	return &TwitchClient{}
}

func (t *TwitchClient) Auth(user domain.User) (domain.User, error) {
	return domain.User{}, nil
}
