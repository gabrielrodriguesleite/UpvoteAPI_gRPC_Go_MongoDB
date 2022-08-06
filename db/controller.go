package db

type Vote struct {
	Email string
	Score int
}

func Create(email string, score int) error {
	s := Vote{email, score}
	return Insert("vote", s)
}
