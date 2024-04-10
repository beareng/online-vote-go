package main

import (
	// TODO: use sql for the first version and will migration to DynamoDB
	"database/sql"
	"net/http"

	"./backend"
	"./frontend"
)

func main() {
	// TODO: feel free to update the DB connection for now. will migration to dynamoDB later.
	db, err := sql.Open("sqlite3", "elections.sqlite")
	if nil != err {
		panic(err)
	}
	edb, err := backend.ConnectDatabase(db)
	if nil != err {
		panic(err)
	}

	mux := http.NewServeMux()
	frontend.Frontend{Edb: edb}.BindServeMux(mux, "")
	edb.BindServeMux(mux, "")
	http.ListenAndServe(":8080", mux)
}
