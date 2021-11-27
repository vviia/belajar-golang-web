package golangweb

import (
	"net/http"
	"testing"
)

func TestServer(t *testing.T) {
	server := http.Server{
		Addr: "Localhost:9000",
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
