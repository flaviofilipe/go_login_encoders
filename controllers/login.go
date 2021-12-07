package controllers

import (
	"net/http"
	"text/template"

	"github.com/flaviofilipe/go_login_encoders/models"
	"github.com/flaviofilipe/go_login_encoders/repositories"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

type Data struct {
	User       models.User
	Error      string
	FormAction string
	Title      string
}

func Index(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "Index", nil)
}

func LoginSenhaClara(w http.ResponseWriter, r *http.Request) {
	var data Data
	data.FormAction = "senhaclara"
	data.Title = "Login Senha Clara"

	if r.Method == "POST" {
		email, password := r.FormValue("email"), r.FormValue("password")
		user, errorMessage := repositories.FindUser(email, password, 0)

		if errorMessage != nil {
			data.Error = errorMessage.Error()
		}

		data.User = user
	}

	temp.ExecuteTemplate(w, "Login", data)
}

func LoginMd5(w http.ResponseWriter, r *http.Request) {
	var data Data
	data.FormAction = "md5"
	data.Title = "Login MD5"

	if r.Method == "POST" {
		email, password := r.FormValue("email"), r.FormValue("password")
		user, errorMessage := repositories.FindUser(email, password, 1)

		if errorMessage != nil {
			data.Error = errorMessage.Error()
		}

		data.User = user
	}

	temp.ExecuteTemplate(w, "Login", data)
}

func LoginHashSalt(w http.ResponseWriter, r *http.Request) {
	var data Data
	data.FormAction = "hashsalt"
	data.Title = "Login Hash e Salt"

	if r.Method == "POST" {
		email, password := r.FormValue("email"), r.FormValue("password")
		user, errorMessage := repositories.FindUser(email, password, 2)

		if errorMessage != nil {
			data.Error = errorMessage.Error()
		}

		data.User = user
	}

	temp.ExecuteTemplate(w, "Login", data)
}
