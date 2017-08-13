package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	}
	r := mux.NewRouter()
	r.HandleFunc("/", ControladorInicio)

	// Colección de mensajes
	mensajes := r.Path("/mensajes").Subrouter()
	mensajes.Methods("GET").HandlerFunc(ControladorIndiceMensajes)
	mensajes.Methods("POST").HandlerFunc(ControladorCreaMensajes)

	// Mensaje singular
	mensaje := r.PathPrefix("/mensajes/{id}").Subrouter()
	mensaje.Methods("GET").Path("/editar").HandlerFunc(ControladorEditarMensaje)
	mensaje.Methods("GET").HandlerFunc(ControladorMostrarMensaje)
	mensaje.Methods("PUT", "POST").HandlerFunc(ControladorActualizaMensaje)
	mensaje.Methods("DELETE").HandlerFunc(ControladorEliminaMensaje)

	fmt.Println("Iniciando el servidor en %s ", port)
	http.ListenAndServe(port, r)
}

func ControladorInicio(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "Inicio")
}

func ControladorIndiceMensajes(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "índice de mensajes")
}

func ControladorCreaMensajes(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "crear mensajes")
}

func ControladorMostrarMensaje(rw http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	fmt.Fprintln(rw, "mostrando mensaje", id)
}

func ControladorActualizaMensaje(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "actualizando mensaje")
}

func ControladorEliminaMensaje(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "borrar mensaje")
}

func ControladorEditarMensaje(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "editar mensaje")
}
