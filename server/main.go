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
curl -i http://127.0.0.1:8080/static/
curl -i http://127.0.0.1:8080/static/kml/example.kml
curl -i http://127.0.0.1:8080/static/asd.zip
*/

func main() {

	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)

	router, err := rest.MakeRouter(
		rest.Get("/test", test),
	)

	if err != nil {
		log.Fatal(err)
	}
	api.SetApp(router)

	http.Handle("/api/", http.StripPrefix("/api", api.MakeHandler()))
	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./assets/static"))))

	log.Fatal(http.ListenAndServe(":" + strconv.Itoa(app.PORT), nil))
}

func test(w rest.ResponseWriter, r *rest.Request) {
	w.WriteJson(map[string]string{
		"key": "value",
	})
}