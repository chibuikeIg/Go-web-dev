package main

import (
	"fmt"
	"net/http"
	"time"

	uuid "github.com/satori/go.uuid"
)

func GetUser(w http.ResponseWriter, r *http.Request) User {

	c, err := r.Cookie("session")

	if err != nil {
		sID := uuid.NewV4()

		c = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
	}

	c.MaxAge = sessionLength
	http.SetCookie(w, c)

	// fetch session

	var u User

	if s, ok := DBSessions[c.Value]; ok {
		s.lastActivity = time.Now()
		DBSessions[c.Value] = s
		u = DBUsers[s.un]
	}

	return u
}

func AlreadyLoggedIn(r *http.Request) bool {
	c, err := r.Cookie("session")

	if err != nil {
		return false
	}

	c.MaxAge = sessionLength

	s := DBSessions[c.Value]
	_, ok := DBUsers[s.un]

	return ok
}

func CleanSessions() {
	fmt.Println("BEFORE CLEAN")

	ShowSessions()

	for k, v := range DBSessions {
		if time.Now().Sub(v.lastActivity) > (time.Second * 30) {
			delete(DBSessions, k)
		}
	}
}

func ShowSessions() {
	fmt.Println("*************")

	for k, v := range DBSessions {
		fmt.Println(k, v.un)
	}

	fmt.Println("")
}
