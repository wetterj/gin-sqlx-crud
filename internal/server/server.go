package server

import (
	"github.com/gin-gonic/gin"
	"github.com/wetterj/gin-sqlx-crud/internal/controllers"
	"github.com/wetterj/gin-sqlx-crud/internal/models"
	"github.com/wetterj/gin-sqlx-crud/internal/models/sql"
)

// This is the thing I will inject dependencies into.
type Server struct {
	PersonService models.PersonService
	Gin           *gin.Engine
}

// This creates a new server using environment variables to
// configure DB connection.
func NewServer() (*Server, error) {
	sql, err := sql.NewSQL()
	if err != nil {
		return nil, err
	}

	personService, err := sql.NewPersonService(sql)
	if err != nil {
		return nil, err
	}

	r := gin.Default()
	{
		route := r.Group("/person")
		ctrl := controllers.NewPerson(personService)
		route.POST("", ctrl.Post)
		route.PUT("/:id", ctrl.Put)
		route.GET("/:id", ctrl.Get)
		route.DELETE("/:id", ctrl.Delete)
	}

	return &Server{
		PersonService: personService,
		Gin:           r,
	}, nil
}
