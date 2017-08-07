package postgres

import (
	"database/sql"

	"github.com/bbjj040471/transporter/client"
)

var _ client.Session = &Session{}

// Session serves as a wrapper for the underlying *sql.DB
type Session struct {
	pqSession *sql.DB
	db        string
}
