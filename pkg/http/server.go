package http

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	Address string
	Port    int
	Mux     *mux.Router
}

func (s *Server) Run() {
	fmt.Println("Server running on", fmt.Sprintf("%v:%v", s.Address, s.Port))

	http.ListenAndServe(fmt.Sprintf("%v:%v", s.Address, s.Port), s.Mux)
}

func NewServer(
	Address string,
	Port int,
	mux *mux.Router,
) *Server {
	return &Server{
		Address: Address,
		Port:    Port,
		Mux:     mux,
	}
}
