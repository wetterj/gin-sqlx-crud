package models

import (
	"time"

	"github.com/guregu/null"
	"github.com/wetterj/gin-sqlx-crud/internal/forms"
)

// This is the public data that should hide the SQL implementation
// details from the rest of the code.
type Person struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
	FirstName null.String
	LastName  null.String
	Address   null.String
	Age       null.Int
}

// The data mapping layer interface, again hiding implementation details.
type PersonService interface {
	Create(*forms.CreatePerson) (*Person, error)
	GetByID(string) (*Person, error)
	Update(*Person) error
	Delete(*Person) error
}
