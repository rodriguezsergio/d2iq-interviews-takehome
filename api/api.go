package api

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
	"github.com/rodriguezsergio/d2iq-interviews-takehome/lru"
	"github.com/rs/zerolog/log"
)

type Response struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Server struct {
	MaxItems int
	MuLock   sync.Mutex
	Cache    map[string]string
	LRU      *lru.DoublyLinkedList
}

func respondWithError(w http.ResponseWriter, statusCode int, message string) {
	respondWithJSON(w, statusCode, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, statusCode int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		log.Error().Err(err).Msg("")
		respondWithError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(response)
}

func (s *Server) AddKey(w http.ResponseWriter, r *http.Request) {
	s.MuLock.Lock()
	defer s.MuLock.Unlock()

	payload := map[string]string{}
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		log.Error().Err(err).Msg("")
		respondWithError(w, http.StatusBadRequest, "bad input provided")
		return
	}

	for key, val := range payload {
		if _, ok := s.Cache[key]; ok {
			log.Debug().Msgf("Moving key, %v, to the front of the LRU cache.", key)
			s.LRU.RemoveNode(key)
			s.LRU.AddNode(key)
		} else {
			s.LRU.AddNode(key)
			s.Cache[key] = val
		}
	}

	if s.LRU.Size > s.MaxItems {
		keyToEvict := s.LRU.End.Data
		s.LRU.RemoveNode(keyToEvict)
		delete(s.Cache, keyToEvict)
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"status": "key(s) successfully added"})
}

func (s *Server) GetKey(w http.ResponseWriter, r *http.Request) {
	s.MuLock.Lock()
	defer s.MuLock.Unlock()

	vars := mux.Vars(r)
	key := vars["key"]

	if value, ok := s.Cache[key]; ok {
		log.Debug().Msgf("Moving key, %v, to the front of the LRU cache.", key)
		s.LRU.RemoveNode(key)
		s.LRU.AddNode(key)

		respondWithJSON(w, http.StatusOK, &Response{Key: key, Value: value})
		return
	}
	respondWithError(w, http.StatusNotFound, "key not found")
}
