package models

import (
	"errors"
	"stargazer/internal/interfaces"
	"stargazer/internal/types"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID          uuid.UUID        `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Username    string           `json:"username" gorm:"not null;uniqueIndex;size:50;index"`
	Email       string           `json:"email" gorm:"not null;uniqueIndex;size:255;index"`
	Avatar      string           `json:"avatar" gorm:"size:255"`
	Password    string           `json:"-" gorm:"not null;size:255"`
	Nickname    string           `json:"nickname" gorm:"not null;size:50;index"`
	Description string           `json:"description" gorm:"size:500"`
	Status      types.UserStatus `json:"status" gorm:"default:'active';index"`
	CreatedAt   time.Time        `json:"created_at" gorm:"autoCreateTime;index"`
	UpdatedAt   time.Time        `json:"updated_at" gorm:"autoUpdateTime;index"`
	DeletedAt   gorm.DeletedAt   `json:"deleted_at" gorm:"index"`
}

func (u *User) GetID() uuid.UUID {
	return u.ID
}

func (u *User) SetID(id uuid.UUID) {
	u.ID = id
}

func (u *User) GetUsername() string {
	return u.Username
}

func (u *User) SetUsername(username string) {
	u.Username = username
}

func (u *User) GetEmail() string {
	return u.Email
}

func (u *User) SetEmail(email string) {
	u.Email = email
}

func (u *User) GetAvatar() string {
	return u.Avatar
}

func (u *User) SetAvatar(avatar string) {
	u.Avatar = avatar
}

func (u *User) GetPassword() string {
	return u.Password
}

func (u *User) SetPassword(password string) {
	u.Password = password
}

func (u *User) GetNickname() string {
	return u.Nickname
}

func (u *User) SetNickname(nickname string) {
	u.Nickname = nickname
}

func (u *User) GetDescription() string {
	return u.Description
}

func (u *User) SetDescription(description string) {
	u.Description = description
}

func (u *User) GetStatus() types.UserStatus {
	return u.Status
}

func (u *User) SetStatus(status types.UserStatus) {
	u.Status = status
}

func (u *User) GetCreatedAt() time.Time {
	return u.CreatedAt
}

func (u *User) SetCreatedAt(createdAt time.Time) {
	u.CreatedAt = createdAt
}

func (u *User) GetUpdatedAt() time.Time {
	return u.UpdatedAt
}

func (u *User) SetUpdatedAt(updatedAt time.Time) {
	u.UpdatedAt = updatedAt
}

func (u *User) Validate() error {
	if u.Username == "" {
		return errors.New("model: username is required")
	}
	if len(u.Username) < 3 || len(u.Username) > 50 {
		return errors.New(
			"model: username must be between 3 and 50 characters",
		)
	}
	if u.Email == "" {
		return errors.New(
			"model: email is required",
		)
	}
	return nil
}

func NewUser(username, email, password, nickname,
	description, avatar string) interfaces.UserInterface {
	return &User{
		ID:          uuid.New(),
		Username:    username,
		Email:       email,
		Password:    password,
		Nickname:    nickname,
		Description: description,
		Avatar:      avatar,
		Status:      types.UserStatusActive,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

func (u *User) AsUserInterface() interfaces.UserInterface {
	return u
}

func (u *User) AsBaseInterface() interfaces.BaseInterface {
	return u
}
