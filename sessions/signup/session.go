package main

import (
	"net/http"

	uuid "github.com/satori/go.uuid"
)

func getUser(w http.ResponseWriter, r *http.Request) User {

	c, err := r.Cookie("session")

	if err != nil {
		sID, _ := uuid.NewV4()

		c = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
	}

	http.SetCookie(w, c)

	// fetch session

	var u User

	if un, ok := DBSessions[c.Value]; ok {
		u = DBUsers[un]
	}

	return u
}

func alreadyLoggedIn(r *http.Request) bool {
	c, err := r.Cookie("session")

	if err != nil {
		return false
	}

	un := DBSessions[c.Value]
	_, ok := DBUsers[un]

	return ok
}
