package models

import "database/sql"

type Session struct {
	ID     int
	UserID int
	// token is only set when the session is created. when look up a session
	// tis will be left emptry, as we only store the hash of a seesion token
	// in our database and we cannot reverse it into a raw token.
	Token     string
	TokenHash string
}

type SessionService struct {
	DB *sql.DB
}

func (ss *SessionService) Create(userID int) (*Session, error) {
	// 1. Create the session token
	// tODO
	return nil, nil
}

func (ss *SessionService) User(token string) (*User, error) {
	return nil, nil
}
