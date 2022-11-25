package share

import "github.com/naosuke884/ClipChecker2/later/domain"

type Auth interface {
	Auth(user domain.User) (domain.User, error)
}
