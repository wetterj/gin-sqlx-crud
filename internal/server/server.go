package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// This is the thing I will inject dependencies into.
type Server struct {
	// User Service
	Gin *gin.Engine
}

func NewServer() (*Server, error) {
	return nil, fmt.Errorf("not made")
}
