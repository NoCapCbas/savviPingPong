package healthchecker

import (
  "fmt"
  "net/http"
  "time"
  "net"
)

func HealthCheckSite(url string) (string, error) {
  // Health Check Website/App via HTTP GET
  client := &http.Client{
    Timeout: 30 * time.Second, 
  }

  resp, err := client.Get(url)
  if err != nil {
    return "", fmt.Errorf("could not make GET request: %w", err)
  }

  defer resp.Body.Close()

  if resp.StatusCode == http.StatusOK {
    return "Healthy", nil
  }

  return fmt.Sprintf("Unhealthy (status code: %d)", resp.StatusCode), nil
}

func HealthCheckPort(host string, port string) (string, error) {
  // Health Check Port via TCP
  addr := fmt.Sprintf("%s:%s", host, port)
  timeout := 5 * time.Second

  conn, err := net.DialTimeout("tcp", addr, timeout)
  if err != nil {
    return "", fmt.Errorf("could not connect to %s:%s %w", host, port, err)
  }

  defer conn.Close()

  return "Port is reachable", nil
}


