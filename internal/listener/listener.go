package listener

import (
	"flag"
	"log"
	"net"

	"golang.org/x/net/netutil"
)

const (
	defaultBindAddr = ":8080"
	defaultMaxConn  = 20
)

func GetHttpListener() (net.Listener, error){
	var (
		bindAddr string
		maxConn  uint64
	)

	flag.StringVar(&bindAddr, "b", defaultBindAddr, "TCP address the server will bind to")
	flag.Uint64Var(&maxConn, "c", defaultMaxConn,
		"maximum number of client connections the server will accept, 0 means unlimited")
	flag.Parse()

	listener, err := net.Listen("tcp", bindAddr)
	if err != nil {
		return nil, err
	}

	if maxConn > 0 {
		listener = netutil.LimitListener(listener, int(maxConn))
		log.Println("max connections set to: ", maxConn)
	}

	return listener, nil
}
