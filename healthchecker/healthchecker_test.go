package healthchecker

import (
    "testing"
)

func TestHealthCheckSite(t *testing.T) {
  url := "https://example.com"

  status, err := HealthCheckSite(url)
  if err != nil {
    t.Errorf("Error checking site health: %v", err)
  }

  t.Logf("Site health check result for %s: %s", url, status)
}

func TestHealthCheckPort(t *testing.T) {
    host := "example.com"
    port := "80" // Replace with the port you want to test

    status, err := HealthCheckPort(host, port)
    if err != nil {
        t.Errorf("Error checking port health: %v", err)
    }
    
    t.Logf("Port health check result for %s:%s: %s", host, port, status)
}

