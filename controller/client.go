package controller

import (
	"net/http"
	"sync"
	"github.com/ant0ine/go-json-rest/rest"
)

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
