// forms.go
package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type ContactDetails struct {
	Email   string
	Subject string
	Message string
}

func main() {
	tmpl := template.Must(template.ParseFiles("forms.html"))
	db, err := sql.Open("mysql", "root:localhost@tcp(127.0.0.1:3307)/testdb")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	fmt.Println("Successfully connected to mysql database")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}

		details := ContactDetails{
			Email:   r.FormValue("email"),
			Subject: r.FormValue("subject"),
			Message: r.FormValue("message"),
		}

		// do something with details
		//fmt.Printf("\n%s, %s, %s", details.Email, details.Subject, details.Message)
		insert, err := db.Query("INSERT INTO users VALUES(?,?,?)", details.Email, details.Subject, details.Message)
		if err != nil {
			panic(err.Error())
		}
		defer insert.Close()
		tmpl.Execute(w, struct{ Success bool }{true})
	})

	http.ListenAndServe(":8080", nil)
}
