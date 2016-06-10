package main

import (
	"log"
	"net/http"
	"github.com/jiangyang5157/vr-server/controller"
	"github.com/ant0ine/go-json-rest/rest"
)

/*
Using "localhost" or "127.0.0.1" as the hostname

Demo:
curl -i -H "Content-Type:application/json" -d '{"name":"Antoine"}' http://127.0.0.1:8080/users
curl -i -H "Content-Type:application/json" -d '{"id":"1"}' http://127.0.0.1:8080/users
curl -i -H "Content-Type:application/json" -d '{"id":"2","name":"Antoine2"}' http://127.0.0.1:8080/users
curl -i -H "Content-Type:application/json" -d '{"id":"1a2a3a4a5a"}' http://127.0.0.1:8080/users
curl -i -H "Content-Type:application/json" -d '{"id":"1a2a3a4a5a"}' http://127.0.0.1:8080/users
curl -i http://127.0.0.1:8080/users/12345
curl -i http://127.0.0.1:8080/users/1a2a3a4a5a
curl -i http://127.0.0.1:8080/users
curl -i -H "Content-Type:application/json" -X PATCH -d '{"name":"After Modify"}' http://127.0.0.1:8080/users/0
curl -i -H "Content-Type:application/json" -X PATCH -d '{"name":"After Modify"}' http://127.0.0.1:8080/users/1a2a3a4a5a
curl -i http://127.0.0.1:8080/users
curl -i -X DELETE http://127.0.0.1:8080/users/0
curl -i -X DELETE http://127.0.0.1:8080/users/1
curl -i http://127.0.0.1:8080/users

config port forwarding, then replace 127.0.0.1 by real ip, then:
curl -i http://122.62.240.22:8080/users
 */
func main() {
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)

	users := controller.Clients{
		Store: map[string]*controller.Client{},
	}

	router, err := rest.MakeRouter(
		// SELECT
		rest.Get("/", users.Default),
		rest.Get("/users", users.GetAll),
		rest.Get("/users/:id", users.Get),
		//CREATE
		rest.Post("/users", users.Post),
		// full UPDATE
		rest.Put("/users/:id", users.Put),
		// patch UPDATE
		rest.Patch("/users/:id", users.Patch),
		// DELETE
		rest.Delete("/users/:id", users.Delete),
	)

	if err != nil {
		log.Fatal(err)
	}
	api.SetApp(router)
	log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))
}