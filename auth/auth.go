package auth

import (
	"journal/entity"
)


var profileDtb = make(entity.ProfileDatabase)

func generateKey(username, password string) string {
	return username + "#%$" + password
}

func Login(username, password string) (entity.Profile, error) {
	hashKey := generateKey(username, password)
	profile, err := profileDtb.SearchProfile(hashKey)
	if err != nil {
		return entity.Profile{}, entity.ErrWrongCredentials
	}

	return profile, nil
}

func Register(username, password, firstName, lastName string) error {
	hashKey := generateKey(username, password)
	_, err := profileDtb.SearchProfile(hashKey)

	if err == nil {	
		return entity.ErrDoubleRegistration
	}

	profileDtb.AddProfile(hashKey, entity.Profile{FirstName: firstName, LastName: lastName})
	return nil
}

func DeleteAccount(username, password string) error {
	hashKey := generateKey(username, password)
	_, err := profileDtb.SearchProfile(hashKey)

	if err != nil {
		return entity.ErrDeleteUnavailableProfile
	}

	profileDtb.DeleteProfile(hashKey)
	return nil
}