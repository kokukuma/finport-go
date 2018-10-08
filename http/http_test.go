package http

import (
	"context"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kokukuma/finport-go/log"
)

func TestServerHandle(t *testing.T) {
	server, err := New(log.NewDiscard())
	if err != nil {
		t.Fatal(err.Error())
	}
	s := httptest.NewServer(server.helloWorld())

	//
	resp, err := getHello(s.URL)
	if err != nil {
		t.Fatal(err.Error())
	}
	if resp == "" {
		t.Fatal("response not found")
	}
}

func getHello(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func TestStartServer(t *testing.T) {
	server, err := New(log.NewDiscard())
	if err != nil {
		t.Fatal(err.Error())
	}

	// start server
	httpLn, err := net.Listen("tcp", fmt.Sprintf(":%d", 10001))
	if err != nil {
		t.Fatal(err.Error())
	}
	go server.Serve(httpLn)

	// check response
	resp, err := getHello("http://localhost:10001")
	if err != nil {
		t.Fatal(err.Error())
	}
	if resp == "" {
		t.Fatal("response not found")
	}

	// shutdown
	server.Shutdown(context.Background())

}
