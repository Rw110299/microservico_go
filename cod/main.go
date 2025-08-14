package main

import (
    "fmt"
    "log"
    "net/http"
)

func pingHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "pong")
}

func main() {
    mux := http.NewServeMux()

    // Endpoint /ping
    mux.HandleFunc("/ping", pingHandler)

    // Servir arquivos estáticos da pasta "web"
    fs := http.FileServer(http.Dir("./web"))
    mux.Handle("/", fs)

    fmt.Println("Servidor rodando em http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", mux))
}
