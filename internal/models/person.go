package models

import (
	"time"

	"github.com/guregu/null"
	"github.com/wetterj/gin-sqlx-crud/internal/forms"
)

// Person is the public data that should hide the SQL implementation
// details from the rest of the code.
type Person struct {
	ID        string
	CreatedAt time.Time   `db:"created_at"`
	UpdatedAt time.Time   `db:"updated_at"`
	FirstName string      `db:"first_name"`
	LastName  null.String `db:"last_name"`
	Address   null.String
	Age       null.Int
}

// This updates the person data to match what is in the form.
func (p *Person) ApplyForm(form *forms.CreatePerson) {
	p.FirstName = *form.FirstName
	p.LastName = null.StringFromPtr(form.LastName)
	p.Address = null.StringFromPtr(form.Address)
	p.Age = null.IntFromPtr(form.Age)
}

// PersonService is the data mapping layer interface, again hiding implementation details.
type PersonService interface {
	Create(*forms.CreatePerson) (*Person, error)
	GetByID(string) (*Person, error)
	Update(*Person) error
	Delete(string) error
}
