package main 

import (
  "fmt"
  "net/http"
  "time"
)

func HealthCheck(url string) (string, error) {
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
