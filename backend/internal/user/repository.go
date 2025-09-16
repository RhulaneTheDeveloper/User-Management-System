package user

import (
	"database/sql"
	"errors"
	"time"
)

// DBType enum
type DBType int

const (
	MySQL DBType = iota
	PostgreSQL
)

// UserRepository handles DB operations
type UserRepository struct {
	DB     *sql.DB
	DBType DBType
}

// CreateUser inserts a new user into the DB
func (repo *UserRepository) CreateUser(u *User) error {
	query := `
	INSERT INTO users 
	(email, username, first_name, last_name, password_hash, is_active, is_verified, created_at, updated_at)
	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	if repo.DBType == PostgreSQL {
		query = `
		INSERT INTO users 
		(email, username, first_name, last_name, password_hash, is_active, is_verified, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		`
	}

	_, err := repo.DB.Exec(query,
		u.Email,
		u.Username,
		u.FirstName,
		u.LastName,
		u.PasswordHash,
		u.IsActive,
		u.IsVerified,
		u.CreatedAt,
		u.UpdatedAt,
	)

	return err
}

// GetUserByEmail retrieves a user by email
func (repo *UserRepository) GetUserByEmail(email string) (*User, error) {
	query := `
	SELECT id, email, username, first_name, last_name, password_hash, is_active, is_verified, created_at, updated_at, last_login, failed_login_attempts 
	FROM users WHERE email = ?
	`

	if repo.DBType == PostgreSQL {
		query = `
		SELECT id, email, username, first_name, last_name, password_hash, is_active, is_verified, created_at, updated_at, last_login, failed_login_attempts
		FROM users WHERE email = $1
		`
	}

	row := repo.DB.QueryRow(query, email)
	u := &User{}
	err := row.Scan(
		&u.ID,
		&u.Email,
		&u.Username,
		&u.FirstName,
		&u.LastName,
		&u.PasswordHash,
		&u.IsActive,
		&u.IsVerified,
		&u.CreatedAt,
		&u.UpdatedAt,
		&u.LastLogin,
		&u.FailedLoginAttempts,
	)
	if err != nil {
		return nil, err
	}

	return u, nil
}

// UpdateUser updates user info (like last login)
func (repo *UserRepository) UpdateUser(u *User) error {
	query := `
	UPDATE users SET username=?, first_name=?, last_name=?, password_hash=?, is_active=?, is_verified=?, updated_at=?, last_login=?, failed_login_attempts=? WHERE id=?
	`

	if repo.DBType == PostgreSQL {
		query = `
		UPDATE users SET username=$1, first_name=$2, last_name=$3, password_hash=$4, is_active=$5, is_verified=$6, updated_at=$7, last_login=$8, failed_login_attempts=$9 WHERE id=$10
		`
	}

	_, err := repo.DB.Exec(query,
		u.Username,
		u.FirstName,
		u.LastName,
		u.PasswordHash,
		u.IsActive,
		u.IsVerified,
		time.Now(),
		u.LastLogin,
		u.FailedLoginAttempts,
		u.ID,
	)

	return err
}

// DeleteUser deletes a user by ID
func (repo *UserRepository) DeleteUser(userID int64) error {
	query := `DELETE FROM users WHERE id=?`
	if repo.DBType == PostgreSQL {
		query = `DELETE FROM users WHERE id=$1`
	}

	res, err := repo.DB.Exec(query, userID)
	if err != nil {
		return err
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("no user found with given ID")
	}

	return nil
}

// ListUsers retrieves all users
func (repo *UserRepository) ListUsers() ([]*User, error) {
	query := `
	SELECT id, email, username, first_name, last_name, password_hash, is_active, is_verified, created_at, updated_at, last_login, failed_login_attempts
	FROM users
	`

	rows, err := repo.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*User
	for rows.Next() {
		u := &User{}
		err := rows.Scan(
			&u.ID,
			&u.Email,
			&u.Username,
			&u.FirstName,
			&u.LastName,
			&u.PasswordHash,
			&u.IsActive,
			&u.IsVerified,
			&u.CreatedAt,
			&u.UpdatedAt,
			&u.LastLogin,
			&u.FailedLoginAttempts,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	return users, nil
}

