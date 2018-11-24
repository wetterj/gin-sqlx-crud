package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/guregu/null"
	"github.com/wetterj/gin-sqlx-crud/internal/forms"
	"github.com/wetterj/gin-sqlx-crud/internal/models"
)

// Person provides the handlers for the person entity.
type Person struct {
	personService models.PersonService
}

// NewPerson created the controller using the given data mapper for
// persons.
func NewPerson(personService models.PersonService) (*Person, error) {
	return &Person{
		personService: personService,
	}, nil
}

// Post will create a new person from the given data, if the form is valid.
func (p *Person) Post(c *gin.Context) {
	var form forms.CreatePerson
	if c.ShouldBindWith(&form, binding.JSON) != nil {
		// TODO: Give a better error message.
		c.JSON(
			http.StatusNotAcceptable,
			gin.H{"message": "invalid data."},
		)
		c.Abort()
		return
	}

	person, err := p.personService.Create(&form)
	if err != nil {
		// TODO: An error middleware should log the error,
		// and email admin.
		c.Error(err)
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"message": "internal error"},
		)
		c.Abort()
		return
	}

	// TODO: use a view if part of the person data should not be
	// returned to the client.
	c.JSON(
		http.StatusCreated,
		person,
	)
}

// Patch will perform a partial update of a user.
func (p *Person) Patch(c *gin.Context) {
	var form forms.UpdatePerson
	if c.ShouldBindWith(&form, binding.JSON) != nil {
		// TODO: Give a better error message.
		c.JSON(
			http.StatusNotAcceptable,
			gin.H{"message": "invalid data."},
		)
		c.Abort()
		return
	}
	id := c.Param("id")

	person, err := p.personService.GetByID(id)
	if err == models.ErrNotFound {
		c.JSON(
			http.StatusNotFound,
			gin.H{"message": "user not found"},
		)
		c.Abort()
		return
	} else if err != nil {
		c.Error(err)
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"message": "internal error."},
		)
		c.Abort()
		return
	}

	if form.FirstName != nil {
		person.FirstName = *form.FirstName
	}
	if form.LastName != nil {
		person.LastName = null.StringFrom(*form.LastName)
	}
	if form.Address != nil {
		person.Address = null.StringFrom(*form.Address)
	}
	if form.Age != nil {
		person.Age = null.IntFrom(*form.Age)
	}

	err = p.personService.Update(person)
	if err != nil {
		c.Error(err)
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"message": "internal error."},
		)
		c.Abort()
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{"message": "updated"},
	)
}

// Get will fetch a person by ID.
func (p *Person) Get(c *gin.Context) {
	id := c.Param("id")
	person, err := p.personService.GetByID(id)
	if err == models.ErrNotFound {
		c.JSON(
			http.StatusNotFound,
			gin.H{"message": "user not found"},
		)
		c.Abort()
		return
	} else if err != nil {
		c.Error(err)
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"message": "internal error."},
		)
		c.Abort()
		return
	}

	c.JSON(
		http.StatusOK,
		person,
	)
}

// Delete will remove a person from the DB.
func (p *Person) Delete(c *gin.Context) {
	id := c.Param("id")

	err := p.personService.Delete(id)

	if err != nil {
		c.Error(err)
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"message": "internal error."},
		)
		c.Abort()
	}

	c.JSON(
		http.StatusOK,
		gin.H{"message": "deleted"},
	)
}
