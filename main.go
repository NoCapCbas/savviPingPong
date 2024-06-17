package main

import (
  "flag"
  "fmt"
  "log"
  "os"

  "github.com/NoCapCbas/health-checker/healthchecker"
)

func main() {

  url := flag.String("url", "", "The URL to check the health of")
  flag.Parse()

  if *url == "" {
    fmt.Println("Usage: healthchecker -url <URL>")
    os.Exit(1)
  }

  status, err := healthchecker.HealthCheck(*url)
  if err != nil {
    log.Fatalf("Error checking health: %v\n", err)
  }

  fmt.Printf("Health check result for %s: %s\n", *url, status)

}
