package main

import (
  "github.com/bilbof/shiny/server"
  "log"
)

func main() {
  s := server.Server {
    Port: ":1718",
    BackendUrl: "http://localhost:8000",
  }

  err := s.ListenAndServe()

  if err != nil {
      log.Fatal("ListenAndServe:", err)
  }
}
