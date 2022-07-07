package main

import (
	"net/http"

	"example.com/controllers"
	"github.com/julienschmidt/httprouter"
	mgo "gopkg.in/mgo.v2"
)

func main() {

	r := httprouter.New()

	uc := controllers.NewUserController(getSession())

	r.GET("/user/:id", uc.GetUser)

	r.POST("/user", uc.CreateUser)

	r.DELETE("/user/:id", uc.DeleteUser)

	http.ListenAndServe(":8080", r)
}

func getSession() *mgo.Session {

	s, err := mgo.Dial("mongodb://localhost")

	if err != nil {
		panic(err)
	}

	return s

}
