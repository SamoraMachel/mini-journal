package test

import (
	"testing"
	"journal/entity"
	"journal/auth"
	"journal/lib"
)


const (
	testUserName = "test"
	testPassword = "test"
	testFirstName = "test"
	testLastName = "test"
)


func TestRegistration(t *testing.T) {
	t.Run("register user successfully", func(t *testing.T) {
		err := auth.Register(testUserName, testPassword, testFirstName, testLastName)
		lib.AssertNoError(t, err)
	})

	t.Run("prevent registering with same username", func(t *testing.T) {
		err := auth.Register(testUserName, "random_password", testFirstName, testLastName)
		lib.AssertError(t, err, entity.ErrUserNameTaken)
	})

	t.Run("prevent registering user more than once", func(t *testing.T) {
		err := auth.Register(testUserName, testPassword, testFirstName, testLastName)
		lib.AssertError(t, err, entity.ErrDoubleRegistration)
	})
}

func TestLogin(t *testing.T) {
	t.Run("on right credentials", func(t *testing.T) {
		userProfile := entity.Profile{FirstName: testFirstName, LastName: testLastName}
		auth.Register(testUserName, testPassword, testFirstName, testLastName)

		profile, _ := auth.Login(testUserName, testPassword)

		if userProfile != profile {
			t.Errorf("expected %v got %v", userProfile, profile)
			
		}
	})

	t.Run("on wrong credentials", func(t *testing.T) {
		username := "wrong_username"
		password := "wrong_password"

		_, err := auth.Login(username, password)

		lib.AssertError(t, err, entity.ErrWrongCredentials)
	})

}

func TestDeleteProfile(t *testing.T) {
	t.Run("deleting an available account", func(t *testing.T) {
		err := auth.DeleteAccount(testUserName, testPassword)
		lib.AssertNoError(t, err)
	})

	t.Run("deleting an unavailable account", func(t *testing.T) {
		err := auth.DeleteAccount(testUserName, testPassword)
		lib.AssertError(t, err, entity.ErrDeleteUnavailableProfile)
	})
}


