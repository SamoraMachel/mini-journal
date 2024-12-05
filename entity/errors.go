package entity

import "errors"

var (
	ErrWrongCredentials = errors.New("wrong credentials provided")
	ErrDoubleRegistration = errors.New("cannot register twice")
	ErrDeleteUnavailableProfile = errors.New("cannot delete a user who doesn't exists")
)