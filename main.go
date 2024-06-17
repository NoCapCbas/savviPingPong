package main

import (
  "flag"
  "fmt"
  "log"
  "os"
  "strconv"

  "github.com/NoCapCbas/health-checker/healthchecker"
)

func main() {
  // CLI flags
  url := flag.String("url", "", "The URL to check the health of")
  port := flag.Int("port", 0, "The port number to check the health of")
  host := flag.String("host", "", "The host to check the health of")
  flag.Parse()

  if *url == "" && (*port == 0 && *host == "") {
    fmt.Println("Usage: healthchecker -url <URL> OR -port <PORT> -host <HOST>")
    os.Exit(1)
  }
  
  if *url != "" {
    status, err := healthchecker.HealthCheckSite(*url)
    if err != nil {
      log.Fatalf("Error checking health: %v\n", err)
    }

    fmt.Printf("Site Health check result for %s: %s\n", *url, status)
  } else if *port != 0 && *host != "" {
    portStr := strconv.Itoa(*port)
    status, err := healthchecker.HealthCheckPort(*host, portStr)
    if err != nil {
      log.Fatalf("Error checking port health: %v\n", err)
    }
    fmt.Printf("Port health check result for %s:%s: %s\n", *host, portStr, status)
  }

}
