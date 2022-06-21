package main

import (
	"context"
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/foo", Index)
	http.HandleFunc("/bar", bar)

	http.Handle("/", http.NotFoundHandler())

	http.ListenAndServe(":8080", nil)

}

func Index(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	ctx = context.WithValue(ctx, "userID", 7707)
	ctx = context.WithValue(ctx, "fnamee", "James Bond")

	results := dbAccess(ctx)

	fmt.Fprintln(w, results)
}

func dbAccess(ctx context.Context) int {
	return ctx.Value("userID").(int)
}

func bar(w http.ResponseWriter, req *http.Request) {

	context := req.Context()

	fmt.Fprintln(w, context)

}
