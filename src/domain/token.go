package domain

type Token struct {
	UserEmail string `json:"email"`
	Value     string `json:"value"`
}

type TokenRepository interface {
	Create(token Token) (*Token, error)
	Delete(token Token) error
	FindByEmail(userEmail string) (*[]Token, error)
	DeleteAll(userEmail string) error
	ValidateToken(token, email string) (bool, error)
}

func CreateToken(token Token, repo TokenRepository) (*Token, error) {
	return repo.Create(token)
}

func DeleteToken(token Token, repo TokenRepository) error {
	return repo.Delete(token)
}

func GetToken(userEmail string, repo TokenRepository) (*[]Token, error) {
	return repo.FindByEmail(userEmail)
}

func DeleteAllTokens(userEmail string, repo TokenRepository) error {
	return repo.DeleteAll(userEmail)
}

func TokenIsValid(email, tokenString string, repo TokenRepository) (bool, error) {
	return repo.ValidateToken(tokenString, email)
}
