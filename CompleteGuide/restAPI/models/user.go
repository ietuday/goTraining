package models

import (
	"database/sql"
	"strings"

	"example.com/rest-api/db"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int64  `json:"id"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

// Save creates a user in DB (stores hashed password, not plain password).
func (u *User) Save() error {
	u.Email = strings.TrimSpace(strings.ToLower(u.Email))

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	query := `INSERT INTO users (email, password) VALUES (?, ?)`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(u.Email, string(passwordHash))
	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	u.ID = id
	u.Password = "" // never keep plain password in memory longer than needed
	return nil
}

func (u *User) ValidateCredential() error {
	u.Email = strings.TrimSpace(strings.ToLower(u.Email))
	query := `SELECT id, password FROM users WHERE email = ?`
	var id int64
	var passwordHash string
	err := db.DB.QueryRow(query, u.Email).Scan(&id, &passwordHash)

	if err != nil {
		// includes sql.ErrNoRows
		return err
	}

	// compare hash with plain password
	if err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(u.Password)); err != nil {
		// wrong password -> treat as "not found" to avoid leaking which field was wrong
		return sql.ErrNoRows
	}

	u.ID = id
	u.Password = "" // don't keep plain password
	return nil
}
