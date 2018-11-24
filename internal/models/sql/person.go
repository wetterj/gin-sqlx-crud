package sql

import (
	"errors"

	"github.com/jmoiron/sqlx"
	"github.com/wetterj/gin-sqlx-crud/internal/forms"
	"github.com/wetterj/gin-sqlx-crud/internal/models"
)

type PersonService struct {
	conn *sqlx.DB
}

func NewPersonService(conn *sqlx.DB) (*PersonService, error) {
	// TODO: Make the tables
	return &PersonService{conn: conn}, nil
}

func (s *PersonService) Create(form *forms.CreatePerson) (*models.Person, error) {
	return nil, errors.New("Not made")
}

func (s *PersonService) Update(p *models.Person) error {
	return errors.New("Not made")
}

func (s *PersonService) GetByID(id string) (*models.Person, error) {
	return nil, errors.New("not made")
}

func (s *PersonService) Delete(p *models.Person) error {
	return errors.New("not made")
}

// Check it implements the interface
var _ models.PersonService = &PersonService{}
