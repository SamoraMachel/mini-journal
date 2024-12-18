package auth

import (
	"journal/entity"
)


var profileDtb = make(entity.ProfileDatabase)
var userNameList = entity.UserNameList{}

func GenerateKey(username, password string) string {
	return username + "#%$" + password
}

func Login(username, password string) (entity.Profile, error) {
	hashKey := GenerateKey(username, password)
	profile, err := profileDtb.SearchProfile(hashKey)
	if err != nil {
		return entity.Profile{}, entity.ErrWrongCredentials
	}

	return profile, nil
}

func Register(username, password, firstName, lastName string) error {
	hashKey := GenerateKey(username, password)
	_, err := profileDtb.SearchProfile(hashKey)

	if err == nil {	
		return entity.ErrDoubleRegistration
	}

	isUserNameValid := userNameList.IsValidUserName(username)
	
	if !isUserNameValid {
		return entity.ErrUserNameTaken
	}

	profileDtb.AddProfile(hashKey, entity.Profile{FirstName: firstName, LastName: lastName})
	userNameList = append(userNameList, username)
	return nil
}

func DeleteAccount(username, password string) error {
	hashKey := GenerateKey(username, password)
	_, err := profileDtb.SearchProfile(hashKey)

	if err != nil {
		return entity.ErrDeleteUnavailableProfile
	}

	profileDtb.DeleteProfile(hashKey)
	return nil
}