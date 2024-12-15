package db

import (
	"errors"
	models "myapp/models"
	service "myapp/service"

	"github.com/lib/pq"
)

var (
	ErrAlreadyExist = errors.New("email or username already exists")
	ErrServer       = errors.New("internal server error")
)

func (d *DB) CreateUser(req models.RegisterRequest) (string, error) {
	d.Connect()
	defer d.CloseConnection()

	hashedPassword, err := service.HashPassword(req.Password)

	if err != nil {
		return "", err
	}

	_, err = d.connection.Exec("INSERT INTO users (email, username, password_hash, created_at) VALUES ($1, $2, $3, NOW())", req.Email, req.Username, hashedPassword)

	if err, ok := err.(*pq.Error); ok {
		switch err.Code {
		case "23505":
			if err.Constraint == "users_email_key" || err.Constraint == "users_username_key" {
				return "", ErrAlreadyExist
			} else {
				return "", ErrServer
			}
		default:
			return "", ErrServer
		}
	}

	token, err := service.GenerateToken(req.Username)

	if err != nil {
		return "", err
	}

	err = service.SendEmail(req.Email, token)

	if err != nil {
		return "", err
	}

	return token, nil

}

func (d *DB) ConfirmMail(username string) error {
	d.Connect()
	defer d.CloseConnection()

	_, err := d.connection.Exec("UPDATE users SET email_confirmed = true WHERE username = $1", username)

	return err
}
