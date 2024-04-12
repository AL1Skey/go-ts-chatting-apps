package seeder

import (
	"encoding/json"
	"os"
)

type User struct {

	// Username is the user's chosen name to login.
	Username string `json:"username" db:"username"`

	// Email is the user's email address.
	Email string `json:"email" db:"email"`

	// Password is the user's chosen password.
	Password string `json:"password" db:"password"`
}

func (s *Seeder) UsersSeed() {
	data, err := os.ReadFile("../data/users.json")
	if err != nil {
		panic(err)
	}

	var users []User
	err = json.Unmarshal(data, &users)
	if err != nil {
		panic(err)
	}

	for _, value := range users {
		stmt, _ := s.db.Prepare(`INSERT INTO customers(name, email) VALUES (?,?)`)
		_, err := stmt.Exec()
	}
}
