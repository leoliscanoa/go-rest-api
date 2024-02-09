package controller

import (
	"fmt"
	"net/http"
)

func Index(writer http.ResponseWriter, request *http.Request) {
	_, err := fmt.Fprintf(writer, "Welcome to Usersß API!")
	if err != nil {
		panic(err)
	}
}
