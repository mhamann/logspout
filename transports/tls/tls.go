package tls

import (
	"crypto/tls"
	"net"

	"github.com/gliderlabs/logspout/adapters/raw"
	"github.com/gliderlabs/logspout/resolver"
	"github.com/gliderlabs/logspout/router"
)

func init() {
	router.AdapterTransports.Register(new(tlsTransport), "tls")
	// convenience adapters around raw adapter
	router.AdapterFactories.Register(rawTLSAdapter, "tls")
}

func rawTLSAdapter(route *router.Route) (router.LogAdapter, error) {
	route.Adapter = "raw+tls"
	return raw.NewRawAdapter(route)
}

type tlsTransport int

func (t *tlsTransport) Dial(addr string, options map[string]string) (net.Conn, error) {
	daddr, err := resolver.ResolveSrvAddr(resolver.DNSConfig{Addr: addr})
	if err != nil {
		return nil, err
	}

	conn, err := tls.Dial("tcp", daddr, nil)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
