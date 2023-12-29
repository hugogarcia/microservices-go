package data

import (
	"database/sql"
	"time"
)

type ProsgresTestRepository struct {
	Conn *sql.DB
}

func NewPostgresTestRepository(db *sql.DB) *ProsgresTestRepository {
	return &ProsgresTestRepository{Conn: db}
}

var usersMock = []*User{
	{ID: 1, Email: "teste", FirstName: "teste", LastName: "teste", Password: "teste", Active: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()},
	{ID: 2, Email: "teste2", FirstName: "teste2", LastName: "teste2", Password: "XXXXX", Active: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()},
}

// GetAll returns a slice of all users, sorted by last name
func (u *ProsgresTestRepository) GetAll() ([]*User, error) {
	return usersMock, nil
}

// GetByEmail returns one user by email
func (u *ProsgresTestRepository) GetByEmail(email string) (*User, error) {
	return usersMock[1], nil
}

// GetOne returns one user by id
func (u *ProsgresTestRepository) GetOne(id int) (*User, error) {
	return usersMock[0], nil
}

// Update updates one user in the database, using the information
// stored in the receiver u
func (u *ProsgresTestRepository) Update(user User) error {
	return nil
}

// DeleteByID deletes one user from the database, by ID
func (u *ProsgresTestRepository) DeleteByID(id int) error {
	return nil
}

// Insert inserts a new user into the database, and returns the ID of the newly inserted row
func (u *ProsgresTestRepository) Insert(user User) (int, error) {
	return 1, nil
}

// ResetPassword is the method we will use to change a user's password.
func (u *ProsgresTestRepository) ResetPassword(password string, user User) error {
	return nil
}

// PasswordMatches uses Go's bcrypt package to compare a user supplied password
// with the hash we have stored for a given user in the database. If the password
// and hash match, we return true; otherwise, we return false.
func (u *ProsgresTestRepository) PasswordMatches(plainText string, user User) (bool, error) {
	return true, nil
}
