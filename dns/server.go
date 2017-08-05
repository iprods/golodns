package dns

import (
	"net"

	"github.com/miekg/dns"
)

type Server struct {
	Address   string
	Port      string
	ResolveIp net.IP
}

var server *dns.Server

// Start listening
func (s *Server) Start() error {
	addr := s.Address + ":" + s.Port
	handler := RequestHandler{
		ResolveIp: s.ResolveIp,
	}
	server = &dns.Server{
		Addr: addr,
		Net: "udp",
		Handler: dns.HandlerFunc(handler.handleRequest),
	}
	return server.ListenAndServe()
}
