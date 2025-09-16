package user

import "time"

type User struct {
    ID                  int64      `db:"id"`
    Email               string     `db:"email"`
    Username            string     `db:"username"`
    FirstName           string     `db:"first_name"`
    LastName            string     `db:"last_name"`
    PasswordHash        string     `db:"password_hash"`
    IsActive            bool       `db:"is_active"`
    IsVerified          bool       `db:"is_verified"`
    VerificationToken   *string    `db:"verification_token"`
    ResetToken          *string    `db:"reset_token"`
    ResetTokenExpires   *time.Time `db:"reset_token_expires"`
    FailedLoginAttempts int        `db:"failed_login_attempts"`
    LockUntil           *time.Time `db:"lock_until"`
    LastLogin           *time.Time `db:"last_login"`
    CreatedAt           time.Time  `db:"created_at"`
    UpdatedAt           time.Time  `db:"updated_at"`
}

