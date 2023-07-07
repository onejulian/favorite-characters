package domain

type Character struct {
	UserEmail   string `json:"email"`
	IdCharacter string `json:"id_character"`
}

type CharacterRepository interface {
	Create(character Character) (*Character, error)
	Delete(userEmail string, idCharacter string) error
	FindByEmail(userEmail string) ([]*Character, error)
	DeleteAll(userEmail string) error
}

func CreateCharacter(character Character, repo CharacterRepository) (*Character, error) {
	return repo.Create(character)
}

func DeleteCharacter(userEmail string, idCharacter string, repo CharacterRepository) error {
	return repo.Delete(userEmail, idCharacter)
}

func GetCharacters(userEmail string, repo CharacterRepository) ([]*Character, error) {
	return repo.FindByEmail(userEmail)
}

func DeleteAllCharacters(userEmail string, repo CharacterRepository) error {
	return repo.DeleteAll(userEmail)
}
