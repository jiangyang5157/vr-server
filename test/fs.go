package main

import (
	"log"
	"net/http"
	"github.com/ant0ine/go-json-rest/rest"
)

/*
curl -i http://127.0.0.1:8080/api/message

curl -i http://127.0.0.1:8080/static/
curl -i http://127.0.0.1:8080/static/asd.txt
*/

func main() {

	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)

	router, err := rest.MakeRouter(
		rest.Get("/message", func(w rest.ResponseWriter, req *rest.Request) {
			w.WriteJson(map[string]string{"Body": "Hello World!"})
		}),
	)

	if err != nil {
		log.Fatal(err)
	}
	api.SetApp(router)

	http.Handle("/api/", http.StripPrefix("/api", api.MakeHandler()))
	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./assets"))))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
