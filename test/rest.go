package main

import (
	"log"
	"net/http"
	"sync"
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

	users := Clients{
		Store: map[string]*Client{},
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

type Client struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Clients struct {
	sync.RWMutex
	Store map[string]*Client
}

func (u *Clients) Default(w rest.ResponseWriter, r *rest.Request) {
	// TODO homepage
	w.WriteHeader(http.StatusNotImplemented)
}

func (u *Clients) GetAll(w rest.ResponseWriter, r *rest.Request) {
	u.RLock()
	clients := make([]Client, len(u.Store))
	// TODO improve
	i := 0
	for _, client := range u.Store {
		clients[i] = *client
		i++
	}
	u.RUnlock()
	w.WriteJson(&clients)
}

func (u *Clients) Get(w rest.ResponseWriter, r *rest.Request) {
	id := r.PathParam("id")
	u.RLock()
	if u.Store[id] == nil {
		rest.NotFound(w, r)
		u.RUnlock()
		return
	}
	// TODO improve
	var client *Client
	client = &Client{}
	*client = *u.Store[id]
	u.RUnlock()
	w.WriteJson(client)
}

func (u *Clients) Post(w rest.ResponseWriter, r *rest.Request) {
	client := Client{}
	err := r.DecodeJsonPayload(&client)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if client.Id == "" {
		rest.Error(w, "[id] required", http.StatusBadRequest)
		return
	}

	u.Lock()
	if u.Store[client.Id] != nil {
		rest.Error(w, "[id] existed", http.StatusBadRequest)
		u.Unlock()
		return
	}
	u.Store[client.Id] = &client
	u.Unlock()
	w.WriteJson(&client)
}

func (u *Clients) Put(w rest.ResponseWriter, r *rest.Request) {
	// TODO homepage
	w.WriteHeader(http.StatusNotImplemented)
}

func (u *Clients) Patch(w rest.ResponseWriter, r *rest.Request) {
	client := Client{}
	err := r.DecodeJsonPayload(&client)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	id := r.PathParam("id")
	u.Lock()
	if u.Store[id] == nil {
		rest.NotFound(w, r)
		u.Unlock()
		return
	}

	client.Id = id
	u.Store[id] = &client
	u.Unlock()
	w.WriteJson(&client)
}

func (u *Clients) Delete(w rest.ResponseWriter, r *rest.Request) {
	id := r.PathParam("id")
	u.Lock()
	if u.Store[id] == nil {
		rest.NotFound(w, r)
		u.Unlock()
		return
	}
	delete(u.Store, id)
	u.Unlock()
	w.WriteHeader(http.StatusOK)
}
