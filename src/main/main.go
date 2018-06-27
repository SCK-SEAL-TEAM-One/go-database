package main

import (
	"net/http"

	"controller"
)

func main() {

	http.HandleFunc("/read", controller.Read)
	http.HandleFunc("/add", controller.Add)
	http.HandleFunc("/edit", controller.Edit)
	http.HandleFunc("/delete", controller.Delete)
	http.ListenAndServe(":8080", nil)
}
