package main

import (
	"github.com/GeertJohan/go.rice"
	"github.com/gorilla/mux"
	"net/http"
)

type SupervisorWebgui struct {
	router     *mux.Router
	supervisor *Supervisor
}

func NewSupervisorWebgui(supervisor *Supervisor) *SupervisorWebgui {
	router := mux.NewRouter()
	return &SupervisorWebgui{router: router, supervisor: supervisor}
}

func (sw *SupervisorWebgui) CreateHandler() http.Handler {
	sw.router.PathPrefix("/").Handler(http.FileServer(rice.MustFindBox("webgui").HTTPBox()))
	return sw.router
}
