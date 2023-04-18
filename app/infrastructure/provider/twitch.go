package provider

import (
	"log"
	"os"

	"github.com/go-resty/resty/v2"
	"github.com/naosuke884/ClipChecker2/later/domain"
)

//TwitchAPIへのリクエスト用の関数
//1. getTwitchUserToken
//2. getTwitchUser
//3. getTwitchUserFollows
//4. getTwitchClips

type TwitchProvider struct {
	token  string
	client *resty.Client
}

type ITwitchProvider interface {
	GetUser() (*domain.User, error)
	GetFollows() ([]domain.BroadCaster, error)
	GetClips() ([]domain.Clip, error)
	GetToken() (string, error)
}

func NewTwitchProvider(client *resty.Client) ITwitchProvider {
	token, err := getTwitchUserToken(client)
	if err != nil {
		log.Fatal(err)
	}
	return &TwitchProvider{
		token:  token,
		client: client,
	}
}

func (t *TwitchProvider) GetToken() (string, error) {
	return t.token, nil
}

func (t *TwitchProvider) GetUser() (*domain.User, error) {
	return nil, nil
}

func (t *TwitchProvider) GetFollows() ([]domain.BroadCaster, error) {
	return nil, nil
}

func (t *TwitchProvider) GetClips() ([]domain.Clip, error) {
	return nil, nil
}

func getTwitchUserToken(client *resty.Client) (string, error) {
	resp, err := client.R().
		SetFormData(map[string]string{
			"client_id":     os.Getenv("CLIENT_ID"),
			"client_secret": os.Getenv("CLIENT_SECRET"),
			"grant_type":    "client_credentials",
		}).
		Post("https://id.twitch.tv/oauth2/token")
	return string(resp.Body()), err
}
