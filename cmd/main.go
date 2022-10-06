package main

import (
	"net/http"

	router "gitlab.okymikhael.io/playground/parking-lot-golang/server"
)

func main() {
	router.Router()
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
