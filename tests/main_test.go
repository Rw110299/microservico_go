package main

import (
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestPingHandler(t *testing.T) {
    req, err := http.NewRequest("GET", "/ping", nil)
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(pingHandler)
    handler.ServeHTTP(rr, req)

    if rr.Body.String() != "pong\n" {
        t.Errorf("Esperado 'pong', recebido '%s'", rr.Body.String())
    }
}
