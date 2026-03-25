package unit

import (
	"os"
	"path/filepath"
	"testing"

	"k8s-cost-optimizer/pkg/config"
)

func TestLoadConfig(t *testing.T) {
	// Percorsi file di test temporanei
	dir := t.TempDir()
	cfgPath := filepath.Join(dir, "config.yaml")
	pricingPath := filepath.Join(dir, "pricing.yaml")

	// Scrivo file di test
	os.WriteFile(cfgPath, []byte(`
server:
  port: 8080
collector:
  scrape_interval: "15s"
  prometheus_url: "http://localhost:9090"
forecast:
  lookback_hours: 24
  anomaly_threshold: 1.5
`), 0644)

	os.WriteFile(pricingPath, []byte(`
cpu:
  price_per_core: 0.05
memory:
  price_per_gb: 0.01
`), 0644)

	cfg, err := config.LoadConfig(cfgPath, pricingPath)
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}

	if cfg.Server.Port != 8080 {
		t.Errorf("Expected port 8080, got %d", cfg.Server.Port)
	}
	if cfg.Cost.CPUPricePerCore != 0.05 {
		t.Errorf("Expected CPU price 0.05, got %f", cfg.Cost.CPUPricePerCore)
	}
}
