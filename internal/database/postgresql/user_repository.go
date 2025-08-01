package postgresql

import "stargazer/internal/interfaces"

type UserRepository struct {
	*BaseRepository
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		BaseRepository: NewBaseRepository(),
	}
}

// Create Raw SQL
func (r *UserRepository) Create(user interfaces.UserInterface) error {
	query := `
	insert into users (id, username, email, avatar, password, nickname, description, created_at, updated_at)
	values (?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	return r.Raw(
		query,
		user.GetID(),
		user.GetUsername(),
		user.GetEmail(),
		user.GetAvatar(),
		user.GetPassword(),
		user.GetNickname(),
		user.GetDescription(),
		user.GetCreatedAt(),
		user.GetUpdatedAt(),
	).Scan(&user).Error
}
