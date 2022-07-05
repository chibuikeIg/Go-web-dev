package main

import (
	mgo "gopkg.in/mgo.v2"
)

func main() {

}

func getSession() *mgo.Session {

	s, err := mgo.Dial("mongodb://localhost")

	if err != nil {
		panic(err)
	}

	return s

}
