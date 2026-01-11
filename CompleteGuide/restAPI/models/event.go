package models

import (
	"fmt"
	"time"

	"example.com/rest-api/db"
)

type Event struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	DateTime    time.Time `json:"dateTime" binding:"required"`
	UserID      int       `json:"userId"`
}

func (e *Event) Save() error {
	fmt.Println("Inside  Save", e)
	query := `
	INSERT INTO events (name, description, location, dateTime, user_id)
	VALUES (?, ?, ?, ?, ?)
	`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	e.ID = id
	return nil
}

func GetAllEvents() ([]Event, error) {
	query := `
	SELECT id, name, description, location, dateTime, user_id
	FROM events
	ORDER BY dateTime
	`

	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	events := make([]Event, 0)

	for rows.Next() {
		var event Event
		if err := rows.Scan(
			&event.ID,
			&event.Name,
			&event.Description,
			&event.Location,
			&event.DateTime,
			&event.UserID,
		); err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return events, nil
}

func GetEventByID(id int64) (*Event, error) {
	query := `
	SELECT id, name, description, location, dateTime, user_id
	FROM events
	WHERE id = ?
	`
	var event Event

	err := db.DB.QueryRow(query, id).Scan(
		&event.ID,
		&event.Name,
		&event.Description,
		&event.Location,
		&event.DateTime,
		&event.UserID,
	)

	if err != nil {
		return nil, err
	}

	return &event, nil

}

func (e *Event) UpdateEvent() (int64, error) {
	query := `
	UPDATE events
	SET name = ?, description = ?, location = ?, dateTime = ?, user_id = ?
	WHERE id = ?
	`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID, e.ID)
	if err != nil {
		return 0, err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rows, nil
}

func (e *Event) Delete() (int64, error) {
	query := `DELETE FROM events WHERE id = ?`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(e.ID)
	if err != nil {
		return 0, err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rows, nil
}


