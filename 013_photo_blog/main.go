package main

import (
	"crypto/sha1"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
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

	//display pictures without /public
	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("./public"))))

	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	c := getCookie(w, req)

	// process form submission
	if req.Method == http.MethodPost {
		mf, fh, err := req.FormFile("f")
		if err != nil {
			fmt.Println(err)
		}
		defer mf.Close()
		// create sha for file name from FileHeader
		ext := strings.Split(fh.Filename, ".")[1]
		h := sha1.New()
		io.Copy(h, mf)
		fname := fmt.Sprintf("%x", h.Sum(nil)) + "." + ext
		// create new file
		wd, err := os.Getwd() //dir
		if err != nil {
			fmt.Println(err)
		}
		path := filepath.Join(wd, "public", "pictures", fname)
		nf, err := os.Create(path)
		if err != nil {
			fmt.Println(err)
		}
		defer nf.Close()
		// copy
		mf.Seek(0, 0) //get head back to begining
		io.Copy(nf, mf)
		// add filename to this user's cookie
		c = appendValue(w, c, fname)
	}

	xs := strings.Split(c.Value, "|")
	tpl.ExecuteTemplate(w, "index.gohtml", xs[1:]) //remove session_id
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

//deprecated
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

func appendValue(w http.ResponseWriter, c *http.Cookie, fname string) *http.Cookie {
	s := c.Value
	if !strings.Contains(s, fname) {
		s += "|" + fname
	}
	c.Value = s
	http.SetCookie(w, c)
	return c
}
