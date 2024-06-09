package main

import (
	"fmt"
	"net/http"
  "log"
)

func main ()  {
  fmt.Println("Starting Signaling Server...")
  if err := http.ListenAndServe(":8080", nil); err != nil {
    log.Fatal(err)
  }
}
