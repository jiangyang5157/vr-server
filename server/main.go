package main

import (
	"log"
	"net/http"
	"github.com/ant0ine/go-json-rest/rest"
	"github.com/jiangyang5157/vr-server/app"
	"strconv"
)

/*
curl -i http://127.0.0.1:8080/api/patch
curl -i http://127.0.0.1:8080/static/
curl -i http://127.0.0.1:8080/static/asd.txt
*/

func main() {

	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)

	router, err := rest.MakeRouter(
		rest.Get("/patch", patch),
	)

	if err != nil {
		log.Fatal(err)
	}
	api.SetApp(router)

	http.Handle("/api/", http.StripPrefix("/api", api.MakeHandler()))
	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./assets"))))

	log.Fatal(http.ListenAndServe(":" + strconv.Itoa(app.PORT), nil))
}

func patch(w rest.ResponseWriter, r *rest.Request) {
	w.WriteJson(map[string]string{
		"sequence": strconv.Itoa(app.PATCH_SEQUENCE),
		"path": "asd.zip",
	})
}