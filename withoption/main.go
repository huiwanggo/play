package main

import "fmt"

func main() {
	server, err := New("127.0.0.1", WithProtocol("http"), WithPort(80))
	if err != nil {
		return
	}
	fmt.Printf("%#v", server)
}

type Server struct {
	Addr     string
	Port     int
	Protocol string
}

// Option Functional Options
type Option func(s *Server)

func New(addr string, opts ...Option) (*Server, error) {
	s := &Server{
		Addr: addr,
	}

	for _, opt := range opts {
		opt(s)
	}
	return s, nil
}

func WithPort(p int) Option {
	return func(s *Server) {
		s.Port = p
	}
}

func WithProtocol(p string) Option {
	return func(s *Server) {
		s.Protocol = p
	}
}
