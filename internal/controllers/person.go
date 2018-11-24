package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wetterj/gin-sqlx-crud/internal/models"
)

type Person struct {
	PersonService models.PersonService
}

func NewPerson(personService models.PersonService) (*Person, error) {
	return &Person{
		PersonService: personService,
	}, nil
}

func (p *Person) Post(c *gin.Context) {
	c.JSON(
		http.StatusNotImplemented,
		gin.H{"message": "not implemented"},
	)
	c.Abort()
}

func (p *Person) Put(c *gin.Context) {
	c.JSON(
		http.StatusNotImplemented,
		gin.H{"message": "not implemented"},
	)
	c.Abort()
}

func (p *Person) Get(c *gin.Context) {
	c.JSON(
		http.StatusNotImplemented,
		gin.H{"message": "not implemented"},
	)
	c.Abort()
}

func (p *Person) Delete(c *gin.Context) {
	c.JSON(
		http.StatusNotImplemented,
		gin.H{"message": "not implemented"},
	)
	c.Abort()
}
