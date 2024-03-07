package mysql

import (
	"codes/pkg/models"
	"database/sql"
	"errors"
)

type AppointmentModel struct {
	DB *sql.DB
}

func (m *AppointmentModel) Insert(doctorType, userName, phoneNumber string) (int, error) {
	stmt := `INSERT INTO appointments (doctorType, userName, phoneNumber, created)
             VALUES(?, ?, ?, UTC_TIMESTAMP())`
	result, err := m.DB.Exec(stmt, doctorType, userName, phoneNumber)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (m *AppointmentModel) Get(id int) (*models.Appointment, error) {
	stmt := `SELECT id, doctorType, userName, phoneNumber, created FROM appointments
WHERE id = ?`

	row := m.DB.QueryRow(stmt, id)

	s := &models.Appointment{}

	err := row.Scan(&s.ID, &s.Doctortype, &s.Username, &s.Phonenumber, &s.Created)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}
	return s, nil
}

func (m *AppointmentModel) GetAll() ([]*models.Appointment, error) {
	stmt := `SELECT id, doctorType, userName, phoneNumber, created FROM appointments ORDER BY created DESC`

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	appointments := []*models.Appointment{}

	for rows.Next() {
		s := &models.Appointment{}
		err = rows.Scan(&s.ID, &s.Doctortype, &s.Username, &s.Phonenumber, &s.Created)

		if err != nil {
			return nil, err
		}

		appointments = append(appointments, s)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return appointments, nil
}

//func (m *SnippetModel) LastCategory(category string) ([]*models.Snippet, error) {
//	stmt := `SELECT id, title, content, created,expires,  category FROM
//   snippets WHERE category = ? ORDER BY created DESC LIMIT 10`
//
//	rows, err := m.DB.Query(stmt, category)
//	if err != nil {
//		return nil, err
//	}
//
//	defer rows.Close()
//
//	snippets := []*models.Snippet{}
//	for rows.Next() {
//		s := &models.Snippet{}
//
//		err = rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires, &s.Category)
//		if err != nil {
//			return nil, err
//		}
//		snippets = append(snippets, s)
//	}
//
//	if err = rows.Err(); err != nil {
//		return nil, err
//	}
//
//	return snippets, nil
//}
