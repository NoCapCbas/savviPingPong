package main

import (
  "flag"
  "fmt"
  "log"
  "os"
)

func main() {

  url := flag.String("url", "", "The URL to check the health of")
  flag.Parse()

  if *url == "" {
    fmt.Println("Usage: healthchecker -url <URL>")
    os.Exit(1)
  }

  status, err := HealthCheck(*url)
  if err != nil {
    log.Fatalf("Error checking health: %v\n", err)
  }

  fmt.Printf("Health check result for %s: %s\n", *url, status)

}
