package share

import (
	"ClipChecker2/domain"
)

type Auth interface {
	Auth(user domain.User) (domain.User, error)
}
