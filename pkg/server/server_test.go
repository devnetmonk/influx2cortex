package server

import (
	"flag"
	"net/http"
	"testing"

	"github.com/go-kit/log"

	"github.com/gorilla/mux"

	"github.com/stretchr/testify/require"
)

// TestServerRun ensures that after initializing the server and configuring a route
// the server is able to handle http requests
func TestServerRun(t *testing.T) {
	var cfg Config
	cfg.RegisterFlags(flag.NewFlagSet("", flag.ExitOnError))
	cfg.HTTPListenPort = 0

	logger := log.NewNopLogger()

	server, err := NewServer(logger, cfg, mux.NewRouter(), nil)
	require.NoError(t, err)

	server.Router.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(204)
	})

	go server.Run()
	defer server.Shutdown(nil)

	req, err := http.NewRequest("GET", fmt.Sprintf("http://127.0.0.1:%d/test", cfg.HTTPListenPort), http.NoBody)
	require.NoError(t, err)
	_, err = http.DefaultClient.Do(req)
	require.NoError(t, err)
}
