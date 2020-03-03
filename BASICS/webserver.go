package main

import (
  "net/http"
  "io"
)

func main() {
  http.HandleFunc("/", servePage)
	http.ListenAndServe(":7781", nil)
}

func servePage(writer http.ResponseWriter, reqest *http.Request) {
  io.WriteString(writer, "This is a simple web server")
}
