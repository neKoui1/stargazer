package interfaces

type UserRepository interface {
	Create(user UserInterface) error
	GetByID(id string) (UserInterface, error)
	GetByUsername(username string) (UserInterface, error)
	GetByEmail(email string) (UserInterface, error)
	Update(user UserInterface) error
	Delete(id string) error
	List(offset, limit int) ([]UserInterface, error)
	Count() (int64, error)
}
