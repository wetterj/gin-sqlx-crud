package sql

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// NewSQL creates and SQL connection using environment variables
// to configure.
func NewSQL() (*sqlx.DB, error) {
	host := strings.TrimSpace(os.Getenv("POSTGRES_HOST"))
	port := strings.TrimSpace(os.Getenv("POSTGRES_PORT"))
	user := strings.TrimSpace(os.Getenv("POSTGRES_USER"))
	password := strings.TrimSpace(os.Getenv("POSTGRES_PASSWORD"))
	db := strings.TrimSpace(os.Getenv("POSTGRES_DB"))

	info := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host,
		port,
		user,
		password,
		db,
	)
	return sqlx.Connect("postgres", info)
}

// validID checks if the given string is a valid id.
func validID(id string) bool {
	_, err := strconv.Atoi(id)
	return err == nil
}
