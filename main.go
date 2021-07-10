package main

import (
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/kelseyhightower/envconfig"
	"github.com/rodriguezsergio/d2iq-interviews-takehome/api"
	"github.com/rodriguezsergio/d2iq-interviews-takehome/lru"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type EnvVars struct {
	LogLevel  string `split_words:"true" default:"info"`
	CacheSize int    `split_words:"true" default:"250"`
}

func main() {
	// add file and line number to logs
	log.Logger = log.With().Caller().Logger()

	// read env vars
	var env EnvVars
	err := envconfig.Process("", &env)
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("Failed to process environment variables")
	}

	// configure logging based on env var
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if strings.ToLower(env.LogLevel) == "debug" {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	// initialize server struct
	ll := &lru.DoublyLinkedList{}

	s := api.Server{
		Cache:    make(map[string]string),
		LRU:      ll,
		MaxItems: env.CacheSize,
	}

	// set up endpoints
	r := mux.NewRouter()
	r.HandleFunc("/key/{key}", s.GetKey).
		Methods("GET")
	r.HandleFunc("/key", s.AddKey).
		Methods("POST")

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal().
			Err(err).
			Msg("Failed to start server")
	}
}
