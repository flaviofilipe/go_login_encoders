package repositories

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"

	"github.com/flaviofilipe/login/models"
	"github.com/flaviofilipe/login/utils"
)

const (
	SenhaClara      = 0
	Md5             = 1
	HashSalt        = 2
	PasswordDefault = "admin"
)

func FindUser(email, password string, method int) (models.User, error) {
	var user models.User
	var errorMessage error
	switch method {
	case SenhaClara:
		user, errorMessage = validarSenhaClara(email, password)
	case Md5:
		user, errorMessage = validarMd5(email, password)
	case HashSalt:
		user, errorMessage = validarHashSalt(email, password)
	default:
		errorMessage = fmt.Errorf("Método inválido")
	}
	return user, errorMessage
}

func validarSenhaClara(email, password string) (models.User, error) {
	var user models.User
	var errorMessage error

	userByEmail := findUserByEmail(email)

	if password == userByEmail.Password {
		user = userByEmail
	} else {
		errorMessage = fmt.Errorf("Senha inválida")
	}

	return user, errorMessage
}

func validarMd5(email, password string) (models.User, error) {
	var user models.User
	var errorMessage error

	userByEmail := findUserByEmail(email)

	userPasswordEncoded := utils.GetMD5Hash(userByEmail.Password)
	passwordEncodedMd5 := utils.GetMD5Hash(password)

	if passwordEncodedMd5 == userPasswordEncoded {
		user = userByEmail
	} else {
		errorMessage = fmt.Errorf("Senha inválida")
	}

	return user, errorMessage
}

func validarHashSalt(email, password string) (models.User, error) {
	var user models.User
	var errorMessage error
	userByEmail := findUserByEmail(email)
	passwordEncodedMd5, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)

	if err != nil {
		log.Println(err)
	}

	err = bcrypt.CompareHashAndPassword(passwordEncodedMd5, []byte(userByEmail.Password))

	if err == nil {
		user = userByEmail
	} else {
		errorMessage = fmt.Errorf("Senha inválida")
	}

	return user, errorMessage
}

func findUserByEmail(email string) models.User {
	return models.User{Email: email, Password: PasswordDefault}
}
