package main

import (
	"fmt"
	"log"
	"net/http"
)

type MyServer struct {
	PanicCountdown int
}

func (m *MyServer) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	defer func() {
		m.PanicCountdown = 100
	}()
	if req.URL.Path != "/" {
		http.NotFound(w, req)
		return
	}

	if m.PanicCountdown == 0 {
		panic("Panic!!!")
	}

	fmt.Fprintf(w, "Panic at %d...", m.PanicCountdown)
	m.PanicCountdown--
}

func main() {
	h := &MyServer{
		PanicCountdown: 3,
	}
	log.Fatalln(http.ListenAndServe(":8080", h))
}
