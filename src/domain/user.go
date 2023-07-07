package domain

type User struct {
	Email     string   `json:"email" binding:"required"`
	FirtsName string   `json:"firstName" binding:"required"`
	LastName  string   `json:"lastName" binding:"required"`
	Password  string   `json:"password"`
	IsActive  bool     `json:"is_active"`
}

type UserRepository interface {
	Create(user User) (*User, error)
	Update(user User) (*User, error)
	Delete(email string) error
	FindByEmail(email string) (*User, error)
	ChangePassword(email, password string) error
}

func GetUser(email string, repo UserRepository) (*User, error) {
	return repo.FindByEmail(email)
}

func CreateUser(user User, repo UserRepository) (*User, error) {
	return repo.Create(user)
}

func UpdateUser(user User, repo UserRepository) (*User, error) {
	return repo.Update(user)
}

func DeleteUser(email string, repo UserRepository) error {
	return repo.Delete(email)
}

func ChangePassword(email, password string, repo UserRepository) error {
	return repo.ChangePassword(email, password)
}
