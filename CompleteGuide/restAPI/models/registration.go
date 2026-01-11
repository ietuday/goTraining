package models

import (
	"strings"
	"time"

	"example.com/rest-api/db"
)

type Registration struct {
	ID        int64     `json:"id"`
	EventID   int64     `json:"eventId"`
	UserID    int64     `json:"userId"`
	CreatedAt time.Time `json:"createdAt"`
}

// Register the logged-in user for an event
func RegisterForEvent(eventID, userID int64) error {
	query := `
	INSERT INTO registrations (event_id, user_id, created_at)
	VALUES (?, ?, ?)
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(eventID, userID, time.Now())
	if err != nil {
		// handle duplicate registration (unique constraint)
		if strings.Contains(strings.ToLower(err.Error()), "unique") {
			return ErrAlreadyRegistered
		}
		return err
	}

	return nil
}

var ErrAlreadyRegistered = sqlErr("already registered")

type sqlErr string
func (e sqlErr) Error() string { return string(e) }

func GetRegistrationsByEventID(eventID int64) ([]User, error) {
	query := `
	SELECT u.id, u.email
	FROM registrations r
	JOIN users u ON u.id = r.user_id
	WHERE r.event_id = ?
	ORDER BY r.created_at DESC
	`

	rows, err := db.DB.Query(query, eventID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make([]User, 0)
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Email); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, rows.Err()
}

func UnregisterForEvent(eventID, userID int64) (int64, error) {
	query := `DELETE FROM registrations WHERE event_id = ? AND user_id = ?`
	res, err := db.DB.Exec(query, eventID, userID)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}


