package main

import (
	"fmt"
	"net/http"
	"time"
)

func getUser(w http.ResponseWriter, req *http.Request) user {
	var u user

	c, err := req.Cookie("session")
	if err != nil {
		return u //{}
	}

	c.MaxAge = sessionLenght
	http.SetCookie(w, c)

	// if the user exists already, get user
	if s, ok := dbSessions[c.Value]; ok {
		u = dbUsers[s.un]
	}
	return u
}

func alreadyLoggedIn(req *http.Request) bool {
	c, err := req.Cookie("session")
	if err != nil {
		return false
	}
	s := dbSessions[c.Value]
	_, ok := dbUsers[s.un] //cool
	return ok
}

func cleanSessions() {
	showMeTheSessions()

	for k, _ := range dbSessions {
		if time.Now().Sub(dbSessionsCleaned) > (time.Second * 30) {
			delete(dbSessions, k)
		}
	}

	dbSessionsCleaned = time.Now()
}

func showMeTheSessions() {
	for k, v := range dbSessions {
		fmt.Println(k, v.un, v.lastActive)
	}
}
