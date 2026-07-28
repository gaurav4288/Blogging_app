package repository

import (
	"database/sql"
	"errors"

	"github.com/gaurav4288/go_tutorial/blogging_app/internal/models"
	"github.com/gaurav4288/go_tutorial/blogging_app/internal/pkg/database"
)

type (
	UserRepository interface {
		GetUserByID(id string) (*models.User, error)
		CreateUser(user *models.User) error
		UpdateUser(user *models.User) error
		DeleteUser(id string) error
		IsUserExistsById(id string) (bool, error)
	}

	user struct {
		db *database.Database
	}
)

// NewUserRepository returns a UserRepository backed by the given DB connection
func NewUserRepository(db *database.Database) UserRepository {
	return &user{db: db}
}

func (u *user) GetUserByID(id string) (*models.User, error) {
	query := `SELECT user_id, firstname, lastname, email, role, created_at FROM users WHERE user_id = $1`

	var usr models.User
	err := u.db.Conn.QueryRow(query, id).Scan(
		&usr.UserId,
		&usr.FirstName,
		&usr.LastName,
		&usr.Email,
		&usr.Role,
		&usr.CreatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &usr, nil
}

func (u *user) CreateUser(user *models.User) error {
	query := `INSERT INTO users (user_id, firstname, lastname, email, password, created_at)
	          VALUES ($1, $2, $3, $4, $5, $6)`

	_, err := u.db.Conn.Exec(query,
		user.UserId,
		user.FirstName,
		user.LastName,
		user.Email,
		user.Password,
		user.CreatedAt,
	)
	return err
}

func (u *user) UpdateUser(user *models.User) error {
	query := `UPDATE users SET firstname = $1, lastname = $2, email = $3, password = $4 WHERE user_id = $5`

	result, err := u.db.Conn.Exec(query,
		user.FirstName,
		user.LastName,
		user.Email,
		user.Password,
		user.UserId,
	)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return errors.New("user not found")
	}

	return nil
}

func (u *user) DeleteUser(id string) error {
	query := `DELETE FROM users WHERE user_id = $1`

	result, err := u.db.Conn.Exec(query, id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return errors.New("user not found")
	}

	return nil
}

func (u *user) IsUserExistsById(id string) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM users WHERE user_id = $1)`
	var exists bool
	err := u.db.Conn.QueryRow(query, id).Scan(&exists)
	return exists, err
}
