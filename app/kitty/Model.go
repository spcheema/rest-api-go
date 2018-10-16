package kitty

type Kitty struct {
	Name 		string `json:"name"`
	Breed 		string `json:"breed"`
	BirthDate 	string `json:"birthDate"`
}

func NewKitty(name string, breed string, birthDate string) *Kitty {
	return &Kitty{
		Name:		name,
		Breed:		breed,
		BirthDate:	birthDate,
	}
}