package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Server struct {
	cache *Cache
}

func NewServer(c *Cache) *Server {
	return &Server{
		cache: c,
	}
}

func (s *Server) getHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	key := r.Form.Get("key")
	val := s.cache.Get(key)
	json.NewEncoder(w).Encode(&val)
}

func (s *Server) addHandler(w http.ResponseWriter, r *http.Request) {
	// Get our values from the form
	r.ParseForm()
	key := r.Form.Get("key")
	val := r.Form.Get("val")
	s.cache.Put(key, val)

	// Marshal a json representation to a slice of bytes
	b, err := json.Marshal([]*update{
		&update{
			Action: "add",
			Data: map[string]string{
				key: val,
			},
		},
	})

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// create a signal channel and broadcast our update
	notify := make(chan struct{})
	broadcasts.QueueBroadcast(&broadcast{
		msg:    append([]byte("d"), b...),
		notify: notify,
	})
	<-notify
	log.Println("successfully broadcast add to memberlist")
}

func (s *Server) delHandler(w http.ResponseWriter, r *http.Request) {
	// Get our values from the form
	r.ParseForm()
	key := r.Form.Get("key")
	s.cache.Delete(key)

	// Marshal a json representation to a slice of bytes
	b, err := json.Marshal([]*update{
		&update{
			Action: "del",
			Data: map[string]string{
				key: "",
			},
		},
	})

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// create a signal channel and broadcast our update
	notify := make(chan struct{})
	broadcasts.QueueBroadcast(&broadcast{
		msg:    append([]byte("d"), b...),
		notify: notify,
	})
	<-notify
	log.Println("successfully broadcast delete to memberlist")
}
