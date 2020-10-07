package main

import (
	"html/template"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

type user struct {
	UserName string
	First    string
	Last     string
	Loggedin bool
}

var dbUsers = map[string]user{}      // user ID, user
var dbSessions = map[string]string{} // session ID, user ID

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/other", other)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	oreo, err := req.Cookie("session")
	if err != nil {
		sID, _ := uuid.NewV4()
		oreo = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(w, oreo)
	}

	var u user
	if un, ok := dbSessions[oreo.Value]; ok {
		u = dbUsers[un]
	}

	//process the form
	if req.Method == http.MethodPost {
		m := req.FormValue("username")
		f := req.FormValue("firstname")
		l := req.FormValue("lastname")
		li := req.FormValue("loggedin") == "on"

		u = user{m, f, l, li}

		dbSessions[oreo.Value] = m
		dbUsers[m] = u
	}

	tpl.ExecuteTemplate(w, "index.gohtml", u)
}

func other(w http.ResponseWriter, req *http.Request) {
	oreo, err := req.Cookie("session")
	if err != nil {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	un, ok := dbSessions[oreo.Value] //always get zero value even if entry does not exist
	if !ok {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	user := dbUsers[un]

	tpl.ExecuteTemplate(w, "other.gohtml", user)
}

func setCookie(w http.ResponseWriter, req *http.Request) {
	oreo, err := req.Cookie("session")
	if err != nil {
		sID, _ := uuid.NewV4()
		oreo = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(w, oreo)
	}
}
