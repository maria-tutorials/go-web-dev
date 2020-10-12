package main

import (
	"html/template"
	"net/http"
	"strings"

	uuid "github.com/satori/go.uuid"
)

type user struct {
	UserName string
	Password []byte
	First    string
	Last     string
	Role     string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	c := getCookie(w, req)
	c = appendMultipleValues(w, c)
	xs := strings.Split(c.Value, "|")
	tpl.ExecuteTemplate(w, "index.gohtml", xs)
}

func getCookie(w http.ResponseWriter, req *http.Request) *http.Cookie {
	c, err := req.Cookie("session")

	if err != nil {
		sID, _ := uuid.NewV4()
		c = &http.Cookie{
			Name:     "session",
			Value:    sID.String(),
			HttpOnly: true,
			Path:     "/",
		}
		http.SetCookie(w, c)
	}
	return c
}

func appendMultipleValues(w http.ResponseWriter, c *http.Cookie) *http.Cookie {
	p1 := "pic1.jpg"
	p2 := "pic2.jpg"
	p3 := "pic3.jpg"
	ps := []string{c.Value, p1, p2, p3}

	if !strings.Contains(c.Value, p1) { // cheating because I know I always set them all together gg
		s := strings.Join(ps, "|")
		c.Value = s

		http.SetCookie(w, c)
	}

	return c
}
