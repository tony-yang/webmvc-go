package main

import (
  "net/http"
  "webmvc"
)

func main() {
  webmvc.Debug("Starting the WebMVC Go Framework")
  addr := ":80"
  server := webmvc.CreateNewServer()
  if err := http.ListenAndServe(addr, server); err != nil {
    webmvc.Critical("The WebMVC Go Framework failed on port 80:", err)
  }
}
