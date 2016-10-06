package main

import (
	"log"
	"net/http"
	"github.com/ant0ine/go-json-rest/rest"
	"github.com/jiangyang5157/vr-server/app"
	"strconv"
)

/*
curl -i http://127.0.0.1:8080/api/test
curl -i http://192.168.1.67:8080/api/test

curl -i http://127.0.0.1:8080/assets/
curl -i http://127.0.0.1:8080/assets/static/
curl -i http://127.0.0.1:8080/assets/static.zip
curl -i http://127.0.0.1:8080/assets/static/layer/example.kml
*/

func main() {

	//RESTful server
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	router, err := rest.MakeRouter(
		rest.Get("/test", func(w rest.ResponseWriter, r *rest.Request) {
			w.WriteJson(map[string]string{
				"Key": "Hello World!",
			})
		}),
	)
	if err != nil {
		log.Fatal(err)
	}
	api.SetApp(router)
	http.Handle("/api/", http.StripPrefix("/api", api.MakeHandler()))

	// file server
	http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("./assets"))))
	log.Fatal(http.ListenAndServe(":" + strconv.Itoa(app.PORT), nil))
}

