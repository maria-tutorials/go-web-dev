package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../models"

	"github.com/julienschmidt/httprouter"
)

type UserController struct{}

// NewUserController returns a *UserController, kinda like the express router with module.exports
func NewUserController() *UserController {
	return &UserController{}
}

// GetUser handles the HTTP GET requests for /user/:id
func (uc UserController) GetUser(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	u := models.User{
		ID:     p.ByName("id"),
		Name:   "FAKE USER",
		Gender: "fake",
		Age:    100,
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

	json.NewDecoder(req.Body).Decode(&u) //save it in a variable
	uj, _ := json.Marshal(&u)            //so we can send it back

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n", uj)
}

// DeleteUser handles the HTTP DELETE requests for /user/:id
func (uc UserController) DeleteUser(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	//TODO: will remove users from db by id
	w.WriteHeader(http.StatusNotImplemented)
}
