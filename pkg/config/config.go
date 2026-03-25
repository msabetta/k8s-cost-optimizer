package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type ServerConfig struct {
	Port int `yaml:"port"`
}

type CollectorConfig struct {
	ScrapeInterval string `yaml:"scrape_interval"`
	PrometheusURL  string `yaml:"prometheus_url"`
}

type ForecastConfig struct {
	LookbackHours     int     `yaml:"lookback_hours"`
	AnomalyThreshold  float64 `yaml:"anomaly_threshold"`
}

type CostConfig struct {
	CPUPricePerCore float64 `yaml:"cpu_price_per_core"`
	MemPricePerGB   float64 `yaml:"memory_price_per_gb"`
}

type AppConfig struct {
	Server    ServerConfig    `yaml:"server"`
	Collector CollectorConfig `yaml:"collector"`
	Forecast  ForecastConfig  `yaml:"forecast"`
	Cost      CostConfig      `yaml:"cost"`
}

func LoadConfig(configPath, pricingPath string) (*AppConfig, error) {
	cfg := &AppConfig{}

	// 🔹 Config principale
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read config.yaml: %w", err)
	}
	if err := yaml.Unmarshal(data, cfg); err != nil {
		return nil, fmt.Errorf("failed to parse config.yaml: %w", err)
	}

	// 🔹 Config pricing
	data2, err := os.ReadFile(pricingPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read pricing.yaml: %w", err)
	}
	var cost CostConfig
	if err := yaml.Unmarshal(data2, &cost); err != nil {
		return nil, fmt.Errorf("failed to parse pricing.yaml: %w", err)
	}
	cfg.Cost = cost

	return cfg, nil
}
