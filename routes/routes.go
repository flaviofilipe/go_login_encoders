package routes

import (
	"net/http"

	"github.com/flaviofilipe/go_login_encoders/controllers"
)

func Routes() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/md5", controllers.LoginMd5)
	http.HandleFunc("/senhaclara", controllers.LoginSenhaClara)
	http.HandleFunc("/hashsalt", controllers.LoginHashSalt)
}
