package main

import (
	"crypto/sha1"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	uuid "github.com/satori/go.uuid"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("template/*"))
}

func main() {

	http.HandleFunc("/", Index)

	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("./public"))))

	http.Handle("/favicon.ico", http.NotFoundHandler())

	http.ListenAndServe(":8080", nil)
}

func Index(w http.ResponseWriter, r *http.Request) {

	c := GetCookie(w, r)

	if r.Method == http.MethodPost {

		UploadImage(w, r, c)

	}

	images := strings.Split(c.Value, "|")

	tpl.ExecuteTemplate(w, "index.html", images)
}

func GetCookie(w http.ResponseWriter, req *http.Request) *http.Cookie {

	c, err := req.Cookie("session")

	if err != nil {
		id := uuid.NewV4()

		c = &http.Cookie{
			Name:  "session",
			Value: id.String(),
		}

		http.SetCookie(w, c)
	}

	return c

}

func UploadImage(w http.ResponseWriter, req *http.Request, cookie *http.Cookie) {

	mf, fh, err := req.FormFile("img")

	if err != nil {
		fmt.Println(err)
	}

	defer mf.Close()

	fnameArr := strings.Split(fh.Filename, ".")
	ext := fnameArr[len(fnameArr)-1]
	h := sha1.New()
	io.Copy(h, mf)

	fname := fmt.Sprintf("%x", h.Sum(nil)) + "." + ext

	wd, err := os.Getwd()

	if err != nil {
		fmt.Println(err)
	}

	path := filepath.Join(wd, "public", "pics", fname)

	nf, err := os.Create(path)

	if err != nil {
		fmt.Println(err)
	}

	defer nf.Close()

	mf.Seek(0, 0)
	io.Copy(nf, mf)

	if strings.Contains(cookie.Value, fname) == false {
		cookie.Value = cookie.Value + "|" + fname
		http.SetCookie(w, cookie)
	}
}
