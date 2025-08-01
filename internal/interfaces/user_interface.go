package interfaces

import (
	"stargazer/internal/types"
)

type UserInterface interface {
	BaseInterface
	GetUsername() string
	GetEmail() string
	GetAvatar() string
	GetNickname() string
	GetDescription() string
	GetPassword() string
	GetStatus() types.UserStatus
	SetUsername(username string)
	SetEmail(email string)
	SetAvatar(avatar string)
	SetNickname(nickname string)
	SetDescription(description string)
	SetPassword(password string)
	SetStatus(status types.UserStatus)
	Validate() error
}
