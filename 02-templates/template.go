package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// Estructuras
type User struct {
	Username string
	Age      int
	Active   bool
	Admin    bool
	Cursos   []Curso
}

type Curso struct {
	Nombre string
}

func Saludar(nombre string) string {
	return "Hola " + nombre + " desde una funcion"
}

var templates = template.Must(template.New("T").ParseGlob("templates/**/*.html"))
var errorTemplate = template.Must(template.ParseFiles("templates/error/error.html"))

func handlerError(rw http.ResponseWriter, status int) {
	rw.WriteHeader(status)
	errorTemplate.Execute(rw, nil)
}

func renderTemplate(rw http.ResponseWriter, name string, data interface{}) {
	err := templates.ExecuteTemplate(rw, name, data)

	if err != nil {
		//http.Error(rw, "No es posible renderizar la pagina", http.StatusBadRequest)
		handlerError(rw, http.StatusInternalServerError)
	}
}

func Index(rw http.ResponseWriter, r *http.Request) {
	c1 := Curso{"Go"}
	c2 := Curso{"Python"}
	c3 := Curso{"Java"}
	c4 := Curso{"Javascript"}
	//fmt.Fprintln(rw, "Hola Mundo")
	//template := template.Must(template.ParseFiles("index.html", "base.html"))
	/*funciones := template.FuncMap{
		"saludar": Saludar,
	}*/
	//template, error := template.New("index.html").Funcs(funciones).ParseFiles("index.html")
	cursos := []Curso{c1, c2, c3, c4}
	usuario := User{"Juan", 28, true, false, cursos}
	renderTemplate(rw, "index.html", usuario)
	//err := templates.ExecuteTemplate(rw, "index.html", usuario)

	//if err != nil {
	//	panic(err)
	//} //else {
	//template.Execute(rw, usuario)
	//}
}

func Registro(rw http.ResponseWriter, r *http.Request) {
	/*err := templates.ExecuteTemplate(rw, "registro.html", nil)

	if err != nil {
		panic(err)
	}*/
	renderTemplate(rw, "registro.html", nil)
}

func main() {
	//Mux
	mux := http.NewServeMux()
	mux.HandleFunc("/", Index)
	mux.HandleFunc("/registro", Registro)
	//Router
	/*http.HandleFunc("/", Hola)
	  http.HandleFunc("/page", PageNotFound)
	  http.HandleFunc("/error", Error)
	  http.HandleFunc("/saludar", Saludar)*/
	//Crear servidor
	server := &http.Server{
		Addr:    "localhost:3000",
		Handler: mux,
	}
	fmt.Println("El servidor esta corriendo en el puerto 3000")
	fmt.Println("Run server: http://localhost:3000/")
	//log.Fatal(http.ListenAndServe("localhost:3000", mux))
	log.Fatal(server.ListenAndServe())
}
