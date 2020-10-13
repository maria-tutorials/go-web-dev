package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gopkg.in/mgo.v2/bson"

	"gopkg.in/mgo.v2"

	"../models"

	"github.com/julienschmidt/httprouter"
)

type UserController struct {
	session *mgo.Session
}

// NewUserController returns a *UserController, kinda like the express router with module.exports
func NewUserController(s *mgo.Session) *UserController {
	return &UserController{s}
}

// GetUser handles the HTTP GET requests for /user/:id
func (uc UserController) GetUser(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	oid := bson.ObjectIdHex(id)

	u := models.User{}

	if err := uc.session.DB("test").C("users").FindId(oid).One(&u); err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)

		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "ups something went wrong")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", uj)
}

// CreateUser handles the HTTP POST erequests for /user
func (uc UserController) CreateUser(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	u := models.User{}

	if err := json.NewDecoder(req.Body).Decode(&u); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// create id for db
	u.ID = bson.NewObjectId()
	// store in collection
	uc.session.DB("test").C("users").Insert(u)

	uj, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n", uj)
}

// DeleteUser handles the HTTP DELETE requests for /user/:id
func (uc UserController) DeleteUser(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	oid := bson.ObjectIdHex(id)

	if err := uc.session.DB("test").C("users").RemoveId(oid); err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Deleted user with id %s\n", oid)
}
