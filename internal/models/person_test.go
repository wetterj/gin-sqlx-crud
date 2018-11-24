package models

import (
	"testing"

	"github.com/guregu/null"
	"github.com/wetterj/gin-sqlx-crud/internal/forms"
)

func makePtr(str string) *string {
	return &str
}

// Test any business logic here.
func TestPerson(t *testing.T) {
	p := Person{
		ID:        "0",
		FirstName: "James",
		Address:   null.StringFrom("an address"),
	}
	f := forms.CreatePerson{}
	f.FirstName = makePtr("Not James")
	f.LastName = makePtr("Wetter")
	p.ApplyForm(&f)

	if p.FirstName != "Not James" {
		t.Fatal("first name not updated")
	}
	if p.LastName != null.StringFrom("Wetter") {
		t.Fatal("last name not updated")
	}
	if p.Address.Valid {
		t.Fatal("address not updated")
	}
}
