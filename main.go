package main

import (
	"fmt"
	"net"

	"github.com/spf13/pflag"
)

var url string

func main() {
	var port = pflag.Int("port", 3000, "Server port")
	var origin = pflag.String("origin", "", "usage string")
	pflag.Parse()

	url = *origin

	server(*port)
}

func server(port int) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic("womp womp! the server didn't listen.")
	}

	println("listening on port: ", port)
	// God bless goroutines. Imagine doing this with C++ threads lmao xd!
	for {
		request, err := listener.Accept()
		if err != nil {
			println("womp womp! could not accept request")
			continue
		}
		go proxy(request)
	}
}

func proxy(request net.Conn) {
	// Now this is where I'll deal with the request:
	// 	Should check if request destination is cached
	// 	If not cached:
	// 		Should forward the request to the origin url (url global variable)
	// 		Should cache the respose for X amount of time (I'll figure out X)
	// 		Should return the response from origin to the client that requested it.
	// 	If cached:
	// 		Should return the cached data.
	// 	Ps: Need to add a header to the response to tell it was gotten from the cache or not. (this is the hard part)
	println("Got a request;")
	return
}
