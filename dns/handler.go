package dns

import (
	"fmt"
	"log"
	"net"

	"github.com/miekg/dns"
)

type RequestHandler struct {
	ResolveIp net.IP
}

// Handle the requests
func (h *RequestHandler) handleRequest(w dns.ResponseWriter, r *dns.Msg) {
	q := r.Question[0]

	info := fmt.Sprintf("Question: Type=%s Class=%s Name=%s", dns.TypeToString[q.Qtype], dns.ClassToString[q.Qclass], q.Name)

	if q.Qtype == dns.TypeA && q.Qclass == dns.ClassINET {
		m := new(dns.Msg)
		m.SetReply(r)
		a := new(dns.A)
		a.Hdr = dns.RR_Header{Name: q.Name, Rrtype: dns.TypeA, Class: dns.ClassINET, Ttl: 600}
		a.A = h.ResolveIp
		m.Answer = []dns.RR{a}
		w.WriteMsg(m)
		log.Printf("%s (RESOLVED)\n", info)
	} else {
		m := new(dns.Msg)
		m.SetReply(r)
		m.Rcode = dns.RcodeNameError // NXDOMAIN
		w.WriteMsg(m)
		log.Printf("%s (NXDOMAIN)\n", info)
	}
}
