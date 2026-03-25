package collector

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"io"
)

type promResponse struct {
	Data struct {
		Result []struct {
			Metric map[string]string `json:"metric"`
			Value  []interface{}     `json:"value"`
		} `json:"result"`
	} `json:"data"`
}

func queryPrometheus(promURL, query string) ([]map[string]interface{}, error) {
	endpoint := fmt.Sprintf("%s/api/v1/query?query=%s", promURL, url.QueryEscape(query))

	resp, err := http.Get(endpoint)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result promResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	var parsed []map[string]interface{}

	for _, r := range result.Data.Result {
		parsed = append(parsed, map[string]interface{}{
			"metric": r.Metric,
			"value":  r.Value[1],
		})
	}

	return parsed, nil
}


// PrometheusClient gestisce le chiamate al server Prometheus
type PrometheusClient struct {
	baseURL string
	client  *http.Client
}

// NewPrometheusClient crea un nuovo client con l’URL del server Prometheus
func NewPrometheusClient(baseURL string) *PrometheusClient {
	return &PrometheusClient{
		baseURL: baseURL,
		client:  &http.Client{},
	}
}

// QueryResult è il formato standard della risposta Prometheus
type QueryResult struct {
	Status string `json:"status"`
	Data   struct {
		ResultType string      `json:"resultType"`
		Result     []MetricRow `json:"result"`
	} `json:"data"`
}

// MetricRow rappresenta una singola riga di metrica
type MetricRow struct {
	Metric map[string]string `json:"metric"`
	Value  [2]interface{}    `json:"value"` // timestamp, valore
}

// FetchMetrics esegue una query PromQL e restituisce le metriche
func (c *PrometheusClient) FetchMetrics(query string) ([]MetricRow, error) {
	url := fmt.Sprintf("%s/api/v1/query?query=%s", c.baseURL, query)
	resp, err := c.client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to query Prometheus: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("prometheus returned status %d: %s", resp.StatusCode, string(bodyBytes))
	}

	var result QueryResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode prometheus response: %w", err)
	}

	return result.Data.Result, nil
}
