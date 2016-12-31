package client

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type Page struct {
	Name, Description string
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// fmt.Fprint(w, "Welcome!\n")
	t := Page{
		Name:        "Hello",
		Description: "Welcome to index page!",
	}
	renderTemplate(w, "index", "base", t)
}

// func CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
// 	r.ParseForm()
// 	email := r.PostFormValue("email")
// 	password := r.PostFormValue("password")
// 	fmt.Print(email)
// 	fmt.Print(password)
// 	http.Redirect(w, r, "/", 302)
// }
func RegisterRoutes(router *httprouter.Router) {
	//router.GET("/users/register", controllers.Register)
	router.GET("/", Index)

}
