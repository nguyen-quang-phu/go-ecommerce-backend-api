package repo

type IUserRepository interface {
	GetUserByEmail(email string) bool
}

type userRepository struct{}

// GetUserByEmail implements IUserRepository.
func (u userRepository) GetUserByEmail(email string) bool {
	panic("unimplemented")
}

func NewUserRepository() IUserRepository {
	return userRepository{}
}
