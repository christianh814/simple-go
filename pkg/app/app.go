package app

import (
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// static port for now
const HttpPort string = "8080"

func Start() {
	//create/register a new request multiplexer
	router := mux.NewRouter()

	//Set up router for /
	router.HandleFunc("/", appRoot)
	router.HandleFunc("/info", appInfo)

	// try to start the app and log output
	log.Info("Starting server on port " + HttpPort)
	log.Fatal(http.ListenAndServe(":"+HttpPort, router))
}
