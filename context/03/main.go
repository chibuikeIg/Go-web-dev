package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func main() {

	http.HandleFunc("/", Index)
	http.HandleFunc("/bar", bar)

	http.Handle("/favicon.ico", http.NotFoundHandler())

	http.ListenAndServe(":8080", nil)

}

func Index(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	ctx = context.WithValue(ctx, "userID", 7707)
	ctx = context.WithValue(ctx, "fnamee", "James Bond")

	results, err := dbAccess(ctx)

	if err != nil {
		http.Error(w, err.Error(), http.StatusRequestTimeout)
	}

	fmt.Fprintln(w, results)
}

func dbAccess(ctx context.Context) (int, error) {

	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)

	defer cancel()
	ch := make(chan int)

	go func() {

		// ridiculous long task

		uuid := ctx.Value("userID").(int)

		time.Sleep(10 * time.Second)

		if ctx.Err() != nil {
			return
		}

		ch <- uuid

	}()

	select {
	case <-ctx.Done():
		return 0, ctx.Err()
	case i := <-ch:
		return i, ctx.Err()
	}

}

func bar(w http.ResponseWriter, req *http.Request) {

	context := req.Context()

	fmt.Fprintln(w, context)

}
