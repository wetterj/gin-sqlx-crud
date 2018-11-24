package sql

import (
	"database/sql"

	"github.com/guregu/null"
	"github.com/jmoiron/sqlx"
	"github.com/wetterj/gin-sqlx-crud/internal/forms"
	"github.com/wetterj/gin-sqlx-crud/internal/models"
)

// PersonService is the implementation of the person data mapping layer
// using SQL.
type PersonService struct {
	conn *sqlx.DB
}

// NewPersonService creates the person service using the given
// connection pool to a postgres DB.
func NewPersonService(conn *sqlx.DB) (*PersonService, error) {
	// TODO: It would be better to use a DB management tool
	// to make migrations painless.
	_, err := conn.Exec(`
CREATE TABLE IF NOT EXISTS persons (
	id         SERIAL PRIMARY KEY,
	created_at TIMESTAMP NOT NULL DEFAULT NOW(),
	updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
	first_name TEXT NOT NULL,
	last_name  TEXT,
	address    TEXT,
	age        INTEGER
);
`)
	if err != nil {
		return nil, err
	}
	return &PersonService{conn: conn}, nil
}

// Create will try to add the person to the DB.
func (s *PersonService) Create(form *forms.CreatePerson) (*models.Person, error) {
	q := `
INSERT INTO persons(first_name, last_name, address, age)
VALUES ($1, $2, $3, $4)
RETURNING *;`

	var output models.Person
	err := s.conn.Get(
		&output,
		q,
		*form.FirstName,
		null.StringFromPtr(form.LastName),
		null.StringFromPtr(form.Address),
		null.IntFromPtr(form.Age),
	)

	if err != nil {
		return nil, err
	}
	return &output, nil
}

// Update will replace the values of the give person with those provided.
func (s *PersonService) Update(p *models.Person) error {
	if !validID(p.ID) {
		return models.ErrNotFound
	}

	q := `
UPDATE persons
SET updated_at = NOW(),
    first_name = $1,
    last_name = $2,
    address = $3,
    age = $4
WHERE id = $5;
`

	_, err := s.conn.Exec(
		q,
		p.FirstName,
		p.LastName,
		p.Address,
		p.Age,
		p.ID,
	)
	if err != nil {
		return err
	}
	return nil
}

// GetByID fetches the person with the given id.
func (s *PersonService) GetByID(id string) (*models.Person, error) {
	if !validID(id) {
		return nil, models.ErrNotFound
	}

	q := `
SELECT *
FROM persons
WHERE id = $1;`

	var output models.Person
	err := s.conn.Get(
		&output,
		q,
		id,
	)
	// Replace the SQL error with our own error type.
	if err == sql.ErrNoRows {
		return nil, models.ErrNotFound
	} else if err != nil {
		return nil, err
	} else {
		return &output, nil
	}
}

// Delete removes the person with the given id from the DB.
// TODO: this should just mark the object as deleted,
// not actually get rid of the data.
func (s *PersonService) Delete(id string) error {
	if !validID(id) {
		return models.ErrNotFound
	}

	q := `
DELETE FROM persons
WHERE id = $1;
`

	_, err := s.conn.Exec(
		q,
		id,
	)
	return err
}

// Check it implements the interface
var _ models.PersonService = &PersonService{}
