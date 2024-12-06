package entity

import "errors"

var (
	// Authentication Errors

	ErrWrongCredentials = errors.New("wrong credentials provided")
	ErrDoubleRegistration = errors.New("cannot register twice")
	ErrDeleteUnavailableProfile = errors.New("cannot delete a user who doesn't exists")
	ErrUserNameTaken = errors.New("the inputted username is already taken")
	ErrProfileNotFound = errors.New("could not find the specified profile")

	// Journal Errors
	ErrJournalNotFound = errors.New("could not find a journal record for the user")
	ErrJournalAlreadyCreated = errors.New("user already has a journal created")

)