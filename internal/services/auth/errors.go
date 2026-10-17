package auth_services

import "errors"

var ErrInvalidCredentials = errors.New("invalid email or password")
var ErrEmailAlreadyExists = errors.New("email already exists")