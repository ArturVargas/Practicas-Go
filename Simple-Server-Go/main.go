package main

import (
	"net/http"
)

func main() {

	//rutas
	http.HandleFunc("/", homeRoute)
	http.HandleFunc("/contact", contactRoute)

	// iniciar server
	http.ListenAndServe(":3000", nil)
}

func homeRoute(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Primer Server en Go"))
}

func contactRoute(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Pagina de Contacto"))
}
