package userValidations

import (
	"errors"
	"net/http"

	"github.com/chandanaavadhani/usermanagement/models"
	"github.com/chandanaavadhani/usermanagement/userDB"

	_ "github.com/go-sql-driver/mysql"
)

func SignUpValidation(user models.User) (int, error) {
	if user.Username == " " || user.Username == "" {
		return http.StatusBadRequest, errors.New("username is Required")
	}
	if user.PWord == " " || user.PWord == "" {
		return http.StatusBadRequest, errors.New("password is Required")
	}
	if len(user.PWord) < 8 {
		return http.StatusBadRequest, errors.New("password is Not Valid")
	}
	if user.ConfirmPassword != user.PWord {
		return http.StatusBadRequest, errors.New("passwords don't match")
	}
	count, _, err := userDB.Validationsquery(user)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	if count > 0 {
		return http.StatusBadRequest, errors.New("user already exist")
	}
	return 200, nil
}

func LoginValidation(user models.User) (int, error) {
	if user.Username == " " || user.Username == "" {
		return http.StatusBadRequest, errors.New("username is Required")
	}
	count, password, err := userDB.Validationsquery(user)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	if count == 0 {
		return http.StatusBadRequest, errors.New("username not found")
	}
	if password != user.PWord {
		return http.StatusBadRequest, errors.New("passwords don't match")
	}
	return 200, nil
}

func UpdateValidation(user models.User) (int, error) {
	if user.Username == " " || user.Username == "" {
		return http.StatusBadRequest, errors.New("username is Required")
	}
	count, password, err := userDB.Validationsquery(user)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	if count == 0 {
		return http.StatusBadRequest, errors.New("username not found")
	}
	if user.OldPassword != password {
		return http.StatusBadRequest, errors.New("old password is wrong")
	}
	if len(user.NewPassword) < 8 {
		return http.StatusBadRequest, errors.New("password is Not Valid")
	}
	if user.ConfirmPassword != user.NewPassword {
		return http.StatusBadRequest, errors.New("passwords don't match")
	}
	return 200, nil
}

func DeleteUsername(user string) (int, error) {
	count, _, err := userDB.StringValidations(user)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	if count == 0 {
		return http.StatusBadRequest, errors.New("username not found")
	}
	return 200, nil
}
